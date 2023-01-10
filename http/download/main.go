package main

import (
	"fmt"
	"log"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/xuri/excelize/v2"
	"github.com/gorilla/handlers"

)

func downloadFile(ctx http.Context) error {
	f := excelize.NewFile()
	index := f.NewSheet("Sheet2")
	f.SetCellValue("Sheet2", "A2", "Hello world.")
	f.SetCellValue("Sheet1", "B2", 100)
	f.SetActiveSheet(index)
	disposition := fmt.Sprintf("attachment; filename=%s", "demo.xlsx")
	ctx.Response().Header().Set("Content-Type", "application/vnd.openxmlformats")
	ctx.Response().Header().Set("Content-Disposition", disposition)
	ctx.Response().Header().Set("Access-Control-Expose-Headers", "Content-Disposition")

	return f.Write(ctx.Response())
}

func main() {
	var opts = []http.ServerOption{
		http.Address(":8001"),
		http.Filter(
			handlers.CORS(
				handlers.AllowedOrigins([]string{"*"}),
			),
		),
	}

	httpSrv := http.NewServer(
		opts...,
	)
	route := httpSrv.Route("/")
	route.POST("/download", downloadFile)

	app := kratos.New(
		kratos.Name("download"),
		kratos.Server(
			httpSrv,
		),
	)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
