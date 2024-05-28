// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"context"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"otel/internal/biz"
	"otel/internal/conf"
	"otel/internal/data"
	"otel/internal/dep"
	"otel/internal/server"
	"otel/internal/service"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(contextContext context.Context, bootstrap *conf.Bootstrap, confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	textMapPropagator := dep.NewTextMapPropagator()
	tracerProvider, err := dep.NewTracerProvider(contextContext, bootstrap, textMapPropagator)
	if err != nil {
		return nil, nil, err
	}
	meterProvider, err := dep.NewMeterProvider(bootstrap)
	if err != nil {
		return nil, nil, err
	}
	db, cleanup, err := dep.NewDB(confData, tracerProvider, meterProvider)
	if err != nil {
		return nil, nil, err
	}
	client, err := dep.NewRedis(confData, tracerProvider, meterProvider)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	dataData, err := data.NewData(db, client)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	asyncProducer, cleanup2, err := dep.NewProducer(bootstrap, logger, tracerProvider, textMapPropagator)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	consumerGroup, cleanup3, err := dep.NewConsumerGroup(bootstrap)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	greeterRepo := data.NewGreeterRepo(contextContext, dataData, logger, asyncProducer, consumerGroup, tracerProvider, textMapPropagator)
	greeterUsecase := biz.NewGreeterUsecase(greeterRepo, logger)
	greeterService := service.NewGreeterService(greeterUsecase)
	grpcServer := server.NewGRPCServer(confServer, greeterService, logger)
	meter, err := dep.NewMeter(bootstrap, meterProvider)
	if err != nil {
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	httpServer, err := server.NewHTTPServer(confServer, greeterService, logger, meter, tracerProvider)
	if err != nil {
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	pprofServer, err := server.NewPprof(confServer)
	if err != nil {
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	app := newApp(logger, grpcServer, httpServer, pprofServer)
	return app, func() {
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}