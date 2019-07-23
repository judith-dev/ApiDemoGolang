package main

//import "fmt"

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

func main() {
	
	app := iris.New()
	app.Logger().SetLevel("debug")

	app.Use(recover.New())
	app.Use(logger.New())

	app.Handle("GET","/",func (context iris.Context)  {
		context.HTML("<h1>WELCOME<h1>")
	})
	app.Get("/ping", func (context  iris.Context){
		context.WriteString("ping gopher")
	})

	app.Run(iris.Addr(":8080"))
	// fmt.Printf("hello golang")
}