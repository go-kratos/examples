package data

import (
	"context"
	"database/sql"
	"github.com/go-kratos/examples/transaction/sqlc/internal/biz"
	"github.com/go-kratos/examples/transaction/sqlc/internal/conf"
	"github.com/go-kratos/examples/transaction/sqlc/internal/data/queries"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	// init mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewDB, NewTransaction, NewUserRepo)

// Data .
type Data struct {
	db  *sql.DB
	log *log.Helper
}

type contextTxKey struct{}

func (d *Data) InTx(ctx context.Context, fn func(ctx context.Context) error) error {

	tx, err := d.db.Begin()
	if err != nil {
		return err
	}
	defer func() { _ = tx.Rollback() }()

	err = fn(context.WithValue(ctx, contextTxKey{}, queries.New(tx)))

	if err != nil {
		return err
	}

	return tx.Commit()
}

func (d *Data) DB(ctx context.Context) *queries.Queries {
	tx, ok := ctx.Value(contextTxKey{}).(*queries.Queries)
	if ok {
		return tx
	}
	return queries.New(d.db)
}

// NewTransaction .
func NewTransaction(d *Data) biz.Transaction {
	return d
}

// NewData .
func NewData(db *sql.DB, logger log.Logger) (*Data, func(), error) {
	l := log.NewHelper(log.With(logger, "module", "transaction/data"))
	d := &Data{
		db:  db,
		log: l,
	}
	return d, func() {
	}, nil
}

// NewDB sql Connecting to a Database
func NewDB(conf *conf.Data, logger log.Logger) (*sql.DB, func()) {
	l := log.NewHelper(log.With(logger, "module", "order-service/data/sqlc"))
	db, err := sql.Open("mysql", conf.Database.Source)
	if err != nil {
		l.Fatal(err)
	}
	cleanup := func() {
		_ = db.Close()
	}
	return db, cleanup
}
