package data

import (
	"context"
	"time"
	"url-shorten/app/shorten/internal/biz"
	"url-shorten/app/shorten/internal/data/dao/model"
	"url-shorten/app/shorten/internal/data/dao/query"

	"github.com/HarryBird/mo-kit/msgr"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/pkg/errors"
)

var _ biz.ShortenRepo = (*shortenRepo)(nil)

type shortenRepo struct {
	data *Data
	log  *log.Helper
}

func NewShortenRepo(data *Data, logger log.Logger) biz.ShortenRepo {
	return &shortenRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "repo.shorten")),
	}
}

func (r *shortenRepo) Create(ctx context.Context, url *biz.ShortenURL) (*biz.ShortenURL, error) {
	fname := "Create"
	dao := query.Use(r.data.db).URLShortened
	u := model.URLShortened{
		URLFull:   url.URLFull,
		URLHost:   url.URLHost,
		URLURI:    url.URLUri,
		URLQuery:  url.URLQuery,
		URLCode:   url.URLCode,
		CreatedAt: time.Now().Unix(),
	}

	r.log.WithContext(ctx).Infof("%s prepare data: %+v", msgr.W(fname), u)
	if err := dao.WithContext(ctx).Create(&u); err != nil {
		return nil, errors.Wrap(err, "repo: create shorten url fail")
	}

	r.log.WithContext(ctx).Infof("%s after insert db: %+v", msgr.W(fname), u)
	return url, nil
}
