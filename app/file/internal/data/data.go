package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"

	"github.com/xuanmingyi/data-collection/app/file/internal/conf"
	"github.com/xuanmingyi/data-collection/app/file/internal/data/ent"

	// init mysql driver
	_ "github.com/go-sql-driver/mysql"
)

var ProviderSet = wire.NewSet(NewData, NewFileRepo)

type Data struct {
	db *ent.Client
}

func NewData(conf *conf.Data, logger log.Logger) (*Data, func(), error) {
	log := log.NewHelper(log.With(logger, "module", "file/data"))

	client, err := ent.Open(
		conf.Database.Driver,
		conf.Database.Source,
	)

	if err != nil {
		log.Errorf("failed opening connection to sqlite: %v", err)
		return nil, nil, err
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Errorf("failed creating schema resources: %v", err)
		return nil, nil, err
	}

	d := &Data{db: client}
	return d, func() {
		if err := d.db.Close(); err != nil {
			log.Error(err)
		}
	}, nil
}
