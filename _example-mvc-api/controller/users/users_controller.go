package users

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"go-iris-sample/_example-mvc-api/model"
	"go-iris-sample/_example-mvc-api/service"
)

// UsersController ユーザーコントローラー
type UsersController struct {
	UserService service.UserService
	Ctx         iris.Context
}

// メソッド名でパスの違い・リクエストメソッド・パラメータを受け付ける
// [例]
// GET: "http://localhost:8080/users/list"                → GetList()
// POST: "http://localhost:8080/users"                    → Post()
// PUT: "http://localhost:8080/users/details"             → PutDetails()
// PUT: "http://localhost:8080/users/details/{id:int}"    → PutDetailsBy(id uint)
// DELETE: "http://localhost:8080/users/details"          → DeleteDetails()
// DELETE: "http://localhost:8080/users/details/{id:int}" → DeleteDetailsBy(id uint)
// [その他の例]
// POST: http://localhost:8080/users/details/example -> PostDetailsExample()
// PUT: http://localhost:8080/users                  -> Put()

// GetList ユーザー一覧取得
func (c UsersController) GetList() mvc.Response {
	users, err := c.UserService.GetUserList()
	if err != nil {
		return mvc.Response{
			Code: iris.StatusInternalServerError,
		}
	}

	return mvc.Response{
		Code:   iris.StatusOK,
		Object: users,
	}
}

// Post ユーザー追加
func (c UsersController) Post() mvc.Response {
	// リクエストボディのjsonデータを構造体（struct）に格納する
	var user model.User
	err := c.Ctx.ReadJSON(&user)

	if err != nil {
		c.Ctx.StopWithError(iris.StatusBadRequest, err)
		return mvc.Response{
			Code: iris.StatusBadRequest,
		}
	}

	err = c.UserService.CreateUser(&user)
	if err != nil {
		return mvc.Response{
			Code: iris.StatusInternalServerError,
		}
	}

	return mvc.Response{Code: iris.StatusCreated}
}

// PutDetailsBy ユーザー更新
func (c UsersController) PutDetailsBy(id int) mvc.Response {
	// リクエストボディのjsonデータを構造体（struct）に格納する
	var user model.User
	err := c.Ctx.ReadJSON(&user)

	if err != nil {
		c.Ctx.StopWithError(iris.StatusBadRequest, err)
		return mvc.Response{
			Code: iris.StatusBadRequest,
		}
	}

	err = c.UserService.UpdateUser(id, &user)
	if err != nil {
		return mvc.Response{
			Code: iris.StatusInternalServerError,
		}
	}

	return mvc.Response{Code: iris.StatusOK}
}

// DeleteDetailsBy ユーザー削除
func (c UsersController) DeleteDetailsBy(id int) mvc.Response {
	err := c.UserService.DeleteUser(id)
	if err != nil {
		return mvc.Response{
			Code: iris.StatusInternalServerError,
		}
	}

	return mvc.Response{Code: iris.StatusOK}
}
