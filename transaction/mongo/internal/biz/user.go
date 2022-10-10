package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type User struct {
	Name  string
	Email string
}

type UserRepo interface {
	CreateUser(ctx context.Context, a *User) (int64, error)
}

type UserUseCase struct {
	userRepo UserRepo
	cardRepo CardRepo
	txn      Transaction

	log *log.Helper
}

func NewUserUseCase(user UserRepo, card CardRepo, txn Transaction, logger log.Logger) *UserUseCase {
	return &UserUseCase{
		userRepo: user,
		cardRepo: card,
		txn:      txn,

		log: log.NewHelper(log.With(logger, "module", "usercase/interface")),
	}
}

func (u *UserUseCase) CreateUser(ctx context.Context, user *User) (int, error) {
	var (
		err error
		id  int64
	)

	sessionCtx, end, err := u.txn.StartSession(ctx)
	if err != nil {
		return 0, err
	}
	defer end(context.Background())

	if err := u.txn.ExecTx(sessionCtx, func(ctx context.Context) error {
		id, err = u.userRepo.CreateUser(ctx, user)
		if err != nil {
			return err
		}

		_, err = u.cardRepo.CreateCard(ctx, id)
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return 0, err
	}
	return int(id), nil
}
