package models


type Payload struct {
	Photos Photos `json:"photos"`
}

type Photos struct {
	PhotoItems []PhotoItem `json:"photo"`
}

type PhotoItem struct {
	Id string `json:"id"`
	UrlMedium string `json:"url_m"`
	UrlThumbnail string `json:"url_t"`
	UrlMedium2 string `json:"url_n"`
	UrlSquare string `json:"url_sq"`
	UrlSmall string `json:"url_s"`
	UrlLarge string `json:"url_l"`
}
