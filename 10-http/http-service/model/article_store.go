package model

type ArticleStoreInmemory struct {
	ArticleMap []Article
}

func NewArticleStoreInmemory() *ArticleStoreInmemory {
	return &ArticleStoreInmemory{
		ArticleMap: []Article{
			Article{ID: 1, Title: "Membuat website", Body: "Hallo ini bydi"},
		},
	}
}

func (store *ArticleStoreInmemory) Save(article *Article) error {
	lastID := len(store.ArticleMap)

	// set article id
	article.ID = lastID + 1

	// push to article map slices
	store.ArticleMap = append(store.ArticleMap, *article)

	return nil
}
