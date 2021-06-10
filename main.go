package main

import (
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"

	"context"

	"github.com/go-kratos/kratos/v2/registry"

	etcd_registry "github.com/go-kratos/etcd/registry"
)

func main(){
	client, err := clientv3.New(clientv3.Config{
		Endpoints: []string{"127.0.0.1:2379"},
	})

	if err != nil {
		panic(err)
	}

	defer client.Close()

	ctx := context.Background()

	s := &registry.ServiceInstance{
		ID: "0",
		Name: "hellworld",
		Version: "v1.0",
		Metadata: map[string]string{
			"sss": "sss",
			"aaa": "bbbb",
			"ccc": "dddd",
		},
		Endpoints: []string{"127.0.0.1:8080", "127.0.0.1:8081"},
	}

	r := etcd_registry.New(client)

	w, err := r.Watch(ctx, s.Name)

	go func() {
		for {
			res, err := w.Next()

			if err != nil {
				return
			}

			for _, r := range res {
				fmt.Printf("next: %+v\n", r)
			}
		}
	}()

	if err := r.Register(ctx, s); err != nil {
		panic(err)
	}

	go func() {
		time.Sleep(1 * time.Second)
		a, e := r.GetService(ctx, s.Name)
		if e != nil {
			panic(e)
		}
		fmt.Printf("%v\n", a[0].Endpoints[0])
		//fmt.Println(e)
	}()

	select{}
}
