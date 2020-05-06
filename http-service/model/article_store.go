package model

type ArticleStore interface {
	All() []Article
	Save(*Article) error
	Find(int) *Article
	Update(*Article) error
	Delete(article *Article) error
}
