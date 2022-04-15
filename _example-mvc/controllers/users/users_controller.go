package users

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"go-iris-sample/_example-mvc/models/users"
)

type UsersController struct {
	Ctx iris.Context
}

// Get "http://localhost:8080/users"
// メソッド名でパスの違いを受け付けます（超便利）
// （↑）したがって，どのパスが来るのかわかりやすいようにディレクトリ分けがおすすめです．
// [例]
// GET: http://localhost:8080/users -> Get()
// GET: http://localhost:8080/users/example -> GetExample()
func (c *UsersController) Get() mvc.Response {
	response := []users.User{
		{
			Name: "Tom",
			Age:  16,
		},
		{
			Name: "Bob",
			Age:  13,
		},
		{
			Name: "Taro",
			Age:  34,
		},
		{
			Name: "Hanako",
			Age:  28,
		},
	}

	return mvc.Response{
		Code:   iris.StatusOK,
		Object: response,
	}
}

func (c *UsersController) GetExample() {
	// "example"が返る
	_, _ = c.Ctx.WriteString("example")
}