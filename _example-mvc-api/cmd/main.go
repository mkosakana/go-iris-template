// Write your go code in the editor on the left and watch it previewed here on the right.
package main

import (
	"log"

	"github.com/kataras/iris/v12"
	"github.com/mkosakana/go-iris-template/_example-mvc-api/configs"
)

// Test
type Test int

// TestFn
func TestFn() int {
	return 1
}

func main() {
	app := iris.New()

	// middleware
	app.Use(iris.Compression)
	app.Configure(iris.WithoutBodyConsumptionOnUnmarshal)

	// routes
	err := configs.SetRoutes(app)
	if err != nil {
		log.Fatalf("func %s failed: %v", "SetRoutes", err)
	}

	// port
	err = configs.SetPort(app)
	if err != nil {
		log.Fatalf("func %s failed: %v", "SetPort", err)
	}
}
