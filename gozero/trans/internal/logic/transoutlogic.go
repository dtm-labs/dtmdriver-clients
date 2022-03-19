package logic

import (
	"context"
	"log"

	"github.com/dtm-labs/dtmdriver-clients/gozero/trans/internal/svc"
	"github.com/dtm-labs/dtmdriver-clients/gozero/trans/pb"
	"github.com/zeromicro/go-zero/core/logx"
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
