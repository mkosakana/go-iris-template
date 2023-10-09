package configs

import "github.com/kataras/iris/v12"

// SetPort port
func SetPort(app *iris.Application) error {
	return app.Listen(":8080")
}
