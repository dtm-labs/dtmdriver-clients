package main

import (
	"github.com/dtm-labs/client/dtmcli/logger"
	"github.com/dtm-labs/client/dtmgrpc"
	_ "github.com/dtm-labs/driver-kratos"
	"github.com/dtm-labs/dtmdriver-clients/busi"
)

const (
	dtmServer  = "etcd://localhost:2379/dtmservice"
	busiServer = "discovery://127.0.0.1:2379/trans"
)

func main() {
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
