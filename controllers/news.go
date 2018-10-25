package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/koreset/homefapp/services"
	"net/http"
)

func GetNews(context *gin.Context) {
	payload := make(map[string]interface{})
	newsItems := services.GetPosts(0,10)
	recentPosts := services.GetPosts(0, 5)
	payload["recentPosts"] = recentPosts
	payload["newsitems"] = newsItems
	payload["title"] = "News"

	context.HTML(http.StatusOK, "news_index", payload)
}
