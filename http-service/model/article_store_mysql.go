package model

import (
	"database/sql"
	"log"
)

type ArticleStoreMySQL struct {
	DB *sql.DB
}

func NewArticleStoreMySQL() ArticleStore {
	dsn := "root:password@tcp(localhost:3307)/database?parseTime=true&clientFoundRows=true"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	return &ArticleStoreMySQL{DB: db}
}

func (store *ArticleStoreMySQL) All() []Article {
	articles := []Article{}
	rows, err := store.DB.Query("SELECT * FROM articles")
	if err != nil {
		return articles
	}

	article := Article{}
	for rows.Next() {
		rows.Scan(&article.ID, &article.Title, &article.Body)
		articles = append(articles, article)
	}

	return articles
}

func (store *ArticleStoreMySQL) Save(article *Article) error {
	result, err := store.DB.Exec(`
		INSERT INTO articles(title, body) VALUES(?,?)`,
		article.Title,
		article.Body,
	)
	if err != nil {
		return err
	}

	_, err = result.RowsAffected()
	if err != nil {
		return err
	}

	// update article.ID
	lastID, err := result.LastInsertId()
	if err != nil {
		return err
	}

	article.ID = int(lastID)

	return nil
}

func (store *ArticleStoreMySQL) Find(id int) *Article {
	article := Article{}

	err := store.DB.
		QueryRow(`SELECT * FROM articles WHERE id=?`, id).
		Scan(
			&article.ID,
			&article.Title,
			&article.Body,
		)

	if err != nil {
		log.Fatal(err)
		return nil
	}

	return &article
}

func (store *ArticleStoreMySQL) Update(article *Article) error {
	result, err := store.DB.Exec(`
		UPDATE articles SET title = ?, body = ? WHERE id = ?`,
		article.Title,
		article.Body,
		article.ID,
	)
	if err != nil {
		return err
	}

	_, err = result.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

func (store *ArticleStoreMySQL) Delete(article *Article) error {
	result, err := store.DB.Exec(`
		DELETE FROM articles WHERE id = ?`,
		article.ID,
	)
	if err != nil {
		return err
	}

	_, err = result.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}
