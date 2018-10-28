package services

import (
	"encoding/json"
	"fmt"
	"github.com/koreset/homefapp/models"
	"io/ioutil"
	"net/http"
	"strconv"
)

func GetFlickrImages(number int) (payload models.Payload, err error){
	photostreamUrl := "https://api.flickr.com/services/rest/?method=flickr.photos.search&api_key=b5fd7ac0bc2b2e1670312fa98fbe0ae8&user_id=100756072%40N02&extras=url_sq%2Curl_t%2Curl_m%2Curl_b%2Curl_l%2Curl_n&per_page="+ strconv.Itoa(number) + "&format=json&nojsoncallback=1"

	response, err := http.Get(photostreamUrl)
	if err != nil {
		fmt.Println(err.Error())
		return payload, err
	}

	defer response.Body.Close() //Response.Body is of type io.ReadCloser *Look this up later"
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err.Error())
	}

	json.Unmarshal(body, &payload)
	fmt.Println("Flickr: ", payload)
	return payload, nil
}
