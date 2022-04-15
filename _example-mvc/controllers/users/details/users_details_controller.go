package details

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"go-iris-sample/_example-mvc/models/users"
)

type UsersDetailsController struct {
	Ctx iris.Context
}

// PutDetails "http://localhost:8080/users/details"
// メソッド名でパスの違いを受け付けます．（超便利）
// [例]
// Put: http://localhost:8080/users -> Put()
// Put: http://localhost:8080/users/details -> PutDetails()
func (c *UsersDetailsController) PutDetails() mvc.Response {
	var requestBody users.Request
	// リクエストボディのjsonデータを構造体（struct）に格納する
	err := c.Ctx.ReadJSON(&requestBody)
	// エラーハンドリング
	if err != nil {
		c.Ctx.StopWithError(iris.StatusBadRequest, err)
		return mvc.Response{
			Code: iris.StatusBadRequest,
		}
	}

	return mvc.Response{
		Code: iris.StatusCreated,
		Object: users.Response{
			Message: requestBody.Firstname + " created successfully",
		},
	}
}
