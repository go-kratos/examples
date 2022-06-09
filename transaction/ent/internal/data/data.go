package data

import (
	"context"
	"github.com/go-kratos/examples/transaction/ent/internal/biz"

	"github.com/go-kratos/examples/transaction/ent/internal/conf"
	"github.com/go-kratos/examples/transaction/ent/internal/data/ent"

	"entgo.io/ent/dialect/sql"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"

	// init mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewUserRepo, NewCardRepo, NewTransaction)

// Data .
type Data struct {
	db *ent.Database
}

// NewTransaction .
func NewTransaction(data *Data) biz.Transaction {
	return data.db
}

// NewData .
func NewData(conf *conf.Data, logger log.Logger) (*Data, func(), error) {
	log := log.NewHelper(logger)
	drv, err := sql.Open(
		conf.Database.Driver,
		conf.Database.Source,
	)
	if err != nil {
		log.Errorf("failed opening connection to sqlite: %v", err)
		return nil, nil, err
	}
	// Run the auto migration tool.
	client := ent.NewClient(ent.Driver(drv))
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Errorf("failed creating schema resources: %v", err)
		return nil, nil, err
	}

	d := &Data{
		db: ent.NewDatabase(ent.Driver(drv)),
	}

	return d, func() {
		log.Info("message", "closing the data resources")
		if err := drv.Close(); err != nil {
			log.Error(err)
		}
	}, nil
}
