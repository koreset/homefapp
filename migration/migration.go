// +build migration

package main

import (
	"fmt"
	"github.com/gosimple/slug"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/koreset/homefapp/models"
	"github.com/koreset/homefapp/utils"
	"github.com/qor/media"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

var newDB *gorm.DB
var homefDB *gorm.DB
var dbError error

func baseMigration() {

	var categoryQuery = "select tid as id, name from taxonomy_term_data"

	catrows, _ := homefDB.Raw(categoryQuery).Rows()

	for catrows.Next() {
		var category models.Category

		homefDB.ScanRows(catrows, &category)
		newDB.Create(&category)
		newDB.Save(&category)
	}

	defer catrows.Close()

	var baseQuery = "select nid as id, title, type, created, changed as updated from node"
	rows, _ := homefDB.Raw(baseQuery).Rows()

	for rows.Next() {
		var post models.Post
		homefDB.ScanRows(rows, &post)
		newDB.Create(&post)
		newDB.Save(&post)
	}

	defer rows.Close()
}

func populateArticleBody() {
	var posts []models.Post
	newDB.Find(&posts)

	for _, v := range posts {
		rows, _ := homefDB.Raw("select body_value as body, body_summary as summary from field_data_body where entity_id = ? ", v.ID).Rows()

		for rows.Next() {
			homefDB.ScanRows(rows, &v)
		}

		v.Body = utils.CleanHtmlBody(v.Body)

		fmt.Println("ID: ", v.ID)
		fmt.Println("Body: ", v.Body)

		newDB.Save(&v)
	}
}

func populateCategories(){
	var posts []models.Post
	newDB.Find(&posts)
	for _, v := range posts {
		rows, _ := homefDB.Raw("select tid as category_id from taxonomy_index where nid  = ? ", v.ID).Rows()

		for rows.Next() {
			homefDB.ScanRows(rows, &v)
		}

		fmt.Println("CategoryID: ", v.CategoryID)
		fmt.Println("Post ID: ", v.ID)

		newDB.Save(&v)
	}

}

type TempImage struct {
	ID       uint
	FileName string
	Url      string
}

func transformString(file string) string {
	parts := strings.Split(file, ".")
	extension := parts[len(parts)-1]
	filename := strings.TrimSuffix(file, "."+extension)
	modFileName := slug.MakeLang(filename, "en") + "." + extension

	return modFileName

}

func populateImages() {
	//storage := filesystem.New("./public")
	var posts []models.Post
	newDB.Find(&posts)
	workDirectory := "/Users/jome/projects/homef/files/"
	for _, v := range posts {

		rows, _ := homefDB.Raw("select entity_id as id, filename as file_name, uri as url from field_data_field_image, file_managed where field_image_fid = fid and entity_id = ?", v.ID).Rows()

		for rows.Next() {
			var image TempImage
			var imageItem models.Image

			homefDB.ScanRows(rows, &image)
			fmt.Printf("Image Temp: %+v \n", image)

			image.Url = strings.Replace(image.Url, "public://", "", -1)
			filePath := workDirectory + image.Url
			theFile, err := os.Open(filePath)

			if err != nil {
				fmt.Println(err.Error())
				continue
			}else{
				newFileName := transformString(image.FileName)
				//newPath := "/content/images/" + strconv.Itoa(int(v.ID)) + "/" + newFileName
				//storage.Put(newPath, theFile)

				imageItem = models.Image{
					PostID: v.ID,
				}

				imageItem.File.Sizes = imageItem.GetSizes()
				imageItem.File.Scan(theFile)
				imageItem.File.FileName = newFileName
				fmt.Printf("Image Item: %+v \n", imageItem.File.FileName)
				v.Images = append(v.Images, imageItem)
			}
		}
		newDB.Save(&v)
	}
}

func populateVideoItems() {
	var posts []models.Post
	newDB.Find(&posts)
	for _, v := range posts {
		rows, _ := homefDB.Raw("select field_video_value as value, field_video_description_value as description from field_data_field_video as fdv, field_data_field_video_description as fvdd where fdv.entity_id = ? and fvdd.entity_id = ?", v.ID, v.ID).Rows()
		for rows.Next() {
			if v.Type == "video" {
				var video models.Video
				homefDB.ScanRows(rows, &video)
				fmt.Println(video)

				v.Body = video.Description
				v.Videos = append(v.Videos, video)
			}

			newDB.Save(&v)
		}
	}
}

func populateLinks() {
	var posts []models.Post
	newDB.Find(&posts)
	for _, v := range posts {
		rows, _ := homefDB.Raw("select field_url_url as url, field_url_title as title from field_data_field_url where entity_id = ?", v.ID).Rows()
		for rows.Next() {
			var link models.Link
			homefDB.ScanRows(rows, &link)
			fmt.Println(link)
			v.Links = append(v.Links, link)
		}

		newDB.Save(&v)
	}

}

func populatePublications() {
	var posts []models.Post
	newDB.Where("type = ?", "publication").Find(&posts)
	for _, v := range posts {
		rows, _ := homefDB.Raw("select field_url_url as url, field_url_title as title from field_data_field_url where entity_id = ?", v.ID).Rows()
		for rows.Next() {
			var link models.Link
			homefDB.ScanRows(rows, &link)
			fmt.Println(link.Url)
			file, err := openFileByURL(link.Url)

			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(file.Name())

				document := models.Document{}

				document.File.Scan(file)
				v.Documents = append(v.Documents, document)
			}
		}

		newDB.Save(&v)
	}

}

func openFileByURL(rawURL string) (*os.File, error) {
	if fileURL, err := url.Parse(rawURL); err != nil {
		return nil, err
	} else {
		path := fileURL.Path
		segments := strings.Split(path, "/")
		fileName := segments[len(segments)-1]

		filePath := filepath.Join(os.TempDir(), fileName)

		if _, err := os.Stat(filePath); err == nil {
			return os.Open(filePath)
		}

		file, err := os.Create(filePath)
		if err != nil {
			return file, err
		}

		check := http.Client{
			CheckRedirect: func(r *http.Request, via []*http.Request) error {
				r.URL.Opaque = r.URL.Path
				return nil
			},
		}
		resp, err := check.Get(rawURL) // add a filter to check redirect
		if err != nil {
			return file, err
		}
		defer resp.Body.Close()
		fmt.Printf("----> Downloaded %v\n", rawURL)

		_, err = io.Copy(file, resp.Body)
		if err != nil {
			return file, err
		}
		return file, nil
	}
}

func main() {
	//awsConnection := "homef:wordpass15@tcp(rds-mysql-homef.cb44dbuhyviz.eu-west-2.rds.amazonaws.com:3306)/homef?charset=utf8&parseTime=True&loc=Local"
	localConnection := "root:wordpass15@tcp(localhost:3306)/homef?charset=utf8&parseTime=True&loc=Local"
	drupalConnection := "onajome:wordpass15@tcp(mysql.homef.org:3306)/homef_db?charset=utf8&parseTime=True&loc=Local"
	newDB, dbError = gorm.Open("mysql", localConnection)
	if dbError != nil {
		panic(dbError)
	}

	homefDB, dbError = gorm.Open("mysql", drupalConnection)
	if dbError != nil {
		panic(dbError)
	}

	//newDB.LogMode(true)
	//homefDB.LogMode(true)

	var category models.Category
	var posts []models.Post
	var post models.Post
	var video []models.Video
	var image []models.Image
	var link []models.Link
	var documents []models.Document

	newDB.Model(&category).Related(&posts)
	newDB.Model(&post).Related(&video)
	newDB.Model(&post).Related(&image)
	newDB.Model(&post).Related(&link)
	newDB.Model(&post).Related(&documents)
	newDB.AutoMigrate(&models.Category{}, &models.Post{}, &models.Document{}, &models.Video{}, &models.Image{}, &models.Link{})
	media.RegisterCallbacks(newDB)

	//baseMigration()
	//populateArticleBody()
	//populateCategories()
	populateImages()
	//populateVideoItems()
	//populateLinks()
	//populatePublications()

}
