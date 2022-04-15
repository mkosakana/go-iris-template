package main

import (
	"github.com/kataras/iris/v12"
	"go-iris-sample/_example-basic-api/models/books"
)

func main() {
	app := iris.New()

	// GET: "/"
	app.Handle("GET", "/", ping)

	// "/books/"から始まるパスを受け取る処理をグループ化
	books := app.Party("/books")
	{
		// GET: "/books/"
		books.Handle("GET", "/", getBook)
		// PUT: "/books/{id:string}"
		books.Handle("PUT", "/{id:string}", updateBook)
	}

	// ポートの指定
	app.Listen(":8080")
}

// サーバーが動いているかどうかの確認
func ping(ctx iris.Context) {
	// "ping"が返る
	_, _ = ctx.WriteString("ping")
}

// 本の取得メソッド
// 通常はDBと接続したりしてデータを取得・更新など行います
func getBook(ctx iris.Context) {
	books := []books.Book{
		{
			Title: "老人と海",
		},
		{
			Title: "コンビニ人間",
		},
		{
			Title: "歯車",
		},
	}

	ctx.JSON(books)
}

// 本の更新メソッド
// 入力例
// PUT: "https://localhost:8080/books/104"
// {
//		"title": "老人と森",
// }
//
// 返却例
// {
//		"id": "104",
//		"message": "老人と森	updated successfully"
// }
func updateBook(ctx iris.Context) {
	// URLからパラメーター("id")を取得
	id := ctx.Params().Get("id")

	var requestBody books.Request
	// リクエストボディのjsonデータを構造体（struct）に格納する
	err := ctx.ReadJSON(&requestBody)
	// エラーハンドリング
	if err != nil {
		ctx.StopWithError(iris.StatusBadRequest, err)
		return
	}

	// レスポンス用のjsonデータを構造体（struct）に格納する
	response := books.Response{
		ID:      id,
		Message: requestBody.Title + " updated successfully",
	}
	// jsonデータ（が格納された構造体）をレスポンスする
	ctx.JSON(response)
}
