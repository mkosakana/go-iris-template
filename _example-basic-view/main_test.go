package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/httptest"
	"testing"
)

func TestApp(t *testing.T) {
	app := iris.New()
	app.RegisterView(iris.HTML("./views", ".html"))

	app.Get("/", func(ctx iris.Context) {
		ctx.ViewData("message", "Hello world!")
		ctx.View("index.html")
	})

	e := httptest.New(t, app)
	e.GET("/").Expect().
		Status(httptest.StatusOK).
		ContentType("text/html", "utf-8").
		Body().
		NotEmpty()
}
