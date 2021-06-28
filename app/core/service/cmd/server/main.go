package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
)

var (
	Name    = "core.service"
	Version string
)

func newApp() *kratos.App {
	return kratos.New(
		kratos.Name(Name),
		kratos.Version(Version),
	)
}

func main() {
	c := config.New()

	if err := c.Load(); err != nil {
		panic(err)
	}

	app, cleanup, err := initApp()
	if err != nil {
		panic(err)
	}
	defer cleanup()

	if err := app.Run(); err != nil {
		panic(err)
	}

}
