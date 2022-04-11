package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/httptest"
	"testing"
)

func TestApp(t *testing.T) {
	app := iris.New()
	app.Handle("GET", "/", ping)
	app.Handle("PUT", "/books/{id:string}", updateBook)

	e := httptest.New(t, app)
	e.GET("/").Expect().
		Status(httptest.StatusOK).
		Body().
		Equal("ping")

	e.PUT("/books/123").WithJSON(
		iris.Map{
			"title": "老人と森",
		},
		).Expect().
		Status(httptest.StatusOK)
}
