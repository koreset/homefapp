package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/koreset/homefapp/services"
	"net/http"
	"strconv"
)

func ResourceIndex(c *gin.Context) {
	payload := make(map[string]interface{})
	posts := services.GetPosts(0, 9)
	videos := services.GetVideos()
	payload["posts"] = posts
	payload["videos"] = videos
	payload["active"] = "hunger_politics"

	c.HTML(http.StatusOK, "resource-index", payload)

}

func ResourcePublications(c *gin.Context) {
	payload := make(map[string]interface{})
	posts := services.GetPosts(0, 9)
	videos := services.GetVideos()
	publications := services.GetPublications(0, 5)
	payload["posts"] = posts
	payload["videos"] = videos
	payload["publications"] = publications
	payload["active"] = "hunger_politics"

	for _, v := range publications {
		fmt.Println(v.Body)
		fmt.Println("===============Images===================")
		for _, img := range v.Images {
			fmt.Printf("Publication Image: %#v \n", img.File.URL("big"))
		}
		fmt.Println("==============Images====================")

		fmt.Println("===============Documents===================")
		for _, doc := range v.Documents {
			fmt.Printf("Publication Attachment: %#v \n", doc.File.URL())
		}
		fmt.Println("==============Documents====================")

	}

	c.HTML(http.StatusOK, "resource-publications", payload)

}

func ResourceAnnualReports(c *gin.Context) {
	payload := make(map[string]interface{})
	posts := services.GetPosts(0, 9)
	videos := services.GetVideos()
	payload["posts"] = posts
	payload["videos"] = videos
	payload["active"] = "hunger_politics"

	c.HTML(http.StatusOK, "resource-annual-reports", payload)

}

func ResourceBooks(c *gin.Context) {
	payload := make(map[string]interface{})
	posts := services.GetPosts(0, 9)
	videos := services.GetVideos()
	payload["posts"] = posts
	payload["videos"] = videos
	payload["active"] = "hunger_politics"

	c.HTML(http.StatusOK, "resource-books", payload)

}

func ResourceEcoInstigator(c *gin.Context) {
	payload := make(map[string]interface{})
	posts := services.GetPosts(0, 9)
	videos := services.GetVideos()
	payload["posts"] = posts
	payload["videos"] = videos
	payload["active"] = "hunger_politics"

	c.HTML(http.StatusOK, "resource-eco-instigator", payload)

}

func ResourceGallery(c *gin.Context) {
	payload := make(map[string]interface{})
	payload["albums"] = services.GetFlickrAlbums()
	c.HTML(http.StatusOK, "resource-gallery", payload)
}

func ResourceGalleryDetail(c *gin.Context) {
	albumId := c.Param("albumid")
	albumTitle := c.Param("albumtitle")
	fmt.Println("Title: ", albumTitle)
	payload := make(map[string]interface{})
	requestedAlbumId, _ := strconv.Atoi(albumId)
	payload["album"] = services.GetFlickrPhotosInAlbum(requestedAlbumId)
	payload["albumtitle"] = albumTitle
	fmt.Println(payload["album"])
	c.HTML(http.StatusOK, "resource-gallery-detail", payload)
}
