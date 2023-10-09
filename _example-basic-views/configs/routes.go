package configs

import (
	"github.com/kataras/iris/v12"
)

// SetRoutesOfViews routes of views
func SetRoutesOfViews(app *iris.Application) error {
	// pages ディレクトリ配下のHTMLファイルを返却するように設定する
	app.RegisterView(iris.HTML("./pages", ".html"))

	app.Get("/", func(ctx iris.Context) {
		ctx.ViewData("message", "Hello world!!")
		err := ctx.View("index.html")
		if err != nil {
			return
		}
	})

	// "/books"から始まるパスを受け取る処理をグループ化
	books := app.Party("/books")
	{
		// GET: /books
		books.Get("/", func(ctx iris.Context) {
			err := ctx.View("/books/list.html")
			if err != nil {
				return
			}
		})
		// GET: /books/purchase
		books.Get("/purchase", func(ctx iris.Context) {
			err := ctx.View("/books/purchase.html")
			if err != nil {
				return
			}
		})
	}

	return nil
}
