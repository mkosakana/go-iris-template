package setups

import (
	"log"

	"github.com/kataras/iris/v12/mvc"
	"github.com/mkosakana/go-iris-template/_example-mvc-api/configs/db"
)

func SetBaseConfiguration(app *mvc.Application) {
	// DB
	db, err := db.InitDb()
	if err != nil {
		log.Fatalf("func %s failed: %v", "InitDb", err)
	}
	app.Register(db)
}
