package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/koreset/homefapp/services"
	"net/http"
)

func FossilPolitics(c *gin.Context) {
	payload := make(map[string]interface{})
	//posts := services.GetPosts(0, 3)
	posts := services.GetPostsForCategory(0, 3, "Fossil Politics")
	videos := services.GetVideos()
	payload["posts"] = posts
	payload["videos"] = videos
	payload["active"] = "fossil_politics"
	payload["title"] = "Fossil Politics"

	c.HTML(http.StatusOK, "fossil-politics", payload)

}

func HungerPolitics(c *gin.Context) {
	payload := make(map[string]interface{})
	posts := services.GetPostsForCategory(0, 3, "Hunger Politics")
	videos := services.GetVideos()
	payload["posts"] = posts
	payload["videos"] = videos
	payload["active"] = "hunger_politics"
	payload["title"] = "Hunger Politics"

	c.HTML(http.StatusOK, "hunger-politics", payload)

}

func SustainabilityAcademy(c *gin.Context) {
	payload := make(map[string]interface{})
	posts := services.GetPostsForCategory(0, 3, "Sustainability Academy")
	gallery, _ := services.GetFlickrImages(9)
	payload["gallery"] = gallery
	payload["posts"] = posts
	payload["active"] = "sustainability_politics"
	payload["title"] = "Sustainability Academy"

	c.HTML(http.StatusOK, "sustainability-academy", payload)

}
