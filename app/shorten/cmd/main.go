package main

import (
	"flag"
	"os"

	"github.com/HarryBird/url-shorten/app/shorten/internal/conf"

	moZap "github.com/HarryBird/mo-kit/log/zap"
	moJaeger "github.com/HarryBird/mo-kit/trace/jaeger"
	zlog "github.com/go-kratos/kratos/contrib/log/zap/v2"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/env"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name = "urlshorten.service.shorten"
	// Version is the version of the compiled software.
	Version string
	// flagconf is the config flag.
	flagconf string

	id, _ = os.Hostname()
)

const (
	EnvVarPrefix = "MO_"
)

func init() {
	flag.StringVar(&flagconf, "conf", "../config", "config path, eg: -conf config.yaml")
}

func newApp(logger log.Logger, gs *grpc.Server, rr registry.Registrar) *kratos.App {
	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			gs,
		),
		kratos.Registrar(rr),
	)
}

func main() {
	flag.Parse()

	logger := log.With(zlog.NewLogger(moZap.DevelopmentLogger()),
		// "ts", log.DefaultTimestamp,
		// "caller", log.DefaultCaller,
		"service.id", id,
		"service.name", Name,
		"service.version", Version,
		"trace_id", tracing.TraceID(),
		"span_id", tracing.SpanID(),
	)

	slog := log.NewHelper(log.With(logger, "mod", "main.bootstrap"))
	slog.Infof("%s service starting...", Name)

	// init config
	c := config.New(
		config.WithSource(
			env.NewSource(EnvVarPrefix),
			file.NewSource(flagconf),
		),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	slog.Infof("with config: server=%+v, data=%+v, registry=%+v, app=%+v", bc.Server, bc.Data, bc.Registry, bc.App)

	// init tracing
	tp, err := moJaeger.DefaultProvider(Name, bc.Trace)
	if err != nil {
		panic(err)
	}

	// init app
	app, cleanup, err := initApp(bc.Server, bc.Data, bc.Registry, bc.App, tp, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}
