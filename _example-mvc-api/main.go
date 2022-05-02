package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/accesslog"
	"github.com/kataras/iris/v12/middleware/recover"
	"github.com/kataras/iris/v12/mvc"

	"go-iris-sample/_example-mvc-api/setups"
)

func main() {
	app := iris.New()

	// ミドルウェアの使用
	app.Use(iris.Compression)
	app.Configure(iris.WithoutBodyConsumptionOnUnmarshal)

	// ログ記録（これも備え付きミドルウェア）
	ac := accesslog.File("./access.log")
	defer ac.Close()
	app.UseRouter(ac.Handler)
	app.UseRouter(recover.New())

	// 静的ファイルの指定
	app.HandleDir("/", iris.Dir("./_example-mvc-api/public"))

	// GET: "http://localhost:8080/"
	app.Handle("GET", "/ping", ping)

	// "/users/"から始まるURLを受け取った際の処理をグループ化
	users := app.Party("/users")
	mvc.Configure(users, setups.ConfigureUsersControllers)

	// ポートの指定
	app.Listen(":8080")
}

// サーバーが動いているかどうかの確認
func ping(ctx iris.Context) {
	// "ping"が返る
	_, _ = ctx.WriteString("ping")
}
