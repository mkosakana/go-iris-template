package setup_controllers

import (
	"github.com/kataras/iris/v12/mvc"
	"go-iris-sample/_example-mvc/controllers/animals"
)

func ConfigureAnimalsControllers(app *mvc.Application) {
	app.Handle(new(animals.AnimalsController))
}
