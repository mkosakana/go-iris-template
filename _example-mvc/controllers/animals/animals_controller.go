package animals

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type AnimalsController struct {
	Ctx iris.Context
}

// Get "http://localhost:8080/animals"
func (c *AnimalsController) Get() mvc.View {
	var animalsListView = mvc.View{
		Name: "/animals/list.html",
		Data: iris.Map{"message": "🐭🐮🐯🐰🐲🐍🐴🐑🐵🐔🐶🐗"},
	}

	return animalsListView
}