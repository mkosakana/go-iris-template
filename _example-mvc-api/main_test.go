package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/httptest"
	"testing"
)

func Test_Ping(t *testing.T) {
	app := iris.New()
	app.Handle("GET", "/ping", ping)

	e := httptest.New(t, app)
	e.GET("/ping").Expect().
		Status(httptest.StatusOK).
		Body().
		NotEmpty()
}

func Test_HandleDir(t *testing.T) {
	app := iris.New()
	app.HandleDir("/", iris.Dir("./public"))

	e := httptest.New(t, app)
	e.GET("/").Expect().
		Status(httptest.StatusOK).
		ContentType("text/html", "utf-8").
		Body().
		NotEmpty()
}
