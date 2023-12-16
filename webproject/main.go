package main

import (
	"github.com/go-zepto/zepto"
	"github.com/kyungmun/webproject/controllers"
)

func main() {
	// Create Zepto
	z := zepto.NewZepto()

	// Routes
	z.Get("/", controllers.HelloIndex)

	z.Start()
}
