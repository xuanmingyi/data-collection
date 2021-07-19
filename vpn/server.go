package vpn

import (
	"context"
	"fmt"
	"sync"
)

type VPNServer struct {
	Ctx   context.Context
	Group *sync.WaitGroup
}

func (s *VPNServer) Run() {
	for {
		select {
		case <-s.Ctx.Done():
			fmt.Println("vpn exit()")
			s.Group.Done()
			return
		}
	}
}

func NewVPNServer(ctx context.Context, wg *sync.WaitGroup) *VPNServer {
	return &VPNServer{
		Ctx:   ctx,
		Group: wg,
	}
}
