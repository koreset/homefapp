package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/koreset/homefapp/services"
)

func Home(c *gin.Context) {
	payload := make(map[string]interface{})
	postsTop := services.GetPosts(0, 3)
	postsBottom := services.GetPosts(4, 3)
	videos := services.GetVideos()
	publications := services.GetPublications()
	payload["postsTop"] = postsTop
	payload["postsBottom"] = postsBottom
	payload["videos"] = videos
	payload["publications"] = publications
	payload["active"] = "home_page"
	payload["title"] = "Home"

	fmt.Println("Request from: ", c.ClientIP())
	fmt.Println(c.Request.UserAgent())
	c.HTML(http.StatusOK, "home", payload)

}
