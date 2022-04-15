package animals

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/httptest"
	"github.com/kataras/iris/v12/mvc"
	"testing"
)

func TestAnimalsController_Get(t *testing.T) {
	app := iris.New()
	app.RegisterView(iris.HTML("../../public", ".html"))

	animals := app.Party("/animals")
	mvc.Configure(animals, func(mvcApp *mvc.Application) {
		mvcApp.Handle(new(AnimalsController))
	})

	e := httptest.New(t, app)
	e.GET("/animals").
		Expect().
		Status(httptest.StatusOK).
		ContentType("text/html", "utf-8").
		Body().
		NotEmpty()
}
