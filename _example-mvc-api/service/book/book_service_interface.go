//go:generate mockgen -package=$GOPACKAGE -source=$GOFILE -destination=mock_$GOFILE
//go:generate goimports mock_$GOFILE
package book

import (
	"github.com/mkosakana/go-iris-template/_example-mvc-api/model/book"
)

type BookServiceInterface interface {
	// GetBooks 本一覧取得
	GetBooks() ([]book.Book, error)

	// CreateBook 本作成
	CreateBook(bookRequest *book.BookRequest) error

	// UpdateBook 本編集
	UpdateBook(id int, bookRequest *book.BookRequest) error

	// DeleteBookById 本削除
	DeleteBookById(id int) error
}
