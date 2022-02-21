package main

import (
	"github.com/kataras/iris/v12"
	"testing"

	"github.com/kataras/iris/v12/httptest"
)

func TestApp(t *testing.T) {
	app := iris.New()
	app.Get("/", index)

	e := httptest.New(t, app)
	e.GET("/").Expect().Status(httptest.StatusOK).Body().Equal("<h1>Index Page</h1>")
}
