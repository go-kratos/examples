package server

import (
	"context"

	"github.com/bufbuild/protovalidate-go"
	v1 "github.com/go-kratos/examples/blog/api/blog/v1"
	"github.com/go-kratos/examples/blog/internal/conf"
	"github.com/go-kratos/examples/blog/internal/service"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/http"
	"google.golang.org/protobuf/proto"
)

func Validator(validator *protovalidate.Validator) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if msg, ok := req.(proto.Message); ok {
				if err := validator.Validate(msg); err != nil {
					if valErr, ok := err.(*protovalidate.ValidationError); ok {
						for _, field := range valErr.Violations {
							return nil, errors.BadRequest("VALIDATOR", field.Message).WithCause(err)
						}
					}

					// 如果proto校验语法写错会报这个错误
					if _, ok := err.(*protovalidate.CompilationError); ok {
						return nil, errors.BadRequest("VALIDATOR", "校验错误").WithCause(err)
					}

				}
			}
			return handler(ctx, req)
		}
	}
}

// NewHTTPServer new a HTTP server.
func NewHTTPServer(validator *protovalidate.Validator, c *conf.Server, logger log.Logger, blog *service.BlogService) *http.Server {
	opts := []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			tracing.Server(),
			logging.Server(logger),
			validate.Validator(),
			Validator(validator), // buf生态Validator
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
	srv := http.NewServer(opts...)
	v1.RegisterBlogServiceHTTPServer(srv, blog)
	return srv
}
