package service

import (
	"context"
	"log"

	pb "trans/api/trans/v1"
)

type TransService struct {
	pb.UnimplementedTransServer
}

func NewTransService() *TransService {
	return &TransService{}
}

func (s *TransService) TransIn(ctx context.Context, req *pb.AdjustInfo) (*pb.Response, error) {
	log.Printf("transfer in %d cents to %d", req.Amount, req.UserID)
	return &pb.Response{}, nil
}
func (s *TransService) TransOut(ctx context.Context, req *pb.AdjustInfo) (*pb.Response, error) {
	log.Printf("transfer out %d cents from %d", req.Amount, req.UserID)
	return &pb.Response{}, nil
}
