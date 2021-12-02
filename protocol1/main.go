package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/yedf/dtmcli/dtmimp"
	"github.com/yedf/dtmdriver"
	"github.com/yedf/dtmdriver-clients/busi"
	_ "github.com/yedf/dtmdriver-protocol1"
	"github.com/yedf/dtmgrpc"
	"google.golang.org/grpc"
)

var busiPort = 57070

func main() {
	s := grpc.NewServer()
	busi.RegisterBusiServer(s, &busi.BusiServerImp{})

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", busiPort))
	dtmimp.FatalIfError(err)
	go func() {
		log.Printf("busi grpc listening at %v", lis.Addr())
		err := s.Serve(lis)
		dtmimp.FatalIfError(err)
	}()
	time.Sleep(100 * time.Millisecond)

	err = dtmdriver.Use("dtm-driver-protocol1")
	dtmimp.FatalIfError(err)
	SampleCall("")
	SampleCall("protocol1://")
	select {}
}

func SampleCall(prefix string) {
	busiServer := fmt.Sprintf("%slocalhost:%d", prefix, busiPort)
	dtmServer := fmt.Sprintf("%slocalhost:36790", prefix)
	gid := dtmgrpc.MustGenGid(dtmServer)
	msg := dtmgrpc.NewMsgGrpc(dtmServer, gid).
		Add(busiServer+"/busi.Busi/TransOut", &busi.BusiReq{Amount: 30, UserId: 1}).
		Add(busiServer+"/busi.Busi/TransIn", &busi.BusiReq{Amount: 30, UserId: 2})

	err := msg.Submit()
	dtmimp.FatalIfError(err)
}
