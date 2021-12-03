package main

import (
	"flag"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/zrpc"
	"github.com/yedf/dtmcli/dtmimp"
	"github.com/yedf/dtmdriver"
	"github.com/yedf/dtmdriver-clients/busi"
	"github.com/yedf/dtmdriver-clients/gozero/trans/pb"
	trans "github.com/yedf/dtmdriver-clients/gozero/trans/pb"
	_ "github.com/yedf/dtmdriver-gozero"
	"github.com/yedf/dtmgrpc"
)

var (
	configFile = flag.String("f", "etc/config.yaml", "the config file")
	outReq     = &trans.AdjustInfo{Amount: 30, UserID: 1}
	inReq      = &trans.AdjustInfo{Amount: 30, UserID: 2}
)

var dtmServer = "etcd://localhost:2379/dtmservice"

func main() {
	flag.Parse()

	var c zrpc.RpcClientConf
	conf.MustLoad(*configFile, &c)

	err := dtmdriver.Use("dtm-driver-gozero")
	dtmimp.FatalIfError(err)

	busiServer, err := c.BuildTarget()
	dtmimp.FatalIfError(err)

	msg(busiServer)
	tcc(busiServer)
}

func msg(busiServer string) {
	gid := dtmgrpc.MustGenGid(dtmServer)
	msg := dtmgrpc.NewMsgGrpc(dtmServer, gid).
		Add(busiServer+"/trans.TransSvc/TransOut", &busi.BusiReq{Amount: 30, UserId: 1}).
		Add(busiServer+"/trans.TransSvc/TransIn", &busi.BusiReq{Amount: 30, UserId: 2})

	err := msg.Submit()
	dtmimp.FatalIfError(err)
}

func tcc(busiServer string) {
	gid := dtmgrpc.MustGenGid(dtmServer)
	err := dtmgrpc.TccGlobalTransaction(dtmServer, gid, func(tcc *dtmgrpc.TccGrpc) error {
		rep := pb.Response{}
		err1 := tcc.CallBranch(&busi.BusiReq{Amount: 30, UserId: 1},
			busiServer+"/trans.TransSvc/TransOut",
			busiServer+"/trans.TransSvc/TransOut", // 应当为TransOutConfirm，后续会完善例子
			busiServer+"/trans.TransSvc/TransOut", // 应当为TransOutCancel，后续会完善例子
			&rep)
		return err1
	})
	dtmimp.FatalIfError(err)
}
