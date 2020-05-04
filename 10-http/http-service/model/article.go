package model

type Article struct {
	ID    int
	Title string
	Body  string
}

func CreateArticle(title, body string) *Article {
	return &Article{
		Title: title,
		Body:  body,
	}
}
