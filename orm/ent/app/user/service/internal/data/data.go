package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/tx7do/kratos-utils/entgo"

	"kratos-ent-example/app/user/service/internal/data/ent"
	"kratos-ent-example/gen/api/go/common/conf"
	"kratos-ent-example/pkg/bootstrap"
)

// Data .
type Data struct {
	log *log.Helper
	db  *entgo.EntClient[*ent.Client]
	rdb *redis.Client
}

// NewData .
func NewData(entClient *entgo.EntClient[*ent.Client], redisClient *redis.Client, logger log.Logger) (*Data, func(), error) {
	l := log.NewHelper(log.With(logger, "module", "data/user-service"))

	d := &Data{
		db:  entClient,
		rdb: redisClient,
		log: l,
	}

	return d, func() {
		l.Info("message", "closing the data resources")
		d.db.Close()
		if err := d.rdb.Close(); err != nil {
			l.Error(err)
		}
	}, nil
}

// NewRedisClient 创建Redis客户端
func NewRedisClient(cfg *conf.Bootstrap, logger log.Logger) *redis.Client {
	l := log.NewHelper(log.With(logger, "module", "redis/data/user-service"))
	return bootstrap.NewRedisClient(cfg, l)
}
