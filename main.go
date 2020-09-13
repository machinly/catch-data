package main

import (
	"fmt"
	"io/ioutil"

	"github.com/kataras/iris"
)

func main() {
	app := iris.New()

	app.Handle("POST", "/", func(ctx iris.Context) {
		bs, err := ioutil.ReadAll(ctx.Request().Body)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(bs))
		ctx.WriteString("got it")
	})

	app.Get("/ping", func(ctx iris.Context) {
		ctx.WriteString("pong")
	})

	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}
