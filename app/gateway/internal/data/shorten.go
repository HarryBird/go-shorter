package data

import (
	"github.com/HarryBird/url-shorten/app/gateway/internal/biz"

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
