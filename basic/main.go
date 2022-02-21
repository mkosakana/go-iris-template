package main

import "github.com/kataras/iris/v12"

func main()  {
	app := iris.New()

	// GET - "127.0.0.1:8080/index"
	app.Get("/", index)

	// 接続ポートの指定
	app.Listen(":8080")
}

func index(ctx iris.Context) {
	ctx.HTML("<h1>Index Page</h1>")
}