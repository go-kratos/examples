package main

import (
	"log"

	"github.com/go-kratos/examples/sse/handler"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/sse", handler.SSEHandler)

	httpSrv := http.NewServer(http.Address(":8080"), http.Timeout(0))
	httpSrv.HandlePrefix("/", router)

	app := kratos.New(
		kratos.Name("sse"),
		kratos.Server(
			httpSrv,
		),
	)

	log.Println("Open http://127.0.0.1:8080/sse in your web browser")
	if err := app.Run(); err != nil {
		log.Println(err)
	}
}
