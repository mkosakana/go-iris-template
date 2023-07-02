package setups

import (
	"github.com/kataras/iris/v12/middleware/accesslog"
	"github.com/kataras/iris/v12/mvc"
	"go-iris-sample/_example-mvc-api/controller/users"
)

// ConfigureUsersControllers
// main.go で定義したURLのグループ（ここではusers）に対してDI（依存性注入）を行っている
func ConfigureUsersControllers(app *mvc.Application) {
	// ログを取得してくれる機能のDI
	app.Register(accesslog.GetFields)

	// URLが "/users" から始まるリクエストを受け取った際に，
	// 以下の Controllerを使用させるという指示
	app.Handle(new(users.UsersController))
}
