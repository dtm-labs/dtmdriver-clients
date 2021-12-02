package main

import (
	"flag"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/zrpc"
	"github.com/yedf/dtmcli/dtmimp"
	"github.com/yedf/dtmdriver"
	"github.com/yedf/dtmdriver-clients/busi"
	trans "github.com/yedf/dtmdriver-clients/gozero/trans/pb"
	_ "github.com/yedf/dtmdriver-gozero"
	"github.com/yedf/dtmgrpc"
)

var (
	configFile = flag.String("f", "etc/config.yaml", "the config file")
	outReq     = &trans.AdjustInfo{Amount: 30, UserID: 1}
	inReq      = &trans.AdjustInfo{Amount: 30, UserID: 2}
)

func main() {
	flag.Parse()

	var c zrpc.RpcClientConf
	conf.MustLoad(*configFile, &c)

	err := dtmdriver.Use("dtm-driver-gozero")
	dtmimp.FatalIfError(err)

	dtmServer := "etcd://localhost:2379/dtmservice"
	busiServer, err := c.BuildTarget()
	dtmimp.FatalIfError(err)

	gid := dtmgrpc.MustGenGid(dtmServer)
	msg := dtmgrpc.NewMsgGrpc(dtmServer, gid).
		Add(busiServer+"/trans.TransSvc/TransOut", &busi.BusiReq{Amount: 30, UserId: 1}).
		Add(busiServer+"/trans.TransSvc/TransIn", &busi.BusiReq{Amount: 30, UserId: 2})

	err = msg.Submit()
	dtmimp.FatalIfError(err)

}
