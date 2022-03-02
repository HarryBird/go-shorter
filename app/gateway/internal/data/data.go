package data

import (
	"context"

	"github.com/HarryBird/url-shorten/app/gateway/internal/conf"

	mredis "github.com/HarryBird/mo-kit/cache/goredis"
	grlog "github.com/HarryBird/mo-kit/kratos/log/goredis"
	sv1 "github.com/HarryBird/url-shorten/api/shorten/v1"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	redis "github.com/go-redis/redis/v8"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewShortenServiceClient, NewRedis, NewShortenRepo)

// Data .
type Data struct {
	sc  sv1.ShortenClient
	rdb *redis.Client
	log *log.Helper
}

func NewShortenServiceClient(logger log.Logger) sv1.ShortenClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		// grpc.WithEndpoint("http://127.0.0.1:9100"),
		grpc.WithEndpoint("127.0.0.1:9100"),
		grpc.WithMiddleware(
			recovery.Recovery(),
			logging.Client(logger),
			validate.Validator(),
		),
	)
	if err != nil {
		panic(err)
	}
	return sv1.NewShortenClient(conn)
}

func NewRedis(conf *conf.Data, logger log.Logger) *redis.Client {
	rlog := log.NewHelper(log.With(logger, "mod", "repo.redis"))
	rdb, err := mredis.NewRedis(conf.Redis)
	if err != nil {
		rlog.Fatalf("repo: failed opening connection to redis %v", err)
	}

	rlog.Infof("repo: redis connection options: %+v", rdb.Options())

	redis.SetLogger(grlog.New(logger))

	return rdb
}

// NewData .
func NewData(rdb *redis.Client, sc sv1.ShortenClient, logger log.Logger) (*Data, func(), error) {
	log := log.NewHelper(log.With(logger, "mod", "repo.data"))

	d := &Data{
		sc:  sc,
		rdb: rdb,
		log: log,
	}
	return d, func() {
	}, nil
}
