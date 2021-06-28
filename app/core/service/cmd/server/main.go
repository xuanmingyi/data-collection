package main

import (
	"flag"
	"os"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/log"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	Name     = "core.service"
	Version  string
	flagconf string
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
}

func newApp(logger log.Logger) *kratos.App {
	return kratos.New(
		kratos.Name(Name),
		kratos.Version(Version),
	)
}

func main() {
	flag.Parse()
	logger := log.With(log.NewStdLogger(os.Stdout),
		"service.name", Name,
		"service.version", Version,
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller)

	c := config.New()

	if err := c.Load(); err != nil {
		panic(err)
	}

	app, cleanup, err := initApp(logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	if err := app.Run(); err != nil {
		panic(err)
	}

}
