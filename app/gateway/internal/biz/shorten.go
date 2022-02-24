package biz

import (
	"github.com/go-kratos/kratos/v2/log"
)

type ShortenRepo interface{}

type ShortenCase struct {
	repo ShortenRepo
	log  *log.Helper
}

func NewShortenCase(repo ShortenRepo, logger log.Logger) *ShortenCase {
	return &ShortenCase{repo: repo, log: log.NewHelper(log.With(logger, "mod", "biz.shorten"))}
}
