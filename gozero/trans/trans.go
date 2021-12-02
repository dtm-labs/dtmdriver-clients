package main

import (
	"flag"
	"fmt"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/core/service"
	"github.com/tal-tech/go-zero/zrpc"
	"github.com/yedf/dtmdriver-clients/gozero/trans/internal/config"
	"github.com/yedf/dtmdriver-clients/gozero/trans/internal/server"
	"github.com/yedf/dtmdriver-clients/gozero/trans/internal/svc"
	"github.com/yedf/dtmdriver-clients/gozero/trans/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/trans.yaml", "the config file")

func main() {
	flag.Parse()
	logx.Disable()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	srv := server.NewTransSvcServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		reflection.Register(grpcServer)
		pb.RegisterTransSvcServer(grpcServer, srv)

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
