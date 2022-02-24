// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/HarryBird/url-shorten/app/gateway/internal/biz"
	"github.com/HarryBird/url-shorten/app/gateway/internal/conf"
	"github.com/HarryBird/url-shorten/app/gateway/internal/data"
	"github.com/HarryBird/url-shorten/app/gateway/internal/server"
	"github.com/HarryBird/url-shorten/app/gateway/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

// Injectors from wire.go:

// initApp init kratos application.
func initApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	client := data.NewRedis(confData, logger)
	shortenClient := data.NewShortenServiceClient()
	dataData, cleanup, err := data.NewData(client, shortenClient, logger)
	if err != nil {
		return nil, nil, err
	}
	shortenRepo := data.NewShortenRepo(dataData, logger)
	shortenCase := biz.NewShortenCase(shortenRepo, logger)
	gatewayService := service.NewGatewayService(shortenCase, logger)
	httpServer := server.NewHTTPServer(confServer, gatewayService, logger)
	grpcServer := server.NewGRPCServer(confServer, gatewayService, logger)
	app := newApp(logger, httpServer, grpcServer)
	return app, func() {
		cleanup()
	}, nil
}
