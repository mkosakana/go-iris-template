package users

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"go-iris-sample/_example-mvc-api/model"
	"go-iris-sample/_example-mvc-api/service"
)

type UsersController struct {
	Ctx iris.Context
}

var userService service.UserService

// メソッド名でパスの違いを受け付けます（超便利）
// GetList()       "http://localhost:8080/users/list"
// Post()          "http://localhost:8080/users"
// PutDetails()    "http://localhost:8080/users/details/{id:string}"
// DeleteDetails() "http://localhost:8080/users/details/{id:string}"
// [その他の例]
// POST: http://localhost:8080/users/details/example -> PostDetailsExample()
// PUT: http://localhost:8080/users -> Put()

func (c *UsersController) GetList() mvc.Response {
	// 一覧取得
	users, err := userService.GetUserList()
	if err != nil {
		fmt.Println(123)
		return mvc.Response{
			Code: iris.StatusInternalServerError,  // エラーハンドリング
		}
	}

	// Iris に備え付きのレスポンス用構造体（struct）
	return mvc.Response{
		Code:   iris.StatusOK,
		Object: users,
	}
}

func (c *UsersController) Post() mvc.Response {
	var user model.User

	// リクエストボディのjsonデータを構造体（struct）に格納する
	err := c.Ctx.ReadJSON(&user)

	// エラーハンドリング（Iris 備え付きのもので作れます）
	if err != nil {
		c.Ctx.StopWithError(iris.StatusBadRequest, err)
		return mvc.Response{
			Code: iris.StatusBadRequest,
		}
	}

	// 新規作成
	err = userService.CreateUser(&user)
	if err != nil {
		_, _ = c.Ctx.WriteString("ping")
		return mvc.Response{
			Code: iris.StatusInternalServerError,  // エラーハンドリング
		}
	}

	// Iris に備え付きのレスポンス用構造体（struct）
	return mvc.Response{ Code: iris.StatusCreated }
}

func (c *UsersController) PutDetails() mvc.Response {
	var user model.User

	// リクエストボディのjsonデータを構造体（struct）に格納する
	err := c.Ctx.ReadJSON(&user)

	// エラーハンドリング（Iris 備え付きのもので作れます）
	if err != nil {
		c.Ctx.StopWithError(iris.StatusBadRequest, err)
		return mvc.Response{
			Code: iris.StatusBadRequest,
		}
	}

	// URLからパラメーター("id")を取得
	id := c.Ctx.Params().Get("id")
	user.Id = id

	// 更新
	err = userService.UpdateUser(&user)
	if err != nil {
		return mvc.Response{
			Code: iris.StatusInternalServerError,  // エラーハンドリング
		}
	}

	// Iris に備え付きのレスポンス用構造体（struct）
	return mvc.Response{ Code: iris.StatusCreated }
}

func (c *UsersController) DeleteDetails() mvc.Response {
	// URLからパラメーター("id")を取得
	id := c.Ctx.Params().Get("id")

	// 削除
	err := userService.DeleteUser(id)
	if err != nil {
		return mvc.Response{
			Code: iris.StatusInternalServerError,  // エラーハンドリング
		}
	}

	// Iris に備え付きのレスポンス用構造体（struct）
	return mvc.Response{ Code: iris.StatusCreated }
}