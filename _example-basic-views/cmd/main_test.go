package main

import (
	"os"
	"testing"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/httptest"
	"github.com/stretchr/testify/assert"
)

func Test_main(t *testing.T) {
	app := iris.New()

	// ※ Dockerコンテナ内でテストを実行する際は以下のコメントアウトを外し、その下の行と据え変えてください。
	// err := configs.SetRoutesOfViews(app)
	err := setRoutesOfViews(app)
	assert.NoError(t, err)

	e := httptest.New(t, app)
	e.GET("/").
		Expect().
		Status(httptest.StatusOK).
		HasContentType("text/html", "utf-8").
		Body().
		NotEmpty()
	e.GET("/books").
		Expect().
		Status(httptest.StatusOK).
		HasContentType("text/html", "utf-8").
		Body().
		NotEmpty()
	e.GET("/books/purchase").
		Expect().
		Status(httptest.StatusOK).
		HasContentType("text/html", "utf-8").
		Body().
		NotEmpty()
}

func setRoutesOfViews(app *iris.Application) error {
	currentDirName, err := os.Getwd()
	if err != nil {
		return err
	}
	app.RegisterView(iris.HTML(currentDirName+"/../pages", ".html"))

	app.Get("/", func(ctx iris.Context) {
		ctx.ViewData("message", "Hello world!")
		err = ctx.View("index.html")
		if err != nil {
			return
		}
	})

	books := app.Party("/books")
	{
		books.Get("/", func(ctx iris.Context) {
			err = ctx.View("/books/list.html")
			if err != nil {
				return
			}
		})
		books.Get("/purchase", func(ctx iris.Context) {
			err = ctx.View("/books/purchase.html")
			if err != nil {
				return
			}
		})
	}

	return nil
}
