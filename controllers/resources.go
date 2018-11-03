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
	var limit = 5
	page, err := strconv.Atoi(c.Param("page"))

	if err != nil{
		page = 0
	}

	payload := make(map[string]interface{})
	posts := services.GetPosts(0, 5)
	publications := services.GetPublications(limit * page, 5)
	payload["posts"] = posts
	payload["publications"] = publications
	payload["active"] = "publications"
	payload["nextPage"] = page + 1

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
