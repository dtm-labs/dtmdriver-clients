package main

import (
	"context"
	"fmt"

	"github.com/dtm-labs/dtmdriver-clients/busi"
	"github.com/gotomicro/ego"
	"github.com/gotomicro/ego-component/eetcd"
	"github.com/gotomicro/ego-component/eetcd/registry"
	"github.com/gotomicro/ego/core/elog"
	"github.com/gotomicro/ego/server"
	"github.com/gotomicro/ego/server/egrpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

// export EGO_DEBUG=true EGO_NAME=hello && go run main.go --config=config.toml
func main() {
	if err := ego.New().
		Registry(registry.Load("registry").Build(registry.WithClientEtcd(eetcd.Load("etcd").Build()))).
		Serve(func() server.Server {
			server := egrpc.Load("server.grpc").Build()
			busi.RegisterBusiServer(server.Server, &Trans{server: server})
			return server
		}()).Run(); err != nil {
		elog.Panic("startup", elog.Any("err", err))
	}
}

type Trans struct {
	server *egrpc.Component
	busi.UnsafeBusiServer
}

func (s *Trans) TransIn(ctx context.Context, in *busi.BusiReq) (*emptypb.Empty, error) {
	elog.Info("[info]", elog.FieldDescription(fmt.Sprintf("transfer in %d cents to %d", in.GetAmount(), in.GetUserId())))
	return &emptypb.Empty{}, nil
}

func (s *Trans) TransOut(ctx context.Context, in *busi.BusiReq) (*emptypb.Empty, error) {
	elog.Info("[info]", elog.FieldDescription(fmt.Sprintf("transfer out %d cents from %d", in.GetAmount(), in.GetUserId())))
	return &emptypb.Empty{}, nil
}
