package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/purwandi/http-service/model"
)

func main() {
	// init data store
	store := model.NewArticleStoreInmemory()

	e := echo.New()
	e.GET("/articles", func(c echo.Context) error {
		articles := store.ArticleMap
		return c.JSON(http.StatusOK, articles)
	})

	e.POST("/articles", func(c echo.Context) error {
		title := c.FormValue("title")
		body := c.FormValue("body")

		// create article instance
		article := model.CreateArticle(title, body)

		store.Save(article)

		return c.JSON(http.StatusOK, article)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
