package setup_controllers

import (
	"github.com/kataras/iris/v12/mvc"
	usersController "go-iris-sample/_example-mvc/controllers/users"
	usersDetailsController "go-iris-sample/_example-mvc/controllers/users/details"
)

// GET: "http://localhost:8080/users"
// PUT: "http://localhost:8080/users/details"
// GET: "http://localhost:8080/users/details/{title:string}"
func ConfigureUsersControllers(app *mvc.Application) {
	app.Handle(new(usersController.UsersController))
	app.Handle(new(usersDetailsController.UsersDetailsController))
}
