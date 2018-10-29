package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/koreset/homefapp/services"
	"net/http"
)

func ResourceIndex(c *gin.Context) {
	payload := make(map[string]interface{})
	posts := services.GetPosts(0, 9)
	videos := services.GetVideos()
	publications := services.GetPublications()
	payload["posts"] = posts
	payload["videos"] = videos
	payload["publications"] = publications
	payload["active"] = "hunger_politics"

	c.HTML(http.StatusOK, "resource-index", payload)

}

func ResourcePublications(c *gin.Context) {
	payload := make(map[string]interface{})
	posts := services.GetPosts(0, 9)
	videos := services.GetVideos()
	publications := services.GetPublications()
	payload["posts"] = posts
	payload["videos"] = videos
	payload["publications"] = publications
	payload["active"] = "hunger_politics"

	c.HTML(http.StatusOK, "resource-publications", payload)

}

func ResourceAnnualReports(c *gin.Context) {
	payload := make(map[string]interface{})
	posts := services.GetPosts(0, 9)
	videos := services.GetVideos()
	publications := services.GetPublications()
	payload["posts"] = posts
	payload["videos"] = videos
	payload["publications"] = publications
	payload["active"] = "hunger_politics"

	c.HTML(http.StatusOK, "resource-annual-reports", payload)

}

func ResourceBooks(c *gin.Context) {
	payload := make(map[string]interface{})
	posts := services.GetPosts(0, 9)
	videos := services.GetVideos()
	publications := services.GetPublications()
	payload["posts"] = posts
	payload["videos"] = videos
	payload["publications"] = publications
	payload["active"] = "hunger_politics"

	c.HTML(http.StatusOK, "resource-books", payload)

}

func ResourceEcoInstigator(c *gin.Context) {
	payload := make(map[string]interface{})
	posts := services.GetPosts(0, 9)
	videos := services.GetVideos()
	publications := services.GetPublications()
	payload["posts"] = posts
	payload["videos"] = videos
	payload["publications"] = publications
	payload["active"] = "hunger_politics"

	c.HTML(http.StatusOK, "resource-eco-instigator", payload)

}

func ResourceGallery(c *gin.Context) {
	payload := make(map[string]interface{})
	payload["albums"] = services.GetFlickrAlbums()
	c.HTML(http.StatusOK, "resource-gallery", payload)
}
