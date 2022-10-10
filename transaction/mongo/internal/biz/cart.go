package biz

import "context"

type CardRepo interface {
	CreateCard(ctx context.Context, id int64) (string, error)
}
