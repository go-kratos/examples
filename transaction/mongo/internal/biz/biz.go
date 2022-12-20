package biz

import (
	"context"

	"github.com/google/wire"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewUserUseCase)

type Transaction interface {
	StartSession(ctx context.Context) (sessionCtx context.Context, endSession func(ctx context.Context), err error)
	ExecTx(sessionCtx context.Context, cmd func(ctx context.Context) error) error
}
