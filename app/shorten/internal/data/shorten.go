package data

import (
	"context"
	"time"
	"url-shorten/app/shorten/internal/biz"
	"url-shorten/app/shorten/internal/data/dao/model"
	"url-shorten/app/shorten/internal/data/dao/query"

	"github.com/HarryBird/mo-kit/msgr"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"gorm.io/gorm"
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

func (r *shortenRepo) Delete(ctx context.Context, opts ...query.Option) error {
	err := r.delete(ctx, opts...)
	if err != nil {
		if err == biz.ErrNotFoundFromDB {
			return err
		}

		return errors.WithMessage(err, "repo: delete url shorten fail")
	}

	return nil
}

func (r *shortenRepo) Get(ctx context.Context, opts ...query.Option) (*biz.ShortenURL, error) {
	// fname := "Get"

	url, err := r.queryOne(ctx, opts...)
	if err != nil {
		if err == biz.ErrNotFoundFromDB {
			return nil, err
		}

		return nil, errors.WithMessage(err, "repo: get url shorten fail")
	}

	return r.modelToBiz(ctx, url), nil
}

func (r *shortenRepo) delete(ctx context.Context, opts ...query.Option) error {
	fname := "delete"

	cond, err := r.applyQueryOption(ctx, opts...)
	if err != nil {
		return errors.WithMessage(err, "repo: delete fail by applyQueryOption")
	}

	fields := query.Use(r.data.db).URLShortened
	result, err := query.UseCondition(ctx, r.data.db, cond).Update(fields.DeletedAt, time.Now().Unix())

	if result.RowsAffected == 0 {
		r.log.WithContext(ctx).Warnf("%s delete shorten url fail: not found", msgr.W(fname))
		return biz.ErrNotFoundFromDB
	}

	if err != nil {
		return errors.WithMessage(err, "repo: delete fail by db")
	}

	return nil
}

func (r *shortenRepo) queryOne(ctx context.Context, opts ...query.Option) (*model.URLShortened, error) {
	fname := "queryOne"
	cond, err := r.applyQueryOption(ctx, opts...)
	if err != nil {
		return nil, errors.WithMessage(err, "repo: query row fail by applyQueryOption")
	}

	url, err := query.UseCondition(ctx, r.data.db, cond).First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			r.log.WithContext(ctx).Warnf("%s query row from db: not found", msgr.W(fname))
			return nil, biz.ErrNotFoundFromDB
		}

		r.log.WithContext(ctx).Errorf("%s query row from db fail: err=%v", msgr.W(fname), err)
		return nil, errors.WithMessage(err, "repo: query row fail by db")
	}

	return url, nil
}

func (r *shortenRepo) applyQueryOption(ctx context.Context, opts ...query.Option) (*query.Condition, error) {
	fname := "applyQueryOption"

	if len(opts) == 0 {
		return nil, errors.WithMessage(biz.ErrParamInvalid, "repo: apply query option fail")
	}

	qo := query.NewCondition()

	for _, opt := range opts {
		opt(qo)
	}

	r.log.WithContext(ctx).Infof("%s query option detail: %+v", msgr.W(fname), qo)
	return qo, nil
}

func (r *shortenRepo) Decode(ctx context.Context, url *biz.ShortenURL) (*biz.ShortenURL, error) {
	fname := "Decode"

	key, ttl := RedisKeyShortenCodeToURL.extract(url.URLCode)
	r.log.WithContext(ctx).Infof("%s try query from redis: key=%s", msgr.W(fname), key)
	v, err := r.data.rdb.Get(ctx, key).Result()

	if err == nil {
		r.log.WithContext(ctx).Infof("%s query from redis: result=%s", msgr.W(fname), v)
		return &biz.ShortenURL{URLFull: v}, nil
	}

	if err != redis.Nil {
		r.log.WithContext(ctx).Errorf("%s query from redis fail: err=%v", msgr.W(fname), err)
		return nil, errors.WithMessage(err, "repo: query shorten url from redis fail")
	}

	r.log.WithContext(ctx).Infof("%s query from redis: not found", msgr.W(fname))

	dao := query.Use(r.data.db).URLShortened

	r.log.WithContext(ctx).Infof("%s try query from db: code=%s", msgr.W(fname), url.URLCode)
	shortenURL, err := dao.WithContext(ctx).Where(dao.URLCode.Eq(url.URLCode)).First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			r.log.WithContext(ctx).Warnf("%s query from db: not found", msgr.W(fname))
			return nil, biz.ErrNotFoundFromDB
		}

		r.log.WithContext(ctx).Errorf("%s query from db fail: err=%v", msgr.W(fname), err)
		return nil, errors.WithMessage(err, "repo: query shorten url by code fail")
	}

	r.log.WithContext(ctx).Infof("%s try write full url to redis : key=%s, val=%s, ttl=%v",
		msgr.W(fname), key, shortenURL.URLFull, ttl)
	if err := r.data.rdb.SetEX(ctx, key, shortenURL.URLFull, ttl).Err(); err != nil {
		r.log.WithContext(ctx).Errorf("%s write full url to redis fail: err=%v", msgr.W(fname), err)
		return nil, errors.WithMessage(err, "repo: write shorten url to redis fail")
	}

	return r.modelToBiz(ctx, shortenURL), nil
}

func (r *shortenRepo) Create(ctx context.Context, url *biz.ShortenURL) (*biz.ShortenURL, error) {
	fname := "Create"
	dao := query.Use(r.data.db).URLShortened

	r.log.WithContext(ctx).Infof("%s check shorten url existed: urlcode=%v", msgr.W(fname), url.URLCode)
	existedShortenURL, err := dao.WithContext(ctx).Where(dao.URLCode.Eq(url.URLCode)).First()

	// 短链已存在
	if err == nil {
		r.log.WithContext(ctx).Infof("%s shorten url had existed: %+v", msgr.W(fname), existedShortenURL)
		if existedShortenURL.URLFull == url.URLFull {
			return r.modelToBiz(ctx, existedShortenURL), nil
		}

		// 短链存在，但原始链接不同，hash 冲突
		return nil, errors.New("repo: shorten url hash conflict")
	}

	// 短链不存在，创建
	if errors.Is(err, gorm.ErrRecordNotFound) {
		newShortenURL, err := r.create(ctx, url)
		if err != nil {
			return nil, err
		}

		return r.modelToBiz(ctx, newShortenURL), nil
	}

	return nil, errors.WithMessage(err, "repo: query shorten url existed fail")
}

func (r *shortenRepo) WithID(id int64) query.Option {
	return func(o *query.Condition) { o.Where.Id = id }
}

func (r *shortenRepo) WithCode(code string) query.Option {
	return func(o *query.Condition) { o.Where.Code = code }
}

func (r *shortenRepo) WithOffset(offset int64) query.Option {
	return func(o *query.Condition) { o.Paging.Offset = cast.ToInt(offset) }
}

func (r *shortenRepo) WithLimit(limit int64) query.Option {
	return func(o *query.Condition) { o.Paging.Limit = cast.ToInt(limit) }
}

func (r *shortenRepo) WithDeleted(isDel bool) query.Option {
	return func(o *query.Condition) { o.Where.IsDeleted = isDel }
}

func (r *shortenRepo) create(ctx context.Context, url *biz.ShortenURL) (*model.URLShortened, error) {
	fname := "create"
	dao := query.Use(r.data.db).URLShortened

	// biz to model
	u := r.bizToModel(ctx, url)

	// store
	r.log.WithContext(ctx).Infof("%s prepare data: %+v", msgr.W(fname), u)
	if err := dao.WithContext(ctx).Create(u); err != nil {
		return nil, errors.Wrap(err, "repo: create shorten url fail")
	}
	r.log.WithContext(ctx).Infof("%s after insert db: %+v", msgr.W(fname), u)

	return u, nil
}

func (r *shortenRepo) bizToModel(ctx context.Context, url *biz.ShortenURL) *model.URLShortened {
	return &model.URLShortened{
		URLFull:  url.URLFull,
		URLHost:  url.URLHost,
		URLURI:   url.URLUri,
		URLQuery: url.URLQuery,
		URLCode:  url.URLCode,
	}
}

func (r *shortenRepo) modelToBiz(ctx context.Context, url *model.URLShortened) *biz.ShortenURL {
	return &biz.ShortenURL{
		ID:       url.ID,
		URLFull:  url.URLFull,
		URLHost:  url.URLHost,
		URLUri:   url.URLURI,
		URLQuery: url.URLQuery,
		URLCode:  url.URLCode,
	}
}
