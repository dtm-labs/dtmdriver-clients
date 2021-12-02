package busi

import (
	context "context"
	"log"

	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type BusiServerImp struct {
	UnimplementedBusiServer
}

func (s *BusiServerImp) TransIn(ctx context.Context, in *BusiReq) (*emptypb.Empty, error) {
	log.Printf("TransIn %d to user %d", in.Amount, in.UserId)
	return &emptypb.Empty{}, nil
}

func (s *BusiServerImp) TransOut(ctx context.Context, in *BusiReq) (*emptypb.Empty, error) {
	log.Printf("TransOut %d from user %d", in.Amount, in.UserId)
	return &emptypb.Empty{}, nil
}
