package bootstrap

import (
	"context"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/durationpb"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/registry"
	kratosGrpc "github.com/go-kratos/kratos/v2/transport/grpc"

	"kratos-gorm-example/gen/api/go/common/conf"
)

const defaultTimeout = 5 * time.Second

// CreateGrpcClient 创建GRPC客户端
func CreateGrpcClient(ctx context.Context, r registry.Discovery, serviceName string, timeoutDuration *durationpb.Duration) grpc.ClientConnInterface {
	timeout := defaultTimeout
	if timeoutDuration != nil {
		timeout = timeoutDuration.AsDuration()
	}

	endpoint := "discovery:///" + serviceName

	conn, err := kratosGrpc.DialInsecure(
		ctx,
		kratosGrpc.WithEndpoint(endpoint),
		kratosGrpc.WithDiscovery(r),
		kratosGrpc.WithTimeout(timeout),
		kratosGrpc.WithMiddleware(
			tracing.Client(),
			recovery.Recovery(),
		),
	)
	if err != nil {
		log.Fatalf("dial grpc client [%s] failed: %s", serviceName, err.Error())
	}

	return conn
}

// CreateGrpcServer 创建GRPC服务端
func CreateGrpcServer(cfg *conf.Bootstrap, m ...middleware.Middleware) *kratosGrpc.Server {
	var opts []kratosGrpc.ServerOption

	var ms []middleware.Middleware
	ms = append(ms, recovery.Recovery())
	ms = append(ms, tracing.Server())
	ms = append(ms, m...)
	opts = append(opts, kratosGrpc.Middleware(ms...))

	if cfg.Server.Grpc.Network != "" {
		opts = append(opts, kratosGrpc.Network(cfg.Server.Grpc.Network))
	}
	if cfg.Server.Grpc.Addr != "" {
		opts = append(opts, kratosGrpc.Address(cfg.Server.Grpc.Addr))
	}
	if cfg.Server.Grpc.Timeout != nil {
		opts = append(opts, kratosGrpc.Timeout(cfg.Server.Grpc.Timeout.AsDuration()))
	}

	return kratosGrpc.NewServer(opts...)
}
