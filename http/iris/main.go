package main

import (
	"context"
	"fmt"
	"log"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"
	kiris "github.com/iris-contrib/kratos"
	"github.com/kataras/iris/v12"
)

func customMiddleware(handler middleware.Handler) middleware.Handler {
	return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
		if tr, ok := transport.FromServerContext(ctx); ok {
			fmt.Println("operation:", tr.Operation())
		}
		reply, err = handler(ctx, req)
		return
	}
}

func main() {
	router := iris.New()
	router.Use(kiris.Middlewares(recovery.Recovery(), customMiddleware))
	router.Get("/helloworld/{name}", func(ctx iris.Context) {
		name := ctx.Params().Get("name")
		if name == "error" {
			kiris.Error(ctx, errors.Unauthorized("auth_error", "no authentication"))
		} else {
			ctx.JSON(iris.Map{"welcome": name})
		}
	})

	if err := router.Build(); err != nil {
		log.Fatal(err)
	}

	httpSrv := http.NewServer(http.Address(":8000"))
	httpSrv.HandlePrefix("/", router)

	kratosApp := kratos.New(
		kratos.Name("iris"),
		kratos.Server(
			httpSrv,
		),
	)

	// http://localhost:8000/helloworld/kataras.
	if err := kratosApp.Run(); err != nil {
		log.Fatal(err)
	}
}
