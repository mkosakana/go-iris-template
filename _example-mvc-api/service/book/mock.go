package book

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mkosakana/go-iris-sample/_example-mvc-api/model/book"
)

func GetMockBookService(t *testing.T) *MockBookServiceInterface {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := NewMockBookServiceInterface(ctrl)
	mock.setGetBooks()
	mock.setCreateBook()
	mock.setUpdateBook()
	mock.setDeleteBookById()

	return mock
}

func (m *MockBookServiceInterface) setGetBooks() {
	m.
		EXPECT().
		GetBooks().
		Return([]book.Book{
			{
				ID:     1,
				Title:  "コンビニ人間",
				Author: "村田 沙耶香",
			},
			{
				ID:     2,
				Title:  "老人と海",
				Author: "アーネスト・ヘミングウェイ",
			},
			{
				ID:     3,
				Title:  "歯車",
				Author: "芥川 龍之介",
			},
		}, nil).
		AnyTimes()
}

func (m *MockBookServiceInterface) setCreateBook() {
	m.
		EXPECT().
		CreateBook(gomock.Any()).
		Return(nil).
		AnyTimes()
}

func (m *MockBookServiceInterface) setUpdateBook() {
	m.
		EXPECT().
		UpdateBook(gomock.Any(), gomock.Any()).
		Return(nil).
		AnyTimes()
}

func (m *MockBookServiceInterface) setDeleteBookById() {
	m.
		EXPECT().
		DeleteBookById(gomock.Any()).
		Return(nil).
		AnyTimes()
}
