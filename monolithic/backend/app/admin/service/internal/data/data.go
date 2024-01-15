package data

import (
	authnEngine "github.com/tx7do/kratos-authn/engine"
	"github.com/tx7do/kratos-authn/engine/jwt"

	authzEngine "github.com/tx7do/kratos-authz/engine"
	"github.com/tx7do/kratos-authz/engine/noop"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/redis/go-redis/v9"

	"github.com/tx7do/go-utils/entgo"
	"github.com/tx7do/kratos-bootstrap"

	"kratos-monolithic-demo/app/admin/service/internal/data/ent"

	conf "github.com/tx7do/kratos-bootstrap/gen/api/go/conf/v1"
)

// Data .
type Data struct {
	log *log.Helper

	rdb *redis.Client
	db  *entgo.EntClient[*ent.Client]

	authenticator authnEngine.Authenticator
	authorizer    authzEngine.Engine
}

// NewData .
func NewData(
	entClient *entgo.EntClient[*ent.Client],
	redisClient *redis.Client,
	authenticator authnEngine.Authenticator,
	authorizer authzEngine.Engine,
	logger log.Logger,
) (*Data, func(), error) {
	l := log.NewHelper(log.With(logger, "module", "data/admin-service"))

	d := &Data{
		db:            entClient,
		rdb:           redisClient,
		log:           l,
		authenticator: authenticator,
		authorizer:    authorizer,
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
func NewRedisClient(cfg *conf.Bootstrap, _ log.Logger) *redis.Client {
	//l := log.NewHelper(log.With(logger, "module", "redis/data/admin-service"))
	return bootstrap.NewRedisClient(cfg.Data)
}

// NewAuthenticator 创建认证器
func NewAuthenticator(cfg *conf.Bootstrap) authnEngine.Authenticator {
	authenticator, _ := jwt.NewAuthenticator(
		jwt.WithKey([]byte(cfg.Server.Rest.Middleware.Auth.Key)),
		jwt.WithSigningMethod(cfg.Server.Rest.Middleware.Auth.Method),
	)
	return authenticator
}

// NewAuthorizer 创建权鉴器
func NewAuthorizer() authzEngine.Engine {
	return noop.State{}
}
