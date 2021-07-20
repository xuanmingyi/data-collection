package conf

import (
	"context"

	"google.golang.org/grpc"

	pb "lithum/api/conf"
)

type ConfServer struct {
	pb.ConfServer
}

func (s *ConfServer) GetConf(ctx context.Context, req *pb.GetConfReq) (reply *pb.GetConfReply, err error) {
	//	var content []byte

	//var flag bool = false
	/*
		for _, client := range Config.Clients {
			if client.Uuid == req.GetUuid() {
				fmt.Println(2222222222)
				content, err = json.Marshal(client)
				fmt.Println(11111111)

				//		flag = true
				if err != nil {
					return nil, err
				}
			}
		}
	*/

	s1 := "helloworld!"
	return &pb.GetConfReply{
		Tid:     req.GetTid(),
		Uuid:    req.GetUuid(),
		Content: s1,
	}, nil

}

func RegisterConfServer(server *grpc.Server) {
	pb.RegisterConfServer(server, &ConfServer{})
}
