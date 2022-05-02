package users

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/httptest"
	"github.com/kataras/iris/v12/mvc"
	"testing"
)

func TestUsersController_Get(t *testing.T) {
	app := iris.New()
	users := app.Party("/users")

	mvc.Configure(users, func(mvcApp *mvc.Application) {
		mvcApp.Handle(new(UsersController))
	})

	e := httptest.New(t, app)
	e.GET("/users").
		Expect().
		Status(httptest.StatusOK)
}

func TestUsersController_GetExample(t *testing.T) {
	app := iris.New()
	users := app.Party("/users")

	mvc.Configure(users, func(mvcApp *mvc.Application) {
		mvcApp.Handle(new(UsersController))
	})

	e := httptest.New(t, app)
	e.GET("/users/example").
		Expect().
		Status(httptest.StatusOK).
		Body().
		Equal("example")
}