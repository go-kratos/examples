package data

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/go-kratos/examples/transaction/mongo/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

var _ biz.CardRepo = (*cardRepo)(nil)

type cardRepo struct {
	data *Data
	log  *log.Helper
}

type Card struct {
	Id        primitive.ObjectID `bson:"_id"`
	UserID    int64              `bson:"userID"`
	Money     int64              `bson:"money"`
	CreatedAt time.Time          `bson:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt"`
}

func NewCardRepo(data *Data, logger log.Logger) biz.CardRepo {
	return &cardRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (c *cardRepo) CreateCard(ctx context.Context, id int64) (string, error) {
	var card Card
	card.Id = primitive.NewObjectID()
	card.UserID = id
	card.Money = 1000
	_, err := c.data.db.Database("test").Collection("cards").InsertOne(ctx, card)
	if err != nil {
		return "", err
	}
	return card.Id.Hex(), nil
}
