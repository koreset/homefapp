package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AboutUs(c *gin.Context){
	payload := make(map[string] interface{})
	payload["active"] = "aboutus"
	c.HTML(http.StatusOK, "about-us", payload)
}
