package setups

import (
	"github.com/kataras/iris/v12/mvc"
	"github.com/mkosakana/go-iris-template/_example-mvc-api/controller/books"
	bookService "github.com/mkosakana/go-iris-template/_example-mvc-api/service/book"
)

func SetBooksController(app *mvc.Application) {
	SetBaseConfiguration(app)
	app.Register(bookService.NewBookService)
	app.Handle(new(books.BooksController))
}
