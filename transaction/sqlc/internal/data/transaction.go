package data

import (
	"context"
	"github.com/go-kratos/examples/transaction/sqlc/internal/biz"
	"github.com/go-kratos/examples/transaction/sqlc/internal/data/queries"
)

type userRepo struct {
	data *Data
}

func NewUserRepo(data *Data) biz.UserRepo {
	return &userRepo{
		data: data,
	}
}

func (u *userRepo) CreateUser(ctx context.Context, name string) (int64, error) {
	result, err := u.data.DB(ctx).CreateUser(ctx, name)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func (u *userRepo) CreateUserDetail(ctx context.Context, id int64, email string) (int64, error) {
	result, err := u.data.DB(ctx).CreateUserDetail(ctx, queries.CreateUserDetailParams{
		ID:    id,
		Email: email,
	})
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}
