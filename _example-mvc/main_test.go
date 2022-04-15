package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/httptest"
	"testing"
)

func TestPing(t *testing.T) {
	app := iris.New()
	app.Handle("GET", "/", ping)

	e := httptest.New(t, app)
	e.GET("/").Expect().
		Status(httptest.StatusOK).
		Body().
		Equal("ping")
}
