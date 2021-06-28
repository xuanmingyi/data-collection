// +build wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"

	"gitee.com/tianyu-psychiatric-team/data-collection/app/core/service/internal/biz"
)

func initApp(log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(biz.ProviderSet, newApp))
}
