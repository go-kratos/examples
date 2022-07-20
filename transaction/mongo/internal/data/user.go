package data

import (
	"context"
	"math/rand"
	"time"

	"github.com/go-kratos/examples/transaction/mongo/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

var _ biz.UserRepo = (*userRepo)(nil)

type userRepo struct {
	data *Data
	log  *log.Helper
}

type User struct {
	Id        int64     `bson:"id"`
	Name      string    `bson:"name"`
	Email     string    `bson:"email"`
	CreatedAt time.Time `bson:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt"`
}

// NewUserRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (u *userRepo) CreateUser(ctx context.Context, m *biz.User) (int64, error) {
	user := User{Id: rand.Int63n(1000000), Name: m.Name, Email: m.Email}

	_, err := u.data.db.Database("test").Collection("users").InsertOne(ctx, user)
	if err != nil {
		return 0, err
	}
	return user.Id, nil
}
