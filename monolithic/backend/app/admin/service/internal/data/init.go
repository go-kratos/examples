//go:build wireinject
// +build wireinject

package data

import "github.com/google/wire"

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,

	NewRedisClient,
	NewEntClient,

	NewAuthenticator,
	NewAuthorizer,

	NewDictRepo,
	NewDictDetailRepo,
	NewMenuRepo,
	NewOrganizationRepo,
	NewPositionRepo,
	NewRoleRepo,
	NewUserRepo,
	NewUserTokenRepo,
)
