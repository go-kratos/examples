package data

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/go-kratos/examples/transaction/mongo/internal/biz"
	"github.com/go-kratos/examples/transaction/mongo/internal/conf"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewMongo, NewTransaction, NewUserRepo, NewCardRepo)

// Data .
type Data struct {
	db  *mongo.Client
	log *log.Helper
}

type contextTxKey struct{}

func (d *Data) StartSession(ctx context.Context) (context.Context, func(context.Context), error) {
	session, err := d.db.StartSession()
	if err != nil {
		return nil, nil, err
	}
	return context.WithValue(ctx, contextTxKey{}, session), session.EndSession, err
}

func (d *Data) ExecTx(ctx context.Context, fn func(ctx context.Context) error) error {
	sessionCtx, ok := ctx.Value(contextTxKey{}).(mongo.Session)
	if !ok {
		return fn(ctx)
	}

	_, err := sessionCtx.WithTransaction(ctx, func(sessCtx mongo.SessionContext) (interface{}, error) {
		return nil, fn(sessCtx)
	})
	return err
}

// NewTransaction .
func NewTransaction(d *Data) biz.Transaction {
	return d
}

// NewData .
func NewData(mongo *mongo.Client, logger log.Logger) (*Data, func(), error) {
	l := log.NewHelper(log.With(logger, "module", "transaction/data"))
	d := &Data{
		db:  mongo,
		log: l,
	}
	return d, func() {
	}, nil
}

func NewMongo(conf *conf.Data, logger log.Logger) *mongo.Client {
	log := log.NewHelper(log.With(logger, "module", "user/data/mongo"))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(conf.Mongo.Url))
	if err != nil {
		log.Fatalf("failed opening connection to mongo: %v", err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	return client
}
