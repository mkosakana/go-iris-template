package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/httptest"
	"testing"
)

func TestApp(t *testing.T) {
	app := iris.New()
	app.Handle("PUT", "/users/{id:string}", updateUser)

	e := httptest.New(t, app)
	e.PUT("/users/123").WithJSON(
		iris.Map{
			"firstname": "Taro",
			"Lastname": "Yamada",
		},
		).Expect().
		Status(httptest.StatusOK)
}
