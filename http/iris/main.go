package main

import (
	"log"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/kataras/iris/v12"
)

func main() {
	router := iris.Default()
	router.Get("/home", func(ctx iris.Context) {
		ctx.JSON(map[string]interface{}{
			"data": "hello iris",
		})
	})

	httpSrv := http.NewServer(http.Address(":8000"))
	httpSrv.HandlePrefix("/", router)

	app := kratos.New(
		kratos.Name("iris"),
		kratos.Server(
			httpSrv,
		),
	)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}

}
