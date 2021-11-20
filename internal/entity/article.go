package entity

type ArticleID int

// Article -.
type Article struct {
	ID     ArticleID `json:"id" example:"1"`
	Author string    `json:"author" example:"Pushkin"`
	Text   string    `json:"text"   example:"en"`
}
