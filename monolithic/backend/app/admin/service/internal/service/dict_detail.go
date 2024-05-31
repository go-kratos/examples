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

type DictDetailService struct {
	adminV1.DictDetailServiceHTTPServer

	uc  *data.DictDetailRepo
	log *log.Helper
}

func NewDictDetailService(logger log.Logger, uc *data.DictDetailRepo) *DictDetailService {
	l := log.NewHelper(log.With(logger, "module", "dict-detail/service/admin-service"))
	return &DictDetailService{
		log: l,
		uc:  uc,
	}
}

func (s *DictDetailService) ListDictDetail(ctx context.Context, req *pagination.PagingRequest) (*systemV1.ListDictDetailResponse, error) {
	ret, err := s.uc.List(ctx, req)
	if err != nil {

		return nil, err
	}

	return ret, nil
}

func (s *DictDetailService) GetDictDetail(ctx context.Context, req *systemV1.GetDictDetailRequest) (*systemV1.DictDetail, error) {
	ret, err := s.uc.Get(ctx, req)
	if err != nil {

		return nil, err
	}

	return ret, nil
}

func (s *DictDetailService) CreateDictDetail(ctx context.Context, req *systemV1.CreateDictDetailRequest) (*emptypb.Empty, error) {
	err := s.uc.Create(ctx, req)
	if err != nil {
		// s.log.Info(err)
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *DictDetailService) UpdateDictDetail(ctx context.Context, req *systemV1.UpdateDictDetailRequest) (*emptypb.Empty, error) {
	err := s.uc.Update(ctx, req)
	if err != nil {

		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *DictDetailService) DeleteDictDetail(ctx context.Context, req *systemV1.DeleteDictDetailRequest) (*emptypb.Empty, error) {
	_, err := s.uc.Delete(ctx, req)
	if err != nil {

		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *DictDetailService) GetDictDetailByCode(ctx context.Context, req *systemV1.GetDictDetailRequest) (*systemV1.DictDetail, error) {
	ret, err := s.uc.Get(ctx, req)
	if err != nil {

		return nil, err
	}

	return ret, nil
}
