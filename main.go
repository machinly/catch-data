package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/kataras/iris"
)

func main() {
	app := iris.New()

	f, err := os.Create("/tmp/catch")
	if err != nil {
		panic(err)
	}
	app.Handle("POST", "/", func(ctx iris.Context) {
		bs, err := ioutil.ReadAll(ctx.Request().Body)
		if err != nil {
			fmt.Println(err)
		}
		f.Write(bs)
		f.Write([]byte("\n"))
		f.Sync()
		ctx.WriteString("got it")
	})

	app.Get("/ping", func(ctx iris.Context) {
		ctx.WriteString("pong")
	})

	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}
