package models


type Payload struct {
	Photos Photos `json:"photos"`
}

type Photos struct {
	Photo []PhotoItem `json:"photo"`
}

type PhotoItem struct {
	Id string `json:"id"`
	UrlMedium string `json:"url_m"`
	UrlThumbnail string `json:"url_t"`
	UrlSquare string `json:"url_sq"`
}
