package biz

import (
	"context"
	"fmt"
	"hash/crc64"
	stdurl "net/url"
	"strings"

	"github.com/HarryBird/mo-kit/msgr"
	"github.com/go-kratos/kratos/v2/log"
	b62 "github.com/jxskiss/base62"
	"github.com/pkg/errors"
	m3 "github.com/spaolacci/murmur3"
)

type ShortenRepo interface {
	Create(ctx context.Context, url *ShortenURL) (*ShortenURL, error)
	Decode(ctx context.Context, url *ShortenURL) (*ShortenURL, error)
}

type ShortenCase struct {
	repo ShortenRepo
	log  *log.Helper
}

func NewShortenCase(repo ShortenRepo, logger log.Logger) *ShortenCase {
	return &ShortenCase{repo: repo, log: log.NewHelper(log.With(logger, "mod", "biz.shorten"))}
}

func (uc *ShortenCase) Decode(ctx context.Context, surl *ShortenURL) (*ShortenURL, error) {
	return uc.repo.Decode(ctx, surl)
}

func (uc *ShortenCase) Create(ctx context.Context, ourl OriginURL) (*ShortenURL, error) {
	fname := "Create"
	uc.log.WithContext(ctx).Infof("%s origin url: %s", msgr.W(fname), ourl.Url)

	surl, err := uc.shorten(ctx, ourl.Url)
	if err != nil {
		return nil, errors.WithMessage(err, "biz: gen shorten url fail")
	}

	uc.log.WithContext(ctx).Infof("%s after shorten: %+v", msgr.W(fname), surl)

	surl, err = uc.repo.Create(ctx, surl)
	if err != nil {
		return nil, errors.WithMessage(err, "biz: create shorten url fail")
	}

	return surl, nil
}

func (uc *ShortenCase) shorten(ctx context.Context, url string) (*ShortenURL, error) {
	u, err := stdurl.Parse(url)

	if err != nil {
		return nil, errors.WithMessage(err, "biz: url parse fail")
	}

	surl := &ShortenURL{
		URLFull: url,
		URLHost: u.Hostname(),
		URLCode: uc.hash(ctx, url),
	}

	uriParts := strings.SplitN(u.RequestURI(), "?", 2)

	surl.URLUri = uriParts[0]
	if len(uriParts) > 1 {
		surl.URLQuery = uriParts[1]
	}

	return surl, nil
}

func (uc *ShortenCase) hash(ctx context.Context, url string) string {
	fname := "hash"

	murHash := uc.murHash(ctx, url)
	crcHash := uc.crcHash(ctx, url)
	urlHash := fmt.Sprintf("%s.%s", murHash, crcHash)

	uc.log.WithContext(ctx).Infof("%s after hash: murHash=%v, crcHash=%v, urlHash=%v",
		msgr.W(fname), murHash, crcHash, urlHash)
	return urlHash
}

func (uc *ShortenCase) murHash(ctx context.Context, url string) string {
	hasher := m3.New64()
	hasher.Write([]byte(url))
	hashSum := hasher.Sum64()
	return string(b62.FormatUint(hashSum))
}

func (uc *ShortenCase) crcHash(ctx context.Context, url string) string {
	hashSum := crc64.Checksum([]byte(url), crc64.MakeTable(crc64.ISO))
	// 62^4 = 14776336 对 10000000 取模再转base62，使hash为4位62进制字符
	return string(b62.FormatUint(hashSum % 10000000))
}
