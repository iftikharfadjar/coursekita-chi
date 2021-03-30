package models

type Post struct {
	ID    ID     `json:"ID"`
	Title string `json:"title"`
	Post  string `json:"post"`
}

type IDCollection struct {
	ID    []string     `json:"ID"`
}
