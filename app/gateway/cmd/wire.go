//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/HarryBird/url-shorten/app/gateway/internal/biz"
	"github.com/HarryBird/url-shorten/app/gateway/internal/conf"
	"github.com/HarryBird/url-shorten/app/gateway/internal/data"
	"github.com/HarryBird/url-shorten/app/gateway/internal/server"
	"github.com/HarryBird/url-shorten/app/gateway/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Data, *conf.App, *conf.Registry, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
	// panic(wire.Build(server.ProviderSet, service.ProviderSet, newApp))
}
