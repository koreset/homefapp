package models

import (
	"fmt"
	"github.com/gosimple/slug"
	"github.com/jinzhu/gorm"
	"github.com/qor/media"
	"github.com/qor/media/media_library"
	"github.com/qor/media/oss"
	"time"
)

type Category struct {
	ID    uint `gorm:"primary_key"`
	Name  string
	Posts []Post `gorm:"many2many:category_post"`
}

type Link struct {
	gorm.Model
	Url      string
	Title    string
	ImageUrl string
	PostID   uint
}

type Video struct {
	gorm.Model
	Url         string
	Value       string `gorm:"type:longtext"`
	Description string `gorm:"type:longtext"`
	PostID      uint
}

type Image struct {
	gorm.Model
	File   media_library.MediaLibraryStorage `gorm:"type:longtext" sql:"size:4294967295;" media_library:"url:/content/{{class}}/{{primary_key}}/{{column}}.{{extension}};path:./public"`
	PostID uint
}

func (Image) GetSizes() map[string]*media.Size {
	return map[string]*media.Size{
		"small":           {Width: 320, Height: 320},
		"middle":          {Width: 640, Height: 640},
		"big":             {Width: 1024, Height: 720},
		"article_preview": {Width: 390, Height: 300},
		"preview":         {Width: 200, Height: 200},
	}
}

type Document struct {
	gorm.Model
	File   oss.OSS `gorm:"type:longtext" sql:"size:4294967295;" media_library:"url:/content/publications/{{basename}}.{{extension}};path:./public"`
	PostID uint
}

type Post struct {
	ID         uint       `gorm:"primary_key"`
	Categories []Category `gorm:"many2many:category_post"`
	Title      string
	Slug       string `gorm:"unique"`
	Body       string `gorm:"type:longtext"`
	Summary    string `gorm:"type:longtext"`
	Images     []Image
	Documents  []Document
	Videos     []Video
	Links      []Link
	Type       string
	Created    int32
	Updated    int32
}

type Event struct {
	ID         uint       `gorm:"primary_key"`
	Categories []Category `gorm:"many2many:category_post"`
	Title      string
	Slug       string `gorm:"unique"`
	Body       string `gorm:"type:longtext"`
	Summary    string `gorm:"type:longtext"`
	Images     []Image
	Documents  []Document
	Videos     []Video
	Links      []Link
	Type       string
	Created    int32
	Updated    int32
	StartDate  int32
	EndDate    int32
}

func (p *Post) BeforeCreate() (err error) {
	if p.Created == 0 {
		p.Created = int32(time.Now().Unix())
	}

	if p.Updated == 0 {
		p.Updated = int32(time.Now().Unix())
	}

	p.Slug = createUniqueSlug(p.Title)

	fmt.Printf("======> New post: %#v", p.Title)
	fmt.Printf("======> New post: %#v", p.Summary)
	fmt.Printf("======> New post: %#v", p.Images)
	fmt.Printf("=======> Images: %#v", p.Images[0].File.FileName)

	for i := range p.Images {
		p.Images[i].File.Sizes = p.Images[i].GetSizes()
		file, _ := p.Images[i].File.Base.FileHeader.Open()
		p.Images[i].File.Scan(file)
	}

	return nil
}

func createUniqueSlug(title string) string {
	slugTitle := slug.Make(title)
	if len(slugTitle) > 50 {
		slugTitle = slugTitle[:50]
		if slugTitle[len(slugTitle)-1:] == "-" {
			slugTitle = slugTitle[:len(slugTitle)-1]
		}
	}
	return slugTitle
}
