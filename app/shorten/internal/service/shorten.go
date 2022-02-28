package service

import (
	"context"
	"errors"

	"github.com/HarryBird/url-shorten/app/shorten/internal/biz"

	"github.com/go-kratos/kratos/v2/log"

	pb "github.com/HarryBird/url-shorten/api/shorten/v1"

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
		log: log.NewHelper(log.With(logger, "mod", "service.shorten")),
	}
}

// CreateShortenURL 创建短链
func (s *ShortenService) CreateShortenURL(ctx context.Context, req *pb.CreateShortenURLRequest) (*pb.CreateShortenURLReply, error) {
	var (
		fname = "CreateShortenURL"
		resp  = new(pb.CreateShortenURLReply)
	)

	mlog.LogRequest(ctx, s.log, fname, req)

	out, err := s.uc.Create(ctx, &biz.CreateIn{URL: req.Url})
	if err != nil {
		mlog.LogErrorStack(ctx, s.log, fname, err)
		return nil, pb.ErrorCreatrShortenUrlFail("%s", "create shorten url fail")
	}

	resp.ShortenUrl = &pb.ShortenURL{
		Id:      out.ID,
		UrlFull: out.URL,
		UrlCode: out.Code,
	}

	mlog.LogResponse(ctx, s.log, fname, resp)

	return resp, nil
}

// GetShortenURL 通过id或code，获取短链信息
func (s *ShortenService) GetShortenURL(ctx context.Context, req *pb.GetShortenURLRequest) (*pb.GetShortenURLReply, error) {
	var (
		fname = "GetShortenURL"
		resp  = new(pb.GetShortenURLReply)
	)

	mlog.LogRequest(ctx, s.log, fname, req)

	out, err := s.uc.Get(ctx, &biz.GetIn{
		ID:   req.GetId(),
		Code: req.GetCode(),
	})
	if err != nil {
		if errors.Is(err, biz.ErrNotFoundFromDB) {
			return nil, pb.ErrorShortenUrlNonexist("%s", "shorten url non exist")
		}

		mlog.LogErrorStack(ctx, s.log, fname, err)
		return nil, pb.ErrorGetShortenUrlFail("%s", "get shorten url fail")
	}

	resp.ShortenUrl = &pb.ShortenURL{
		Id:      out.ID,
		UrlFull: out.URL,
		UrlCode: out.Code,
	}

	mlog.LogResponse(ctx, s.log, fname, resp)

	return resp, nil
}

// DeleteShortenURL 删除短链
func (s *ShortenService) DeleteShortenURL(ctx context.Context, req *pb.DeleteShortenURLRequest) (*pb.DeleteShortenURLReply, error) {
	var (
		fname = "DeleteShortenURL"
		resp  = new(pb.DeleteShortenURLReply)
	)

	mlog.LogRequest(ctx, s.log, fname, req)

	err := s.uc.Delete(ctx, &biz.DeleteIn{
		ID:   req.GetId(),
		Code: req.GetCode(),
	})
	if err != nil {
		if errors.Is(err, biz.ErrNotFoundFromDB) {
			return nil, pb.ErrorShortenUrlNonexist("%s", "shorten url non exist")
		}

		mlog.LogErrorStack(ctx, s.log, fname, err)
		return nil, pb.ErrorDeleteShortenUrlFail("%s", "delete shorten url fail")
	}

	resp.Result = "ok"

	mlog.LogResponse(ctx, s.log, fname, resp)

	return resp, nil
}

func (s *ShortenService) ListShortenURL(ctx context.Context, req *pb.ListShortenURLRequest) (*pb.ListShortenURLReply, error) {
	return &pb.ListShortenURLReply{}, nil
}
