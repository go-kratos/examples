package biz

import (
	"context"
	"fmt"
)

type UserRepo interface {
	CreateUser(ctx context.Context, name string) (int64, error)
	CreateUserDetail(ctx context.Context, id int64, phone string) (int64, error)
}

type UserUsecase struct {
	repo UserRepo
	tx   Transaction
}

func NewUserUsecase(repo UserRepo, tx Transaction) *UserUsecase {
	return &UserUsecase{
		repo: repo,
		tx:   tx,
	}
}

func (u *UserUsecase) CreateUser(ctx context.Context, name, email string) (int64, error) {
	var (
		id  int64
		err error
	)

	err = u.tx.InTx(ctx, func(ctx context.Context) error {
		id, err = u.repo.CreateUser(ctx, name)
		if err != nil {
			return err
		}
		_, err = u.repo.CreateUserDetail(ctx, id, email)
		return err
	})

	if err != nil {
		return 0, fmt.Errorf("user create fail %v", err)
	}

	return id, nil
}
