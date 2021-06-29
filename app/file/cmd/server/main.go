package main

import (
	"flag"
	"os"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/xuanmingyi/data-collection/app/file/internal/conf"
	"google.golang.org/grpc"
)

var (
	Name     string = "file"
	Version  string = "0.0.1"
	flagconf string
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
}

func newApp(gs *grpc.Server, logger log.Logger) *kratos.App {
	return kratos.New(
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Logger(logger),
	)
}

func main() {
	flag.Parse()

	logger := log.With(log.NewStdLogger(os.Stdout),
		"service.name", Name,
		"service.version", Version,
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller)

	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
	)

	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	app, cleanup, err := initApp(bc.Server, bc.Data, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	if err := app.Run(); err != nil {
		panic(err)
	}
}
