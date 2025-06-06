package bootstrap

import (
	"github.com/go-kratos/aegis/ratelimit"
	"github.com/go-kratos/aegis/ratelimit/bbr"
	"github.com/go-kratos/kratos/v2/middleware"
	midRateLimit "github.com/go-kratos/kratos/v2/middleware/ratelimit"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	kratosRest "github.com/go-kratos/kratos/v2/transport/http"

	"github.com/gorilla/handlers"

	"kratos-gorm-example/gen/api/go/common/conf"
)

// CreateRestServer 创建REST服务端
func CreateRestServer(cfg *conf.Bootstrap, m ...middleware.Middleware) *kratosRest.Server {
	var opts = []kratosRest.ServerOption{
		kratosRest.Filter(handlers.CORS(
			handlers.MaxAge(3600),
			handlers.AllowedHeaders(cfg.Server.Rest.Cors.Headers),
			handlers.AllowedMethods(cfg.Server.Rest.Cors.Methods),
			handlers.AllowedOrigins(cfg.Server.Rest.Cors.Origins),
		)),
	}

	var ms []middleware.Middleware
	if cfg.Server != nil && cfg.Server.Rest != nil && cfg.Server.Rest.Middleware != nil {
		if cfg.Server.Rest.Middleware.GetEnableRecovery() {
			ms = append(ms, recovery.Recovery())
		}
		if cfg.Server.Rest.Middleware.GetEnableTracing() {
			ms = append(ms, tracing.Server())
		}
		if cfg.Server.Rest.Middleware.GetEnableValidate() {
			ms = append(ms, validate.Validator())
		}
		if cfg.Server.Rest.Middleware.GetEnableCircuitBreaker() {
		}
		if cfg.Server.Rest.Middleware.Limiter != nil {
			var limiter ratelimit.Limiter
			switch cfg.Server.Rest.Middleware.Limiter.GetName() {
			case "bbr":
				limiter = bbr.NewLimiter()
			}
			ms = append(ms, midRateLimit.Server(midRateLimit.WithLimiter(limiter)))
		}
	}
	ms = append(ms, m...)
	opts = append(opts, kratosRest.Middleware(ms...))

	if cfg.Server.Rest.Network != "" {
		opts = append(opts, kratosRest.Network(cfg.Server.Rest.Network))
	}
	if cfg.Server.Rest.Addr != "" {
		opts = append(opts, kratosRest.Address(cfg.Server.Rest.Addr))
	}
	if cfg.Server.Rest.Timeout != nil {
		opts = append(opts, kratosRest.Timeout(cfg.Server.Rest.Timeout.AsDuration()))
	}

	return kratosRest.NewServer(opts...)
}
