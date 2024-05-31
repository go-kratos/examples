package auth

import (
	"context"
	"strconv"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"

	authnEngine "github.com/tx7do/kratos-authn/engine"
	authn "github.com/tx7do/kratos-authn/middleware"

	authzEngine "github.com/tx7do/kratos-authz/engine"
	authz "github.com/tx7do/kratos-authz/middleware"
)

var action = authzEngine.Action("ANY")

// Server 衔接认证和权鉴
func Server() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			tr, ok := transport.FromServerContext(ctx)
			if !ok {
				return nil, ErrWrongContext
			}

			authnClaims, ok := authn.FromContext(ctx)
			if !ok {
				return nil, ErrWrongContext
			}

			sub := authzEngine.Subject(authnClaims.Subject)
			path := authzEngine.Resource(tr.Operation())
			authzClaims := authzEngine.AuthClaims{
				Subject:  &sub,
				Action:   &action,
				Resource: &path,
			}

			ctx = authz.NewContext(ctx, &authzClaims)

			return handler(ctx, req)
		}
	}
}

type Result struct {
	UserId   uint32
	UserName string
}

func FromContext(ctx context.Context) (*Result, error) {
	claims, ok := authnEngine.AuthClaimsFromContext(ctx)
	if !ok {
		return nil, ErrMissingJwtToken
	}

	userId, err := strconv.ParseUint(claims.Subject, 10, 32)
	if err != nil {
		return nil, ErrExtractSubjectFailed
	}

	return &Result{
		UserId:   uint32(userId),
		UserName: "",
	}, nil
}
