package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type Greeter struct {
	Name string
}

type GreeterUsecase struct {
	log *log.Helper
}

func NewGreeterUsecase(logger log.Logger) *GreeterUsecase {
	return &GreeterUsecase{log: log.NewHelper(logger)}
}

func (uc *GreeterUsecase) SayHello(ctx context.Context, g *Greeter) (*Greeter, error) {
	return &Greeter{Name: g.Name}, nil
}
