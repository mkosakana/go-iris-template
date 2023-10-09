package book

import (
	"time"

	"github.com/kataras/iris/v12/x/errors"
)

type Book struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	Title     string
	Author    string
}

func NewBook(title string, author string) (Book, error) {
	if title == "" || author == "" {
		return Book{}, errors.New("bookテーブルのレコード作成に失敗しました。")
	}
	return Book{
		Title:  title,
		Author: author,
	}, nil
}
