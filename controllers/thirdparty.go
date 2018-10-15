package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/dghubble/oauth1"
	"github.com/gin-gonic/gin"
	"github.com/koreset/go-twitter/twitter"
	"github.com/koreset/homefapp/utils"
	"io/ioutil"
	"net/http"
	"github.com/koreset/homefapp/models"
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

func GetFlickr(c *gin.Context) {
	photostreamUrl := "https://api.flickr.com/services/rest/?method=flickr.photos.search&api_key=b5fd7ac0bc2b2e1670312fa98fbe0ae8&user_id=100756072%40N02&extras=url_sq%2Curl_t%2Curl_m&per_page=9&format=json&nojsoncallback=1"
	response, err := http.Get(photostreamUrl)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer response.Body.Close() //Response.Body is of type io.ReadCloser *Look this up later"
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	//var payload Payload
	var results models.Payload
	json.Unmarshal(body, &results)

	c.JSON(http.StatusOK, results)

}
