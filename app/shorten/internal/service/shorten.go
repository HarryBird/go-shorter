package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	pb "url-shorten/api/shorten/v1"
	"url-shorten/app/shorten/internal/biz"

	"github.com/HarryBird/mo-kit/msgr"
)

type ShortenService struct {
	pb.UnimplementedShortenServer

	uc  *biz.ShortenCase
	log *log.Helper
}

func NewShortenService(uc *biz.ShortenCase, logger log.Logger) *ShortenService {
	return &ShortenService{
		uc:  uc,
		log: log.NewHelper(log.With(logger, "mod", "service.shorten"))}
}

func (s *ShortenService) logRequest(ctx context.Context, fname string, req interface{}) {
	s.log.WithContext(ctx).Infof("%s Request Begin...", msgr.W(fname))
	s.log.WithContext(ctx).Debugf("%s Request Param: %+v", msgr.W(fname), req)
}

func (s *ShortenService) logResponse(ctx context.Context, fname string, resp interface{}) {
	s.log.WithContext(ctx).Info("%s Request End...", msgr.W(fname))
	s.log.WithContext(ctx).Debugf("%s Response: %+v", msgr.W(fname), resp)
}

func (s *ShortenService) CreateShortenURL(ctx context.Context, req *pb.CreateShortenURLRequest) (*pb.CreateShortenURLReply, error) {
	var (
		fname = "CreateShortenURL"
		resp  = &pb.CreateShortenURLReply{}
	)

	s.logRequest(ctx, fname, req)

	surl, err := s.uc.Create(ctx, biz.OriginURL{Url: req.Url})
	if err != nil {
		s.log.WithContext(ctx).Errorf("%s create shorten url fail: %v", msgr.W(fname), err)
		s.log.WithContext(ctx).Errorf("%s error stack=%+v", msgr.W(fname), err)
		return nil, err
	}

	resp = &pb.CreateShortenURLReply{
		ShortUrl: "https://localhost/" + surl.URLCode,
	}

	s.logResponse(ctx, fname, resp)

	return resp, nil
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
