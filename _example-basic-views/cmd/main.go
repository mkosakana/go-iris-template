package main

import (
	"log"

	"github.com/kataras/iris/v12"
	"github.com/mkosakana/go-iris-sample/_example-basic-views/configs"
)

func main() {
	app := iris.New()

	// routes
	err := configs.SetRoutesOfViews(app)
	if err != nil {
		log.Fatalf("func %s failed: %v", "SetRoutes", err)
	}

	// port
	err = configs.SetPort(app)
	if err != nil {
		log.Fatalf("func %s failed: %v", "SetPort", err)
	}
}
