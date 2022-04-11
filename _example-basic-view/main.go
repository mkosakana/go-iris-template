package main

import (
	"github.com/kataras/iris/v12"
)

func main()  {
	app := iris.New()

	// "views" ディレクトリ配下の.htmlファイルを返却するように設定
	// 上書きされバグが潜む可能性があるため基本的には一度だけ設定するようにし，
	// 指定したディレクトリ（今回は"views"）配下にディレクトリを増やしていくようにする
	app.RegisterView(iris.HTML("_example-basic-view/views", ".html"))

	// GET: "/"
	app.Get("/", func(ctx iris.Context) {
		// ファイル内で使用する変数の設定
		ctx.ViewData("message", "Hello world!")
		// 返却するviewファイルの指定
		ctx.View("index.html")
	})

	// "/books/"から始まるパスを受け取る処理をグループ化
	books := app.Party("/books")
	{
		// GET: "/books/"
		books.Get("/", func(ctx iris.Context) {
			ctx.View("/books/list.html")
		})
		// GET: "/books/buy"
		books.Get("/buy", func(ctx iris.Context) {
			ctx.View("/books/purchase.html")
		})
	}
	
	// ポートの指定
	app.Listen(":8080")
}
