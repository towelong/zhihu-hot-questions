package model

type Question struct {
	Title string `json:"title"`
	Url string `json:"url"`
}

type Target struct {
	Title string `json:"title"`
	Id int32 `json:"id"`
}

type Item struct {
	Target Target `json:"target"`
}

type HotList struct {
	Data []Item `json:"data"`
}
