package server

import (
	moNacos "github.com/HarryBird/mo-kit/registry/nacos"
	"github.com/HarryBird/url-shorten/app/gateway/internal/conf"
	"github.com/go-kratos/kratos/contrib/registry/nacos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewGRPCServer, NewHTTPServer, NewRegistrar, NewDiscover)

func NewRegistrar(regCfg *conf.Registry, appCfg *conf.App, logger log.Logger) registry.Registrar {
	rlog := log.NewHelper(log.With(logger, "mod", "server.registry"))
	client, servCfg, clientCfg, err := moNacos.DefaultClient(regCfg.Nacos, appCfg.Runtime)
	rlog.Infof("registry config: serv_config=%+v, client_config=%v", servCfg, clientCfg)

	if err != nil {
		rlog.Fatalf("server: new nacos client failed %v", err)
	}

	return nacos.New(client, nacos.WithCluster(moNacos.ClusterBeijing), nacos.WithGroup(moNacos.GroupDefault))
}

func NewDiscover(regCfg *conf.Registry, appCfg *conf.App, logger log.Logger) registry.Discovery {
	rlog := log.NewHelper(log.With(logger, "mod", "server.discover"))
	client, servCfg, clientCfg, err := moNacos.DefaultClient(regCfg.Nacos, appCfg.Runtime)
	rlog.Infof("discover config: serv_config=%+v, client_config=%v", servCfg, clientCfg)

	if err != nil {
		rlog.Fatalf("server: new nacos client failed %v", err)
	}

	return nacos.New(client, nacos.WithCluster(moNacos.ClusterBeijing), nacos.WithGroup(moNacos.GroupDefault))
}
