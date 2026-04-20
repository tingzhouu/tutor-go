package model

type Bookmark struct {
	Id    int    `json:"id"`
	Url   string `json:"url"`
	Title string `json:"title"`
	Tags  string `json:"tags"`
}

// id, url, title, tags
