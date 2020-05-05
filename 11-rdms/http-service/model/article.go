package model

// Article is wrap article domain
type Article struct {
	ID    int
	Title string
	Body  string
}

// CreateArticle to create article instance
func CreateArticle(title, body string) (*Article, error) {
	return &Article{
		Title: title,
		Body:  body,
	}, nil
}
