package model

// ArticleStoreInmemory is to article store in memory implementation
type ArticleStoreInmemory struct {
	ArticleMap []Article
}

// NewArticleStoreInmemory to create ArticleStoreInmemory instance
func NewArticleStoreInmemory() *ArticleStoreInmemory {
	return &ArticleStoreInmemory{
		ArticleMap: []Article{
			Article{ID: 1, Title: "Membuat website", Body: "Hallo ini bydi"},
			Article{ID: 2, Title: "Membuat website 2", Body: "Hallo ini bydi 2"},
		},
	}
}

// Save for save
func (store *ArticleStoreInmemory) Save(article *Article) error {
	lastID := len(store.ArticleMap)

	// set article id
	article.ID = lastID + 1

	// push to article map slices
	store.ArticleMap = append(store.ArticleMap, *article)

	return nil
}

// Find for find
func (store *ArticleStoreInmemory) Find(id int) *Article {
	article := Article{}
	for _, item := range store.ArticleMap {
		if item.ID == id {
			article = item
		}
	}

	return &article
}

// Update for update article
func (store *ArticleStoreInmemory) Update(article *Article) error {
	for index, item := range store.ArticleMap {
		if item.ID == article.ID {
			store.ArticleMap[index] = *article
		}
	}

	return nil
}

// Delete to delete article
func (store *ArticleStoreInmemory) Delete(article *Article) error {
	articles := []Article{}

	for _, item := range store.ArticleMap {
		if item.ID != article.ID {
			articles = append(articles, item)
		}
	}

	store.ArticleMap = articles
	return nil
}
