package server

import (
	"log"

	"github.com/HarryBird/url-shorten/app/gateway/internal/conf"
	"github.com/go-kratos/kratos/contrib/registry/nacos/v2"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewGRPCServer, NewHTTPServer, NewRegistrar)

func NewRegistrar(conf *conf.Registry) registry.Registrar {
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
		log.Panic(err)
	}

	return nacos.New(client, nacos.WithCluster("MO"), nacos.WithGroup("MO"))
}
