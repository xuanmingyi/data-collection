package main

import (
	"context"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/xuanmingyi/data-collection/client/ent"
	"google.golang.org/grpc"

	pb "github.com/xuanmingyi/data-collection/api/node/service/v1"
)

var (
	database_driver = "sqlite3"
	database_dsn    = "file:a.db?_fk=1"

	address      = "localhost:50051"
	default_name = "world"
	client       *ent.Client
	conn         *grpc.ClientConn
	node_client  pb.NodeClient
)

func init() {
	var err error

	client, err = ent.Open(database_driver, database_dsn)

	if err != nil {
		panic(err)
	}

	if err = client.Schema.Create(context.Background()); err != nil {
		panic(err)
	}

	conn, err = grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		panic(err)
	}

	node_client = pb.NewNodeClient(conn)

}

func main() {
	defer client.Close()
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := node_client.AddNode(ctx, &pb.AddNodeReq{Uuid: "sssssssssssssssss"})
	if err != nil {
		panic(err)
	}
	fmt.Println(r.GetToken())
}
