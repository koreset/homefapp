package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/koreset/homefapp/services"
	"net/http"
)

func GetTest(c *gin.Context){
	//payload := make(map[string] interface{} )
	//
	//posts := services.GetPosts(0, 4)
	//payload["posts"] = posts
	payload := make(map[string]interface{})
	payload["albums"] = services.GetFlickrAlbums()

	c.HTML(http.StatusOK, "testing.html", payload)
}


func GetNew(c *gin.Context){
	payload := make(map[string] interface{} )

	posts := services.GetPosts(0, 4)
	payload["posts"] = posts


	c.HTML(http.StatusOK, "new", payload)
}


func GetBoot(c *gin.Context){
	payload := make(map[string] interface{} )

	postsTop := services.GetPosts(0, 3)
	postsBottom := services.GetPosts(4, 3)
	payload["postsTop"] = postsTop
	payload["postsBottom"] = postsBottom


	c.HTML(http.StatusOK, "home.html", payload)
}

