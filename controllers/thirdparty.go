package controllers

import (
	"fmt"
	"github.com/dghubble/oauth1"
	"github.com/gin-gonic/gin"
	"github.com/koreset/go-twitter/twitter"
	"github.com/koreset/homefapp/utils"
	"net/http"
)

func GetTweets(c *gin.Context) {
	config := oauth1.NewConfig("cmLzY1seoM3RKdu7oCVFKqBiH", "BpFpYH0wTlvMNPoEzzSHrvtEk9Q5lf6q0vwf6pPSw7y5GDm0fg")
	token := oauth1.NewToken("45419796-3ckhqOkynAMdcHLTZqwN8L6859svhHb5H4BGUHEKd", "vamzHB0ZAU4wKaizV17UTgtXQHdiT99wzdy77bZtmHVHw")
	httpClient := config.Client(oauth1.NoContext, token)

	client := twitter.NewClient(httpClient)

	tweets, response, err := client.Timelines.UserTimeline(&twitter.UserTimelineParams{ScreenName: "Health_Earth", Count: 5,})

	shallowTweets := utils.GetShallowTweets(tweets)

	if err != nil {
		panic(err.Error())
	}

	if response.StatusCode == http.StatusOK {
		fmt.Println(shallowTweets)
		c.JSON(http.StatusOK, shallowTweets)
	}else{
		c.JSON(http.StatusOK, "test")
	}
}
