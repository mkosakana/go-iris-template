package main

import (
	models "_example-basic-api/_models"
	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()

	// GET: "127.0.0.1:8080/"
	app.Handle("GET", "/", ping)
	// PUT: "127.0.0.1:8080/{id:string}"
	app.Handle("PUT", "/users/{id:string}", updateUser)

	// ポートの指定
	app.Listen(":8080")
}

// サーバーが動いているかどうかの確認
func ping(ctx iris.Context) {
	// "ping"が返る
	_, _ = ctx.WriteString("ping")
}

// ユーザーの更新メソッド
func updateUser(ctx iris.Context) {
	// URLからパラメーター("id")を取得
	id := ctx.Params().Get("id")

	var requestBody models.Request
	// リクエストボディのjsonデータを構造体（struct）に格納する
	err := ctx.ReadJSON(&requestBody)
	// エラーハンドリング
	if err != nil {
		ctx.StopWithError(iris.StatusBadRequest, err)
		return
	}

	// レスポンス用のjsonデータを構造体（struct）に格納する
	response := models.Response{
		ID:      id,
		Message: requestBody.Firstname + " updated successfully",
	}
	// jsonデータ（が格納された構造体）をレスポンスする
	ctx.JSON(response)
}

// 入力例
// PUT: "127.0.0.1:8080/104"
// {
// 	"firstname": "Taro",
//	"Lastname": "Yamada"
// }
//
// 返却例
// {
// 	"id": "104",
//	"message": "Taro updated successfully"
// }
