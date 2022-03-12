package server

import (
	v1 "github.com/HarryBird/url-shorten/api/shorten/v1"

	"github.com/HarryBird/url-shorten/app/shorten/internal/conf"
	"github.com/HarryBird/url-shorten/app/shorten/internal/service"

	prom "github.com/go-kratos/kratos/contrib/metrics/prometheus/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/metrics"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Server, tp *tracesdk.TracerProvider, serv *service.ShortenService, logger log.Logger) *grpc.Server {
	opts := []grpc.ServerOption{
		grpc.Middleware(
			tracing.Server(
				tracing.WithTracerProvider(tp)),
			recovery.Recovery(),
			logging.Server(logger),
			validate.Validator(),
			metrics.Server(
				metrics.WithSeconds(prom.NewHistogram(_metricSeconds)),
				metrics.WithRequests(prom.NewCounter(_metricRequests)),
			),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	v1.RegisterShortenServer(srv, serv)
	return srv
}
