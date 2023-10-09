package book

import (
	"os"
	"testing"

	"github.com/mkosakana/go-iris-template/_example-mvc-api/configs/db"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func getTestDb() (*gorm.DB, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return nil, nil
	}
	pathToEnvFile := currentDir + "/../../environment/.env"
	testDb, err := db.GetTestGormDb(pathToEnvFile)
	if err != nil {
		return nil, nil
	}

	return testDb, nil
}

func getBookService() (BookService, error) {
	testDb, err := getTestDb()
	if err != nil {
		return BookService{}, nil
	}
	return NewBookService(testDb), nil
}

func TestNewBookService(t *testing.T) {
	currentDir, err := os.Getwd()
	assert.NoError(t, err)
	pathToEnvFile := currentDir + "/../../environment/.env"
	testDb, err := db.GetTestGormDb(pathToEnvFile)
	assert.NoError(t, err)

	bookService := NewBookService(testDb)
	assert.Equal(t, BookService{db: testDb}, bookService)
}

func TestBookService_GetBooks(t *testing.T) {
	bookService, err := getBookService()
	assert.NoError(t, err)

	books, err := bookService.GetBooks()
	assert.NoError(t, err)
	assert.Equal(t, uint(1), books[0].ID)
	assert.Equal(t, "コンビニ人間", books[0].Title)
	assert.Equal(t, "村田 沙耶香", books[0].Author)
	assert.Equal(t, uint(3), books[len(books)-1].ID)
	assert.Equal(t, "歯車", books[len(books)-1].Title)
	assert.Equal(t, "芥川 龍之介", books[len(books)-1].Author)
}
