package books

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	bookModel "github.com/mkosakana/go-iris-sample/_example-mvc-api/model/book"
	"github.com/mkosakana/go-iris-sample/_example-mvc-api/service/book"
)

type BooksController struct {
	Ctx                  iris.Context
	BookServiceInterface book.BookServiceInterface
}

// IrisのMVCコントローラーではメソッド名を対応させることによってパスの違い・リクエストメソッド・パラメータを振り分けることができる
// 例
// app.Party("/books")
// GetList()                GET:    /books/list
// Post()                   POST:   /books
// PostDetailsExample()     POST:   /books/details/example
// Put()                    PUT:    /books
// PutDetails()             PUT:    /books/details
// PutDetailsBy(id uint)    PUT:    /books/details/{id:int}
// DeleteDetails()          DELETE: /books/details
// DeleteDetailsBy(id uint) DELETE: /books/details/{id:int}

func (c *BooksController) GetList() mvc.Response {
	books, err := c.BookServiceInterface.GetBooks()
	if err != nil {
		return mvc.Response{Code: iris.StatusInternalServerError, Err: err}
	}

	return mvc.Response{Code: iris.StatusOK, Object: books}
}

func (c *BooksController) Post() mvc.Response {
	createBookParams, err := c.getRequestBody()
	if err != nil {
		return mvc.Response{Code: iris.StatusBadRequest, Err: err}
	}

	err = c.BookServiceInterface.CreateBook(&createBookParams)
	if err != nil {
		return mvc.Response{Code: iris.StatusInternalServerError, Err: err}
	}

	return mvc.Response{Code: iris.StatusCreated}
}

func (c *BooksController) PutDetailsBy(id int) mvc.Response {
	updateBookParam, err := c.getRequestBody()
	if err != nil {
		return mvc.Response{Code: iris.StatusBadRequest, Err: err}
	}

	err = c.BookServiceInterface.UpdateBook(id, &updateBookParam)
	if err != nil {
		return mvc.Response{Code: iris.StatusInternalServerError, Err: err}
	}

	return mvc.Response{Code: iris.StatusOK}
}

func (c *BooksController) DeleteDetailsBy(id int) mvc.Response {
	err := c.BookServiceInterface.DeleteBookById(id)
	if err != nil {
		return mvc.Response{Code: iris.StatusInternalServerError, Err: err}
	}

	return mvc.Response{Code: iris.StatusOK}
}

// リクエストボディのjsonデータを構造体に格納する
func (c *BooksController) getRequestBody() (bookModel.BookRequest, error) {
	var bookRequest bookModel.BookRequest
	err := c.Ctx.ReadJSON(&bookRequest)
	if err != nil {
		return bookRequest, err
	}

	return bookRequest, nil
}
