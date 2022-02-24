package service

import (
	"context"

	"github.com/HarryBird/url-shorten/app/gateway/internal/biz"
	"github.com/go-kratos/kratos/v2/log"

	pb "github.com/HarryBird/url-shorten/api/gateway/v1"
)

type GatewayService struct {
	pb.UnimplementedGatewayServer

	sc  *biz.ShortenCase
	log *log.Helper
}

func NewGatewayService(sc *biz.ShortenCase, logger log.Logger) *GatewayService {
	return &GatewayService{
		sc:  sc,
		log: log.NewHelper(log.With(logger, "mod", "service.gateway")),
	}
}

func (s *GatewayService) ShortenURL(ctx context.Context, req *pb.ShortenURLRequest) (*pb.ShortenURLReply, error) {
	return &pb.ShortenURLReply{}, nil
}

func (s *GatewayService) DecodeURL(ctx context.Context, req *pb.DecodeURLRequest) (*pb.DecodeURLReply, error) {
	return &pb.DecodeURLReply{}, nil
}
