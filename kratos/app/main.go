package main

import (
	"github.com/dtm-labs/client/dtmcli/logger"
	"github.com/dtm-labs/client/dtmgrpc"
	_ "github.com/dtm-labs/driver-kratos"
	"github.com/dtm-labs/dtmdriver-clients/busi"
	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/transport/grpc/resolver/discovery"
	etcdAPI "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc/resolver"
	"strings"
)

const (
	etcdServer = "localhost:2379"
	dtmServer  = "discovery:///dtmservice"
	busiServer = "discovery:///trans"
)

func main() {
	//create registry
	client, err := etcdAPI.New(etcdAPI.Config{
		Endpoints: strings.Split(etcdServer, ","),
	})
	if err != nil {
		panic(err)
	}
	registry := etcd.New(client)

	//register global resolver so that dtm client can resolve dtm server itself by registry
	resolver.Register(discovery.NewBuilder(registry, discovery.WithInsecure(true)))

	msg(busiServer)
}

func msg(busiServer string) {
	gid := dtmgrpc.MustGenGid(dtmServer)
	m := dtmgrpc.NewMsgGrpc(dtmServer, gid).
		Add(busiServer+"/api.trans.v1.Trans/TransOut", &busi.BusiReq{Amount: 30, UserId: 1}).
		Add(busiServer+"/api.trans.v1.Trans/TransIn", &busi.BusiReq{Amount: 30, UserId: 2})
	m.WaitResult = true
	err := m.Submit()
	logger.FatalIfError(err)
}
