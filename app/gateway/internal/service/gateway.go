package service

import (
	"context"
	"errors"

	"github.com/HarryBird/url-shorten/app/gateway/internal/biz"
	"github.com/go-kratos/kratos/v2/log"

	mlog "github.com/HarryBird/mo-kit/kratos/log/app"
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
	var (
		fname = "ShortenURL"
		resp  = new(pb.ShortenURLReply)
	)

	mlog.LogRequest(ctx, s.log, fname, req)

	reply, err := s.sc.Shorten(ctx, &biz.ShortenIn{URL: req.Url})
	if err != nil {
		mlog.LogErrorStack(ctx, s.log, fname, err)
		return nil, pb.ErrorCreatrShortenUrlFail("%s", "create shorten url fail")
	}

	resp.Url = reply.URL
	mlog.LogResponse(ctx, s.log, fname, resp)

	return resp, nil
}

func (s *GatewayService) DecodeURL(ctx context.Context, req *pb.DecodeURLRequest) (*pb.DecodeURLReply, error) {
	var (
		fname = "DecodeURL"
		resp  = new(pb.DecodeURLReply)
	)

	mlog.LogRequest(ctx, s.log, fname, req)

	reply, err := s.sc.Decode(ctx, &biz.DecodeIn{Code: req.Code})
	if err != nil {
		if errors.Is(err, biz.ErrURLCodeNonexist) {
			return nil, pb.ErrorDecodeShortenUrlNonexist("invalid shorten code")
		}
		mlog.LogErrorStack(ctx, s.log, fname, err)
		return nil, pb.ErrorDecodeShortenUrlFail("%s", "decode shorten url fail")
	}

	resp.Url = reply.URL
	mlog.LogResponse(ctx, s.log, fname, resp)

	return resp, nil
}
