package books

import (
	"testing"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/httptest"
	"github.com/kataras/iris/v12/mvc"
	"github.com/mkosakana/go-iris-sample/_example-mvc-api/service/book"
)

func TestBooksController_GetList(t *testing.T) {
	app := iris.New()

	books := app.Party("/books")
	mvc.Configure(books, func(app *mvc.Application) {
		app.Register(book.GetMockBookService(t))
		app.Handle(new(BooksController))
	})

	e := httptest.New(t, app)
	e.GET("/books/list").
		Expect().
		Status(httptest.StatusOK)
}

func TestBooksController_Post(t *testing.T) {
	app := iris.New()

	books := app.Party("/books")
	mvc.Configure(books, func(app *mvc.Application) {
		app.Register(book.GetMockBookService(t))
		app.Handle(new(BooksController))
	})

	e := httptest.New(t, app)
	e.POST("/books").
		WithJSON(
			iris.Map{
				"title":  "人間失格",
				"author": "太宰治",
			}).
		Expect().
		Status(httptest.StatusCreated)
}

func TestBooksController_PutDetailsBy(t *testing.T) {
	app := iris.New()

	books := app.Party("/books")
	mvc.Configure(books, func(app *mvc.Application) {
		app.Register(book.GetMockBookService(t))
		app.Handle(new(BooksController))
	})

	e := httptest.New(t, app)
	e.PUT("/books/details/1").
		WithJSON(
			iris.Map{
				"title":  "人間失格",
				"author": "太宰治",
			}).
		Expect().
		Status(httptest.StatusOK)
}

func TestBooksController_DeleteDetailsBy(t *testing.T) {
	app := iris.New()

	books := app.Party("/books")
	mvc.Configure(books, func(app *mvc.Application) {
		app.Register(book.GetMockBookService(t))
		app.Handle(new(BooksController))
	})

	e := httptest.New(t, app)
	e.DELETE("/books/details/1").
		Expect().
		Status(httptest.StatusOK)
}
