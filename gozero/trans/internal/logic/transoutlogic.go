package logic

import (
	"context"
	"log"

	"github.com/tal-tech/go-zero/core/logx"
	"github.com/yedf/dtmdriver-clients/gozero/trans/internal/svc"
	"github.com/yedf/dtmdriver-clients/gozero/trans/pb"
)

type TransOutLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTransOutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TransOutLogic {
	return &TransOutLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TransOutLogic) TransOut(in *pb.AdjustInfo) (*pb.Response, error) {
	log.Printf("transfer out %d cents from %d", in.Amount, in.UserID)
	return &pb.Response{}, nil
}
