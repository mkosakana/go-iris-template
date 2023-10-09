package book

import (
	"github.com/kataras/iris/v12/x/errors"
	"github.com/mkosakana/go-iris-sample/_example-mvc-api/model/book"
	"gorm.io/gorm"
)

type BookService struct {
	db *gorm.DB
}

func NewBookService(db *gorm.DB) BookService {
	return BookService{db: db}
}

func (s BookService) GetBooks() ([]book.Book, error) {
	var books []book.Book
	s.db.Raw("SELECT * FROM books ORDER BY id, title, author").Scan(&books)

	return books, nil
}

func (s BookService) CreateBook(bookRequest *book.BookRequest) error {
	bookRecord, err := book.NewBook(bookRequest.Title, bookRequest.Author)
	if err != nil {
		return err
	}

	// https://gorm.io/docs/transactions.html#Transaction
	return s.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&bookRecord).Error; err != nil {
			return err
		}

		return nil
	})
}

func (s BookService) UpdateBook(id int, bookRequest *book.BookRequest) error {
	if id <= 0 {
		return errors.New("不正なIDです。")
	}

	// https://gorm.io/docs/transactions.html#Transaction
	return s.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Table("books").Where("id = ?", id).Updates(bookRequest).Error; err != nil {
			return err
		}

		return nil
	})
}

func (s BookService) DeleteBookById(id int) error {
	if id <= 0 {
		return errors.New("不正なIDです。")
	}

	// https://gorm.io/docs/transactions.html#Transaction
	return s.db.Transaction(func(tx *gorm.DB) error {
		var bookRecord book.Book
		if err := tx.Table("books").Where("id = ?", id).Delete(&bookRecord).Error; err != nil {
			return err
		}

		return nil
	})
}
