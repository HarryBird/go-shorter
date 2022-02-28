package data

import (
	"context"

	sv1 "github.com/HarryBird/url-shorten/api/shorten/v1"
	"github.com/HarryBird/url-shorten/app/gateway/internal/biz"
	"github.com/pkg/errors"

	mlog "github.com/HarryBird/mo-kit/kratos/log/app"
	"github.com/go-kratos/kratos/v2/log"
)

var _ biz.ShortenRepo = (*shortenRepo)(nil)

type shortenRepo struct {
	data *Data
	log  *log.Helper
}

func NewShortenRepo(data *Data, logger log.Logger) biz.ShortenRepo {
	return &shortenRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "mod", "repo.shorten")),
	}
}

func (r *shortenRepo) Shorten(ctx context.Context, in *biz.ShortenIn) (*biz.ShortenOut, error) {
	fname := "Shorten"
	reply, err := r.data.sc.CreateShortenURL(ctx, &sv1.CreateShortenURLRequest{Url: in.URL})
	if err != nil {
		mlog.LogErrorRPC(ctx, r.log, fname, err)
		return nil, errors.WithMessage(err, "repo: create shorten url fail by shorter service")
	}

	return &biz.ShortenOut{Code: reply.ShortenUrl.UrlCode}, nil
}

func (r *shortenRepo) Decode(ctx context.Context, in *biz.DecodeIn) (*biz.DecodeOut, error) {
	fname := "Decode"
	reply, err := r.data.sc.GetShortenURL(ctx, &sv1.GetShortenURLRequest{Query: &sv1.GetShortenURLRequest_Code{Code: in.Code}})
	if err != nil {
		mlog.LogErrorRPC(ctx, r.log, fname, err)
		return nil, errors.WithMessage(err, "repo: decode shorten url fail by shorter service")
	}

	return &biz.DecodeOut{URL: reply.ShortenUrl.UrlFull}, nil
}
