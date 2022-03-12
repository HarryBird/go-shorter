package data

import (
	"context"

	"github.com/HarryBird/url-shorten/app/gateway/internal/conf"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"

	mredis "github.com/HarryBird/mo-kit/cache/goredis"
	grlog "github.com/HarryBird/mo-kit/kratos/log/goredis"
	sv1 "github.com/HarryBird/url-shorten/api/shorten/v1"
	"github.com/go-kratos/kratos/contrib/registry/nacos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	redis "github.com/go-redis/redis/v8"
	"github.com/google/wire"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewShortenServiceClient, NewRedis, NewShortenRepo, NewDiscovery)

func NewDiscovery(conf *conf.Registry) registry.Discovery {
	sc := []constant.ServerConfig{
		*constant.NewServerConfig(
			conf.Nacos.Address,
			conf.Nacos.Port,
			constant.WithScheme(conf.Nacos.Scheme),
			constant.WithContextPath(conf.Nacos.ContextPath)),
	}

	cc := &constant.ClientConfig{
		NamespaceId:         "Service-Registry-Dev", // namespace id
		TimeoutMs:           5000,
		BeatInterval:        5000,
		NotLoadCacheAtStart: true,
		LogDir:              "./tmp/nacos/log",
		CacheDir:            "./tmp/nacos/cache",
		LogRollingConfig: &constant.ClientLogRollingConfig{
			MaxSize:    100,
			MaxAge:     7,
			MaxBackups: 7,
		},
		LogLevel: "info",
	}

	client, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ServerConfigs: sc,
			ClientConfig:  cc,
		},
	)
	if err != nil {
		panic(err)
	}

	return nacos.New(client, nacos.WithCluster("MO"), nacos.WithGroup("MO"))
}

// Data .
type Data struct {
	sc  sv1.ShortenClient
	rdb *redis.Client
	log *log.Helper
}

func NewShortenServiceClient(r registry.Discovery, tp *tracesdk.TracerProvider, logger log.Logger) sv1.ShortenClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		// grpc.WithEndpoint("http://127.0.0.1:9100"),
		// grpc.WithEndpoint("DevLab.urlshorten.shorten:9000"),
		grpc.WithEndpoint("discovery:///urlshorten.service.shorten.grpc"),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			tracing.Client(tracing.WithTracerProvider(tp)),
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
