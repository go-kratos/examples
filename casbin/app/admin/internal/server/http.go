package server

import (
	"context"
	"github.com/casbin/casbin/v2/model"
	fileAdapter "github.com/casbin/casbin/v2/persist/file-adapter"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-kratos/swagger-api/openapiv2"
	jwtV4 "github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/handlers"
	"kratos-casbin/api/admin/v1"
	"kratos-casbin/app/admin/internal/conf"

	casbinM "github.com/tx7do/kratos-casbin/authz/casbin"
	myAuthz "kratos-casbin/app/admin/internal/pkg/authz"
	"kratos-casbin/app/admin/internal/service"
)

// NewWhiteListMatcher 创建jwt白名单
func NewWhiteListMatcher() selector.MatchFunc {
	whiteList := make(map[string]struct{})
	whiteList["/admin.v1.AdminService/Login"] = struct{}{}
	whiteList["/admin.v1.AdminService/Logout"] = struct{}{}
	whiteList["/admin.v1.AdminService/Register"] = struct{}{}
	whiteList["/admin.v1.AdminService/GetPublicContent"] = struct{}{}
	return func(ctx context.Context, operation string) bool {
		if _, ok := whiteList[operation]; ok {
			return false
		}
		return true
	}
}

// NewMiddleware 创建中间件
func NewMiddleware(ac *conf.Auth, logger log.Logger) http.ServerOption {
	m, _ := model.NewModelFromFile("../../configs/authz/authz_model.conf")
	a := fileAdapter.NewAdapter("../../configs/authz/authz_policy.csv")

	return http.Middleware(
		recovery.Recovery(),
		tracing.Server(),
		logging.Server(logger),
		selector.Server(
			jwt.Server(
				func(token *jwtV4.Token) (interface{}, error) {
					return []byte(ac.ApiKey), nil
				},
				jwt.WithSigningMethod(jwtV4.SigningMethodHS256),
			),
			casbinM.Server(
				casbinM.WithCasbinModel(m),
				casbinM.WithCasbinPolicy(a),
				casbinM.WithSecurityUserCreator(myAuthz.NewSecurityUser),
			),
		).
			Match(NewWhiteListMatcher()).Build(),
	)
}

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, ac *conf.Auth, logger log.Logger, s *service.AdminService) *http.Server {
	var opts = []http.ServerOption{
		NewMiddleware(ac, logger),
		http.Filter(handlers.CORS(
			handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"}),
			handlers.AllowedOrigins([]string{"*"}),
		)),
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
	srv := http.NewServer(opts...)

	h := openapiv2.NewHandler()
	srv.HandlePrefix("/q/", h)

	v1.RegisterAdminServiceHTTPServer(srv, s)
	return srv
}
