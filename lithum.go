package main

import (
	"flag"
	"fmt"
	"lithum/conf"
	"net"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var path string

func init() {
	flag.StringVar(&path, "conf", "configs.yaml", "服务器配置文件")

}

func main() {
	flag.Parse()

	if strings.HasPrefix(path, "grpc://") {
		// 远程配置
		if err := conf.InitRPCConfig(path); err != nil {
			panic(err)
		}
	} else {
		// 配置文件
		if err := conf.InitConfig(path); err != nil {
			panic(err)
		}
	}

	context, cancel := context.WithCancel(context.Background())
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT)
	var wg sync.WaitGroup

	switch conf.Config.Type {
	case "server":
		server := grpc.NewServer()

		// 配置服务器
		conf.RegisterConfServer(server)

		lis, err := net.Listen("tcp",
			fmt.Sprintf("%s:%d", conf.Config.Server.Listen, conf.Config.Server.Port))
		if err != nil {
			panic(err)
		}
		if err := server.Serve(lis); err != nil {
			panic(err)
		}

		fmt.Println(context)
		//wg.Add(1)
		//vpnServer := vpn.NewVPNServer(context, &wg)
		//go vpnServer.Run()

	case "client":
		fmt.Println("client")
	}

	switch <-c {
	case syscall.SIGINT:
		fmt.Println("cancel the main gorouting")
		cancel()
	}

	wg.Wait()

	fmt.Println("all gorouting exit")
}
