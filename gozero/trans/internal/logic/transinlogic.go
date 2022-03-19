package logic

import (
	"context"
	"log"

	"github.com/dtm-labs/dtmdriver-clients/gozero/trans/internal/svc"
	"github.com/dtm-labs/dtmdriver-clients/gozero/trans/pb"
	"github.com/zeromicro/go-zero/core/logx"
)

type TransInLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTransInLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TransInLogic {
	return &TransInLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TransInLogic) TransIn(in *pb.AdjustInfo) (*pb.Response, error) {
	log.Printf("transfer in %d cents to %d", in.Amount, in.UserID)
	return &pb.Response{}, nil
}
