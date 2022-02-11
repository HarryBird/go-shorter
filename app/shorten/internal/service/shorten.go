package service

import (
	"context"

	pb "url-shorten/api/shorten/v1"
)

type ShortenService struct {
	pb.UnimplementedShortenServer
}

func NewShortenService() *ShortenService {
	return &ShortenService{}
}

func (s *ShortenService) CreateShortenURL(ctx context.Context, req *pb.CreateShortenURLRequest) (*pb.CreateShortenURLReply, error) {
	return &pb.CreateShortenURLReply{}, nil
}
func (s *ShortenService) DeleteShortenURL(ctx context.Context, req *pb.DeleteShortenURLRequest) (*pb.DeleteShortenURLReply, error) {
	return &pb.DeleteShortenURLReply{}, nil
}
func (s *ShortenService) GetShortenURL(ctx context.Context, req *pb.GetShortenURLRequest) (*pb.GetShortenURLReply, error) {
	return &pb.GetShortenURLReply{}, nil
}
func (s *ShortenService) ListShortenURL(ctx context.Context, req *pb.ListShortenURLRequest) (*pb.ListShortenURLReply, error) {
	return &pb.ListShortenURLReply{}, nil
}
