package biz

import (
	"context"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewUserUsecase)

type Transaction interface {
	InTx(context.Context, func(ctx context.Context) error) error
}
