package main

import (
	"github.com/dtm-labs/dtmdriver-clients/busi"
	"github.com/gotomicro/ego"
	"github.com/gotomicro/ego-component/eetcd"
	"github.com/gotomicro/ego-component/eetcd/registry"
	"github.com/gotomicro/ego/client/egrpc"
	"github.com/gotomicro/ego/client/egrpc/resolver"
	"github.com/gotomicro/ego/core/elog"

	// 下面这行导入ego的dtm驱动
	"github.com/dtm-labs/client/dtmgrpc"
	_ "github.com/dtm-labs/driver-ego"
)

var dtmServer = "etcd:///dtmservice"
var busiServer string

// export EGO_DEBUG=true EGO_NAME=hello-client && go run main.go --config=config.toml
func main() {
	if err := ego.New().Invoker(
		invokerGrpc,
		callGrpc,
	).Run(); err != nil {
		elog.Error("startup", elog.FieldErr(err))
	}
}

// var grpcComp busi.BusiClient

func invokerGrpc() error {
	// 必须注册在grpc前面
	resolver.Register("etcd", registry.Load("registry").Build(registry.WithClientEtcd(eetcd.Load("etcd").Build())))

	grpcConn := egrpc.Load("grpc.hello").Build()
	busiServer = grpcConn.Target()

	// grpcComp = busi.NewBusiClient(grpcConn.ClientConn)
	// if _, err := grpcComp.TransIn(context.Background(), &busi.BusiReq{Amount: 30, UserId: 1}); err != nil {
	// 	fmt.Println(err)
	// }
	return nil
}

func callGrpc() error {
	gid := dtmgrpc.MustGenGid(dtmServer)
	msg := dtmgrpc.NewMsgGrpc(dtmServer, gid).
		Add(busiServer+"/busi.Busi/TransOut", &busi.BusiReq{Amount: 30, UserId: 1}).
		Add(busiServer+"/busi.Busi/TransIn", &busi.BusiReq{Amount: 30, UserId: 2})
	return msg.Submit()
}
