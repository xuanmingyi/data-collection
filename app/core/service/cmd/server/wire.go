// +build wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/google/wire"
)

func initApp() (*kratos.App, func(), error) {
	panic(wire.Build(newApp))
}
