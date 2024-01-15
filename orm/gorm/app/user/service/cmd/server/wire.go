//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/google/wire"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"

	"kratos-gorm-example/app/user/service/internal/biz"
	"kratos-gorm-example/app/user/service/internal/data"
	"kratos-gorm-example/app/user/service/internal/server"
	"kratos-gorm-example/app/user/service/internal/service"
	"kratos-gorm-example/gen/api/go/common/conf"
)

// initApp init kratos application.
func initApp(log.Logger, registry.Registrar, *conf.Bootstrap) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
