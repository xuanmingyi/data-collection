package data

import (
	"log"

	"github.com/google/wire"

	"go.mongodb.org/mongo-driver/mongo"
)

var ProviderSet = wire.NewSet(NewData)

type Data struct {
	db *mongo.Database
}

func NewData(conf *conf.Data, logger log.Logger) (*Data, func(), error) {
}
