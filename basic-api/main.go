package main

import (
	"basic_api/models"
	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()

	// ルーティング一覧
	booksAPI := app.Party("/books")
	{
		// ミドルウェアの使用
		booksAPI.Use(iris.Compression)

		// メソッド・URL・ハンドラー
		booksAPI.Get("/list", list)
		booksAPI.Post("/create", create)
	}

	app.Listen(":8080")
}

// 一覧表示
func list(ctx iris.Context) {
	books := []models.Book{
		{
			Title: "sample book",
			Body:  "This is a sample",
		},
		{
			Title: "Iris",
			Body:  "Iris is an amazing framework of Go!",
		},
		{
			Title: "How to get to Mars",
			Body: "I totally do not know.",
		},
	}

	ctx.JSON(books)
}

// 新規作成
func create(ctx iris.Context) {
	var b models.Book
	err := ctx.ReadJSON(&b)
	if err != nil {
		ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().Title("Book creation failure").DetailErr(err))
		return
	}

	println(b.Title + "を新規作成しました。")

	ctx.StatusCode(iris.StatusCreated)
}