package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/koreset/gtf"
	"github.com/koreset/homefapp/config/bindatafs"
	"github.com/koreset/homefapp/controllers"
	"github.com/koreset/homefapp/models"
	"github.com/koreset/homefapp/services"
	"github.com/koreset/homefapp/utils"
	"github.com/qor/admin"
	"github.com/qor/media"
	"github.com/qor/media/asset_manager"
	"github.com/qor/media/media_library"
	"github.com/satori/go.uuid"
	"github.com/semihalev/gin-stats"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
	"strings"
)

var db *gorm.DB
var funcMaps template.FuncMap
var templates *template.Template

// AutoMigrate run auto migration
func AutoMigrate(values ...interface{}) {
	for _, value := range values {
		db.AutoMigrate(value)
	}
}

func SetupDB() {
	db = services.Init()
	db.AutoMigrate(&models.Post{}, &models.Video{}, &models.Image{}, &models.Link{}, &models.Category{})
	media.RegisterCallbacks(db)
}

func setupTemplateFuncs() template.FuncMap {
	funcMaps = make(template.FuncMap)
	funcMaps["unsafeHtml"] = utils.UnsafeHtml
	funcMaps["stripSummaryTags"] = utils.StripSummaryTags
	funcMaps["displayDateString"] = utils.DisplayDateString
	funcMaps["displayDate"] = utils.DisplayDateV2
	funcMaps["truncateBody"] = utils.TruncateBody

	gtf.Inject(funcMaps)
	return funcMaps
}

func RequestIdMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		xuid, _ := uuid.NewV4()
		c.Writer.Header().Set("X-Request-Id", xuid.String())
		c.Next()
	}
}

func SetupRouter() *gin.Engine {
	mux := http.NewServeMux()

	Admin := admin.New(&admin.AdminConfig{DB: db})
	Admin.SetAssetFS(bindatafs.AssetFS.NameSpace("admin"))

	Admin.MountTo("/admin", mux)

	assetManager := Admin.AddResource(&asset_manager.AssetManager{}, &admin.Config{Invisible: true})
	// Add Media Library
	Admin.AddResource(&media_library.MediaLibrary{}, &admin.Config{Menu: []string{"Site Management"}})

	post := Admin.AddResource(&models.Post{}, &admin.Config{Name: "Posts", Menu: []string{"Content Management"}})
	post.IndexAttrs("ID", "Title", "Body", "Summary", "Images", "Videos", "Links", "Type")
	post.NewAttrs("Title", "Body", "Summary", "Images", "Videos", "Links", "Type")
	post.Meta(&admin.Meta{Name: "Body", Config: &admin.RichEditorConfig{AssetManager: assetManager}})

	router := gin.Default()
	//router.Use(func(c *gin.Context) {
	//	fmt.Println("User Agent: ", c.Request.UserAgent())
	//	c.Next()
	//})
	//router.Use(RequestIdMiddleware())

	if runtime.GOOS == "linux" {
		log.Println("Loading html from binary")
		router.SetHTMLTemplate(templates)
	}

	if runtime.GOOS == "darwin" {
		router.SetFuncMap(setupTemplateFuncs())
		router.LoadHTMLGlob("views/**/*")

	}

	router.GET("/stats", func(context *gin.Context) {
		context.JSON(http.StatusOK, stats.Report())
	})
	router.GET("/", controllers.Home)
	router.GET("/about-us", controllers.AboutUs)
	router.GET("/fossil-politics", controllers.FossilPolitics)
	router.GET("/hunger-politics", controllers.HungerPolitics)
	router.GET("/resources", controllers.ResourceIndex)
	router.GET("/resources/annual-reports", controllers.ResourceAnnualReports)
	router.GET("/resources/publications", controllers.ResourcePublications)
	router.GET("/resources/publications/:page", controllers.ResourcePublications)
	router.GET("/resources/books", controllers.ResourceBooks)
	router.GET("/resources/eco-instigator", controllers.ResourceEcoInstigator)
	router.GET("/resources/gallery", controllers.ResourceGallery)
	router.GET("/resources/gallery/:albumid/:albumtitle", controllers.ResourceGalleryDetail)
	router.GET("/sustainability-academy", controllers.SustainabilityAcademy)
	router.GET("posts/:id", controllers.GetPost)
	router.GET("publications", controllers.GetPublications)
	router.GET("/test", controllers.GetTest)
	router.GET("/news", controllers.GetNews)
	router.GET("/news/:page", controllers.GetNews)
	router.GET("/new", controllers.GetNew)
	router.GET("/boot", controllers.GetBoot)

	//API Calls
	api := router.Group("/api")
	{
		api.GET("/get-tweets", controllers.GetTweets)
		api.GET("/get-flickr", controllers.GetFlickr)
		api.GET("/testdata", controllers.GetTestData)
	}

	router.Static("/public", "./public")
	router.Any("/admin/*resources", gin.WrapH(mux))
	router.NoRoute(func(context *gin.Context) {
		fmt.Println(">>>>>>>>>>>>>>>>>> 404 <<<<<<<<<<<<<<<<<<<")
		context.HTML(http.StatusNotFound, "content_not_found", nil)
	})
	return router
}

func loadTemplates() (*template.Template, error) {
	templates = template.New("")
	templates.Funcs(setupTemplateFuncs())
	var myAssets = Assets.Files

	for name, file := range myAssets {
		if file.IsDir() || !strings.HasSuffix(name, ".html") {
			continue
		}
		h, err := ioutil.ReadAll(file)
		if err != nil {
			return nil, err
		}
		templates, err = templates.New(name).Parse(string(h))
		if err != nil {
			return nil, err
		}
	}

	return templates, nil
}

func main() {
	port := flag.String("port", "4000", "The port the app will listen to")
	host := flag.String("host", "0.0.0.0", "The ip address to listen on")
	compileTemplate := flag.Bool("compile-templates", false, "Set this to true to compile templates to binary")

	flag.Parse()

	if *compileTemplate {
		Admin := admin.New(&admin.AdminConfig{
			DB:      db,
			AssetFS: bindatafs.AssetFS.NameSpace("admin")})
		Admin.SetAssetFS(bindatafs.AssetFS.NameSpace("admin"))
		bindatafs.AssetFS.Compile()
	} else {
		SetupDB()
		defer db.Close()

		loadTemplates()
		r := SetupRouter()
		fmt.Println(*host, *port)
		r.Run(fmt.Sprintf("%s:%s", *host, *port))
	}
}
