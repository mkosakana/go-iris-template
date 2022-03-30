package main

import (
	"github.com/kataras/iris/v12"
)

func main()  {
	app := iris.New()

	// "_views" ディレクトリ配下の.htmlファイルを返却するように設定
	app.RegisterView(iris.HTML("./_views", ".html"))

	// GET - "127.0.0.1:8080/"
	app.Get("/", func(ctx iris.Context) {
		// ファイル内で使用する変数の設定
		ctx.ViewData("message", "Hello world!")
		// 返却するviewファイルの指定
		ctx.View("index.html")
	})
	
	// ポートの指定
	app.Listen(":8080")
}
