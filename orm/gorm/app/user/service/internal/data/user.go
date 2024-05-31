package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/kratos-utils/crypto"
	util "github.com/tx7do/kratos-utils/time"

	"kratos-gorm-example/app/user/service/internal/biz"
	"kratos-gorm-example/app/user/service/internal/data/models"

	"kratos-gorm-example/gen/api/go/common/pagination"
	v1 "kratos-gorm-example/gen/api/go/user/service/v1"
)

var _ biz.UserRepo = (*UserRepo)(nil)

type UserRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	l := log.NewHelper(log.With(logger, "module", "user/repo/user-service"))
	return &UserRepo{
		data: data,
		log:  l,
	}
}

func (r *UserRepo) convertModelToProto(in *models.User) *v1.User {
	if in == nil {
		return nil
	}
	return &v1.User{
		Id:         uint32(in.ID),
		UserName:   &in.UserName,
		NickName:   &in.NickName,
		Password:   &in.Password,
		CreateTime: util.TimeToTimeString(&in.CreatedAt),
		UpdateTime: util.TimeToTimeString(&in.UpdatedAt),
	}
}

func (r *UserRepo) List(_ context.Context, req *pagination.PagingRequest) (*v1.ListUserResponse, error) {
	var results []models.User

	result := r.data.db.
		Limit(int(req.GetPageSize())).
		Offset(int(req.GetPageSize() * (req.GetPage() - 1))).
		Find(&results)
	if result.Error != nil {
		return nil, result.Error
	}

	items := make([]*v1.User, 0, len(results))
	for _, res := range results {
		item := r.convertModelToProto(&res)
		items = append(items, item)
	}

	var count int64
	result = r.data.db.Model(&models.User{}).
		Count(&count)
	if result.Error != nil {
		return nil, result.Error
	}

	return &v1.ListUserResponse{
		Total: int32(count),
		Items: items,
	}, nil
}

func (r *UserRepo) Get(_ context.Context, req *v1.GetUserRequest) (*v1.User, error) {
	res := &models.User{}
	r.data.db.First(res, "id = ?", req.GetId())
	return r.convertModelToProto(res), nil
}

func (r *UserRepo) Create(_ context.Context, req *v1.CreateUserRequest) (*v1.User, error) {
	cryptoPassword, err := crypto.HashPassword(req.User.GetPassword())
	if err != nil {
		return nil, err
	}

	res := &models.User{
		UserName: req.User.GetUserName(),
		NickName: req.User.GetNickName(),
		Password: cryptoPassword,
	}

	result := r.data.db.Create(res)
	if result.Error != nil {
		return nil, result.Error
	}

	return r.convertModelToProto(res), err
}

func (r *UserRepo) Update(_ context.Context, req *v1.UpdateUserRequest) (*v1.User, error) {
	var cryptoPassword string
	var err error
	if req.User.Password != nil {
		cryptoPassword, err = crypto.HashPassword(req.User.GetPassword())
		if err != nil {
			return nil, err
		}
	}

	res := &models.User{
		UserName: req.User.GetUserName(),
		NickName: req.User.GetNickName(),
		Password: cryptoPassword,
	}

	result := r.data.db.Model(res).Updates(res)
	if result.Error != nil {
		return nil, result.Error
	}

	return r.convertModelToProto(res), err
}

func (r *UserRepo) Delete(_ context.Context, req *v1.DeleteUserRequest) (bool, error) {
	result := r.data.db.Delete(&models.User{}, req.GetId())
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}
