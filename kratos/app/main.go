package main

import (
	_ "github.com/dtm-labs/driver-kratos"
	"github.com/dtm-labs/dtmcli/logger"
	"github.com/dtm-labs/dtmdriver-clients/busi"
	v1 "github.com/dtm-labs/dtmdriver-clients/kratos/app/v1"
	"github.com/dtm-labs/dtmgrpc"
)

const (
	dtmServer  = "discovery://localhost:2379/dtmservice"
	busiServer = "discovery://127.0.0.1:2379/trans"
)

func main() {
	tcc(busiServer)
}

func tcc(busiServer string) {
	gid := dtmgrpc.MustGenGid(dtmServer)
	err := dtmgrpc.TccGlobalTransaction(dtmServer, gid, func(tcc *dtmgrpc.TccGrpc) error {
		rep := v1.Response{}
		err1 := tcc.CallBranch(&busi.BusiReq{Amount: 30, UserId: 1},
			busiServer+"/api.trans.v1.Trans/TransIn",
			busiServer+"/api.trans.v1.Trans/TransOut", // 应当为TransOutConfirm，后续会完善例子
			busiServer+"/api.trans.v1.Trans/TransOut", // 应当为TransOutCancel，后续会完善例子
			&rep)
		return err1
	})
	logger.FatalIfError(err)
}
