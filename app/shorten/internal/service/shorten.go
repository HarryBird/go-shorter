package service

import (
	"context"
	"errors"

	"github.com/go-kratos/kratos/v2/log"

	pb "url-shorten/api/shorten/v1"
	"url-shorten/app/shorten/internal/biz"

	mlog "github.com/HarryBird/mo-kit/kratos/log/app"
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

func (s *ShortenService) DecodeShortenURL(ctx context.Context,
	req *pb.DecodeShortenURLRequest) (*pb.DecodeShortenURLReply, error) {
	var (
		fname = "DecodeShortenURL"
	)

	mlog.LogRequest(ctx, s.log, fname, req)
	surl, err := s.uc.Decode(ctx, &biz.ShortenURL{URLCode: req.Code})

	if err != nil {
		if errors.Is(err, biz.ErrNotFoundFromDB) {
			return nil, pb.ErrorShortenCodeInvalid("%s", "invalid shorten code")
		}

		mlog.LogErrorStack(ctx, s.log, fname, err)
		return nil, pb.ErrorDecodeShortenUrlFail("%s", "decode shorten code fail")
	}

	resp := &pb.DecodeShortenURLReply{
		UrlFull: surl.URLFull,
	}
	mlog.LogResponse(ctx, s.log, fname, resp)

	return resp, nil
}

func (s *ShortenService) CreateShortenURL(ctx context.Context,
	req *pb.CreateShortenURLRequest) (*pb.CreateShortenURLReply, error) {
	var (
		fname = "CreateShortenURL"
		resp  = new(pb.CreateShortenURLReply)
	)

	mlog.LogRequest(ctx, s.log, fname, req)

	url, err := s.uc.Create(ctx, &biz.ShortenURL{URLFull: req.Url})
	if err != nil {
		mlog.LogErrorStack(ctx, s.log, fname, err)
		return nil, pb.ErrorCreatrShortenUrlFail("%s", "create shorten url fail")
	}

	resp.ShortenUrl = &pb.ShortenURL{
		Id:      url.ID,
		UrlFull: url.URLFull,
		UrlCode: url.URLCode,
	}

	mlog.LogResponse(ctx, s.log, fname, resp)

	return resp, nil
}

func (s *ShortenService) GetShortenURL(ctx context.Context, req *pb.GetShortenURLRequest) (*pb.GetShortenURLReply,
	error) {
	var (
		fname = "GetShortenURL"
		resp  = new(pb.GetShortenURLReply)
	)

	mlog.LogRequest(ctx, s.log, fname, req)

	url, err := s.uc.Get(ctx, &biz.ShortenURL{
		ID:      req.GetId(),
		URLCode: req.GetCode(),
	})

	if err != nil {
		if errors.Is(err, biz.ErrNotFoundFromDB) {
			if req.GetId() > 0 {
				return nil, pb.ErrorShortenIdInvalid("%s", "invalid shorten id")
			}

			if req.GetCode() != "" {
				return nil, pb.ErrorShortenCodeInvalid("%s", "invalid shorten code")
			}
		}

		mlog.LogErrorStack(ctx, s.log, fname, err)
		return nil, pb.ErrorGetShortenUrlFail("%s", "get shorten url fail")
	}

	resp.ShortenUrl = &pb.ShortenURL{
		Id:      url.ID,
		UrlFull: url.URLFull,
		UrlCode: url.URLCode,
	}

	mlog.LogResponse(ctx, s.log, fname, resp)

	return resp, nil
}

func (s *ShortenService) DeleteShortenURL(ctx context.Context,
	req *pb.DeleteShortenURLRequest) (*pb.DeleteShortenURLReply, error) {
	return &pb.DeleteShortenURLReply{}, nil
}
func (s *ShortenService) ListShortenURL(ctx context.Context,
	req *pb.ListShortenURLRequest) (*pb.ListShortenURLReply, error) {
	return &pb.ListShortenURLReply{}, nil
}
