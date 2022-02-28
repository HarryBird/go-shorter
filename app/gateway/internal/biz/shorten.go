package biz

import (
	"context"
	"fmt"

	"github.com/HarryBird/url-shorten/app/gateway/internal/conf"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/pkg/errors"
)

type ShortenRepo interface {
	Shorten(ctx context.Context, in *ShortenIn) (*ShortenOut, error)
	Decode(ctx context.Context, in *DecodeIn) (*DecodeOut, error)
}

type ShortenCase struct {
	conf *conf.App
	repo ShortenRepo
	log  *log.Helper
}

func NewShortenCase(conf *conf.App, repo ShortenRepo, logger log.Logger) *ShortenCase {
	return &ShortenCase{conf: conf, repo: repo, log: log.NewHelper(log.With(logger, "mod", "biz.shorten"))}
}

func (uc *ShortenCase) Shorten(ctx context.Context, in *ShortenIn) (*ShortenOut, error) {
	out, err := uc.repo.Shorten(ctx, in)
	if err != nil {
		return nil, errors.WithMessage(err, "biz: create shorten url fail")
	}

	out.URL = fmt.Sprintf("%s/%s", uc.conf.Shorten.Host, out.Code)
	return out, nil
}

func (uc *ShortenCase) Decode(ctx context.Context, in *DecodeIn) (*DecodeOut, error) {
	out, err := uc.repo.Decode(ctx, in)
	if err != nil {
		return nil, errors.WithMessage(err, "biz: decode shorten url fail")
	}

	return out, nil
}
