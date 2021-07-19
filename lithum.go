package main

import (
	"fmt"
	"lithum/conf"
)

func init() {

}

func main() {
	Config, err := conf.ReadConfig("configs.yaml")
	if err != nil {
		panic(err)
	}

	fmt.Println(Config)

	switch Config.Type {
	case "server":
		fmt.Println("server")

	case "client":
		fmt.Println("client")
	}
}
