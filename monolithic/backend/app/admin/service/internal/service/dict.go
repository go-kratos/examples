package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/emptypb"

	"kratos-monolithic-demo/app/admin/service/internal/data"

	adminV1 "kratos-monolithic-demo/gen/api/go/admin/service/v1"
	systemV1 "kratos-monolithic-demo/gen/api/go/system/service/v1"

	pagination "github.com/tx7do/kratos-bootstrap/gen/api/go/pagination/v1"
)

type DictService struct {
	adminV1.DictDetailServiceHTTPServer

	uc  *data.DictRepo
	log *log.Helper
}

func NewDictService(logger log.Logger, uc *data.DictRepo) *DictService {
	l := log.NewHelper(log.With(logger, "module", "dict/service/admin-service"))
	return &DictService{
		log: l,
		uc:  uc,
	}
}

func (s *DictService) ListDict(ctx context.Context, req *pagination.PagingRequest) (*systemV1.ListDictResponse, error) {
	ret, err := s.uc.List(ctx, req)
	if err != nil {

		return nil, err
	}

	return ret, nil
}

func (s *DictService) GetDict(ctx context.Context, req *systemV1.GetDictRequest) (*systemV1.Dict, error) {
	ret, err := s.uc.Get(ctx, req)
	if err != nil {

		return nil, err
	}

	return ret, nil
}

func (s *DictService) CreateDict(ctx context.Context, req *systemV1.CreateDictRequest) (*emptypb.Empty, error) {
	err := s.uc.Create(ctx, req)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *DictService) UpdateDict(ctx context.Context, req *systemV1.UpdateDictRequest) (*emptypb.Empty, error) {
	err := s.uc.Update(ctx, req)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *DictService) DeleteDict(ctx context.Context, req *systemV1.DeleteDictRequest) (*emptypb.Empty, error) {
	_, err := s.uc.Delete(ctx, req)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
