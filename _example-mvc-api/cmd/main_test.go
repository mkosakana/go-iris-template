package main

import (
	"testing"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/httptest"
)

func Test_main(t *testing.T) {
	app := iris.New()
	app.Handle("GET", "/ping", healthCheckOfServer)

	e := httptest.New(t, app)
	e.GET("/ping").
		Expect().
		Status(httptest.StatusOK).
		Body().
		Contains("ping")
}

func healthCheckOfServer(ctx iris.Context) {
	_, _ = ctx.WriteString("ping")
}
