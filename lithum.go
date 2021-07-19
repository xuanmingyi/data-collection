package main

import (
	"fmt"
	"lithum/conf"
	"lithum/vpn"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"golang.org/x/net/context"
)

func main() {

	if err := conf.InitConfig("configs.yaml"); err != nil {
		panic(err)
	}

	context, cancel := context.WithCancel(context.Background())
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT)
	var wg sync.WaitGroup

	switch conf.Config.Type {
	case "server":
		wg.Add(1)
		vpnServer := vpn.NewVPNServer(context, &wg)
		go vpnServer.Run()

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
