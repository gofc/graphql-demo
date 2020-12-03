package model

type Article struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	AuthorID string `json:"authorId"`
}
