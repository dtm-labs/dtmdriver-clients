// Code generated by goctl. DO NOT EDIT!
// Source: trans.proto

package server

import (
	"context"

	"github.com/yedf/dtmdriver-clients/gozero/trans/internal/logic"
	"github.com/yedf/dtmdriver-clients/gozero/trans/internal/svc"
	"github.com/yedf/dtmdriver-clients/gozero/trans/pb"
)

type TransSvcServer struct {
	svcCtx *svc.ServiceContext
}

func NewTransSvcServer(svcCtx *svc.ServiceContext) *TransSvcServer {
	return &TransSvcServer{
		svcCtx: svcCtx,
	}
}

func (s *TransSvcServer) TransIn(ctx context.Context, in *pb.AdjustInfo) (*pb.Response, error) {
	l := logic.NewTransInLogic(ctx, s.svcCtx)
	return l.TransIn(in)
}

func (s *TransSvcServer) TransOut(ctx context.Context, in *pb.AdjustInfo) (*pb.Response, error) {
	l := logic.NewTransOutLogic(ctx, s.svcCtx)
	return l.TransOut(in)
}
