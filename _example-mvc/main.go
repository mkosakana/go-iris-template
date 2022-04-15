package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	setUps "go-iris-sample/_example-mvc/setup-controllers"
)

func main() {
	app := iris.New()

	// ミドルウェアの使用
	app.Use(iris.Compression)

	// 静的ファイルの指定
	app.HandleDir("/", iris.Dir("./_example-mvc/public"))

	// "views" ディレクトリ配下の.htmlファイルを返却するように設定
	app.RegisterView(iris.HTML("_example-mvc/public", ".html"))

	// GET: "http://localhost:8080/"
	app.Handle("GET", "/ping", ping)

	// "http://localhost:8080/users/"から始まるパスを受け取る処理をグループ化
	users := app.Party("/users")
	mvc.Configure(users, setUps.ConfigureUsersControllers)

	// "http://localhost:8080/animals/"から始まるパスを受け取る処理をグループ化
	animals := app.Party("/animals")
	mvc.Configure(animals, setUps.ConfigureAnimalsControllers)

	// ポートの指定
	app.Listen(":8080")
}

// サーバーが動いているかどうかの確認
func ping(ctx iris.Context) {
	// "ping"が返る
	_, _ = ctx.WriteString("ping")
}
