package server

import (
	"fmt"
	stdhttp "net/http"
	"strings"

	v1 "github.com/HarryBird/url-shorten/api/gateway/v1"
	"github.com/HarryBird/url-shorten/app/gateway/internal/conf"
	"github.com/HarryBird/url-shorten/app/gateway/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, serv *service.GatewayService, logger log.Logger) *http.Server {
	opts := []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			logging.Server(logger),
			validate.Validator(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}

	// opts = append(opts, http.ResponseEncoder(RedirectResponseEncoder))

	srv := http.NewServer(opts...)
	v1.RegisterGatewayHTTPServer(srv, serv)
	return srv
}

func RedirectResponseEncoder(w stdhttp.ResponseWriter, r *stdhttp.Request, v interface{}) error {
	fmt.Println("========== RedirectResponseEncoder")
	fmt.Printf("val: %#v\n", v)
	if strings.HasPrefix(r.RequestURI, "/v1/url/decode/") {
		fmt.Println("111")
		if reply, ok := v.(*v1.DecodeURLReply); ok {
			fmt.Println("222")
			stdhttp.Redirect(w, r, reply.Url, 302)
			return nil
		}
	}

	return http.DefaultResponseEncoder(w, r, v)
}
