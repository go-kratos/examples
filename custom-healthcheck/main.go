package main

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/status"
	"os"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

type HealthCheckService struct {
	log *log.Helper
}

func NewHealthCheckService(logger log.Logger) *HealthCheckService {
	return &HealthCheckService{
		log: log.NewHelper(logger),
	}
}

func (s *HealthCheckService) Check(ctx context.Context, req *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	s.log.WithContext(ctx).Infof("HealthCheckService.Check: %v", req)

	// TODO: add your health check logic here

	return &grpc_health_v1.HealthCheckResponse{
		Status: grpc_health_v1.HealthCheckResponse_SERVING,
	}, nil
}

func (s *HealthCheckService) Watch(req *grpc_health_v1.HealthCheckRequest, srv grpc_health_v1.Health_WatchServer) error {
	s.log.Infof("HealthCheckService.Watch: %v, %v", req, srv)

	// simple implementation
	resp := &grpc_health_v1.HealthCheckResponse{
		Status: grpc_health_v1.HealthCheckResponse_SERVING,
	}
	if err := srv.Send(resp); err != nil {
		return status.Errorf(codes.Unknown, "failed to send response: %v", err)
	}
	return nil
}

func main() {
	grpcSrv := grpc.NewServer(
		grpc.Address(":9000"),
		grpc.Middleware(),
		grpc.CustomHealth(),
	)
	app := kratos.New(
		kratos.Name("helloworld"),
		kratos.Server(
			grpcSrv,
		),
	)
	logger := log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
	)
	health := NewHealthCheckService(logger)
	grpc_health_v1.RegisterHealthServer(grpcSrv, health)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
