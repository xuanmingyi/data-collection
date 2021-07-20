package conf

import (
	"context"
	"time"

	"google.golang.org/grpc"

	pb "lithum/api/conf"
)

func GetConf(url string, uuid string) (content string) {
	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	confClient := pb.NewConfClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	reply, err := confClient.GetConf(ctx, &pb.GetConfReq{
		Tid:  33,
		Uuid: uuid,
	})

	if err != nil {
		panic(err)
	}

	return reply.Content
}
