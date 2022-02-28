package data

import (
	"context"

	sv1 "github.com/HarryBird/url-shorten/api/shorten/v1"
	"github.com/HarryBird/url-shorten/app/gateway/internal/biz"
	redis "github.com/go-redis/redis/v8"
	"github.com/pkg/errors"

	mlog "github.com/HarryBird/mo-kit/kratos/log/app"
	"github.com/HarryBird/mo-kit/msgr"
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

	// try load from cache
	cacheReply, err := r.decodeByCache(ctx, in)

	// redis key 不存在，rpc 回源
	if errors.Is(err, biz.ErrNotFoundFromRedis) {
		key, ttl := RedisKeyShortenCodeToURL.extract(in.Code)

		rpcReq := &sv1.GetShortenURLRequest{Query: &sv1.GetShortenURLRequest_Code{Code: in.Code}}
		rpcReply, err := r.data.sc.GetShortenURL(ctx, rpcReq)
		r.log.WithContext(ctx).Debugf("%s call shorter rpc: req=%v, reply=%v", msgr.W(fname), rpcReq, rpcReply)

		if err != nil {
			if sv1.IsShortenUrlNonexist(err) {
				// rpc 查询不存在，写入空值进cache
				r.log.WithContext(ctx).Infof("%s write empty to redis: key=%s, val=\"\", ttl=%v", msgr.W(fname), key, ttl)
				if err := r.data.rdb.SetEX(ctx, key, "", ttl).Err(); err != nil {
					r.log.WithContext(ctx).Errorf("%s write empty set to redis fail: err=%v", msgr.W(fname), err)
				}
				return nil, errors.WithMessage(biz.ErrURLCodeNonexist, "repo: [rpc] shorten url code not found")
			}
			mlog.LogErrorRPC(ctx, r.log, fname, err)
			return nil, errors.WithMessage(err, "repo: [rpc] decode shorten url fail by shorter service")
		}

		// rpc 查询结果写入cache
		r.log.WithContext(ctx).Infof("%s write ful url to redis: key=%s, val=%s, ttl=%v", msgr.W(fname), key, rpcReply.ShortenUrl.UrlFull, ttl)
		if err := r.data.rdb.SetEX(ctx, key, rpcReply.ShortenUrl.UrlFull, ttl).Err(); err != nil {
			r.log.WithContext(ctx).Errorf("%s write full url to redis fail: err=%v", msgr.W(fname), err)
		}

		return &biz.DecodeOut{URL: rpcReply.ShortenUrl.UrlFull}, nil
	}

	return cacheReply, err
}

func (r *shortenRepo) decodeByCache(ctx context.Context, in *biz.DecodeIn) (*biz.DecodeOut, error) {
	fname := "decodeByCache"

	key, _ := RedisKeyShortenCodeToURL.extract(in.Code)
	r.log.WithContext(ctx).Infof("%s try query from redis: key=%s", msgr.W(fname), key)
	v, err := r.data.rdb.Get(ctx, key).Result()

	// key 存在
	if err == nil {
		r.log.WithContext(ctx).Infof("%s query from redis: result=%s", msgr.W(fname), v)
		// value 为空，shorten code 不存在
		if v == "" {
			return nil, errors.WithMessage(biz.ErrURLCodeNonexist, "repo: [redis] shorten url code not exist")
		}
		return &biz.DecodeOut{URL: v}, nil
	}

	// key 不存在，需回源
	if err == redis.Nil {
		return nil, errors.WithMessage(biz.ErrNotFoundFromRedis, "repo: [redis] shorten url code not found")
	}

	r.log.WithContext(ctx).Errorf("%s query from redis fail: err=%v", msgr.W(fname), err)
	return nil, errors.WithMessage(err, "repo: [redis] query shorten url fail")
}
