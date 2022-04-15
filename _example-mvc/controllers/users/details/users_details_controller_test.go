package details

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/httptest"
	"github.com/kataras/iris/v12/mvc"
	"testing"
)

func TestUsersDetailsController_PutDetails(t *testing.T) {
	app := iris.New()
	users := app.Party("/users")

	mvc.Configure(users, func(mvcApp *mvc.Application) {
		mvcApp.Handle(new(UsersDetailsController))
	})

	e := httptest.New(t, app)
	e.PUT("/users/details").WithJSON(
		iris.Map{
			"firstname": "Tanaka",
			"lastname": "ryo",
			"age": 15,
		}).
		Expect().
		Status(httptest.StatusCreated)
}
