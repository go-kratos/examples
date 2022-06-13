package service

import (
	"context"

	v1 "test/api/helloworld/v1"
	"test/internal/biz"
)

type GreeterService struct {
	v1.UnimplementedGreeterServer

	uc *biz.GreeterUsecase
}

func NewGreeterService(uc *biz.GreeterUsecase) *GreeterService {
	return &GreeterService{uc: uc}
}

func (s *GreeterService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	g, err := s.uc.SayHello(ctx, &biz.Greeter{Name: in.Name})
	if err != nil {
		return nil, err
	}
	return &v1.HelloReply{Message: "Hello " + g.Name}, nil
}
