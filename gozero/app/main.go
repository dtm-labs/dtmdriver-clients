package main

import (
	"flag"

	"github.com/dtm-labs/client/dtmcli/dtmimp"
	"github.com/dtm-labs/dtmdriver-clients/busi"
	"github.com/dtm-labs/dtmdriver-clients/gozero/trans/pb"
	trans "github.com/dtm-labs/dtmdriver-clients/gozero/trans/pb"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/zrpc"

	// 下面这行导入gozero的dtm驱动
	"github.com/dtm-labs/client/dtmgrpc"
	_ "github.com/dtm-labs/driver-gozero"
)

var (
	configFile = flag.String("f", "etc/config.yaml", "the config file")
	outReq     = &trans.AdjustInfo{Amount: 30, UserID: 1}
	inReq      = &trans.AdjustInfo{Amount: 30, UserID: 2}
)

// dtm已经通过配置，注册到下面这个地址，因此在dtmgrpc中使用该地址
var dtmServer = "etcd://localhost:2379/dtmservice"

func main() {
	flag.Parse()

	// 下面从配置文件中Load配置，然后通过BuildTarget获得业务服务的地址
	var c zrpc.RpcClientConf
	conf.MustLoad(*configFile, &c)
	busiServer, err := c.BuildTarget()
	dtmimp.FatalIfError(err)

	msg(busiServer)
	tcc(busiServer)
}

func msg(busiServer string) {
	gid := dtmgrpc.MustGenGid(dtmServer)
	msg := dtmgrpc.NewMsgGrpc(dtmServer, gid).
		// 事务的第一步为调用trans.TransSvcClient.TransOut
		// 可以从trans.pb.go中找到上述方法对应的Method名称为"/trans.TransSvc/TransOut"
		// dtm需要从dtm服务器调用该方法，所以不走强类型
		// 而是走动态的url: busiServer+"/trans.TransSvc/TransOut"
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
