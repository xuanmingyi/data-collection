// +build wireinject

package main

import (
	"github.com/xuanmingyi/data-collection/app/file/internal/biz"
	"github.com/xuanmingyi/data-collection/app/file/internal/conf"
	"github.com/xuanmingyi/data-collection/app/file/internal/data"
	"github.com/xuanmingyi/data-collection/app/file/internal/server"
	"github.com/xuanmingyi/data-collection/app/file/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"

	"github.com/google/wire"
)

func initApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
