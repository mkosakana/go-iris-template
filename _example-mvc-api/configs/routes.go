package configs

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/mkosakana/go-iris-template/_example-mvc-api/setups"
)

// SetRoutes routes
func SetRoutes(app *iris.Application) error {
	// health check
	app.Handle("GET", "/ping", healthCheckOfServer)

	// "/books"から始まるパスを受け取る処理をグループ化
	books := app.Party("/books")
	mvc.Configure(books, setups.SetBooksController)

	return nil
}

func healthCheckOfServer(ctx iris.Context) {
	_, _ = ctx.WriteString("ping")
}
