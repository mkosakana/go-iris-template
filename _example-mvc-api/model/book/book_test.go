package book

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBook(t *testing.T) {
	title := "人間失格"
	author := "太宰治"
	book, err := NewBook(title, author)
	assert.NoError(t, err)
	assert.Equal(t, Book{Title: title, Author: author}, book)
}
