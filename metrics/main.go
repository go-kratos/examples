package main

import (
	"context"
	"fmt"
	"log"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"

	"github.com/go-kratos/examples/helloworld/helloworld"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/middleware/metrics"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name = "metrics"
	// Version is the version of the compiled software.
	// Version = "v1.0.0"

	meter = otel.Meter(Name)

	_metricRequests metric.Int64Counter
	_metricSeconds  metric.Float64Histogram
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	helloworld.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	return &helloworld.HelloReply{Message: fmt.Sprintf("Hello %+v", in.Name)}, nil
}

func init() {
	var err error
	_metricRequests, err = metrics.DefaultRequestsCounter(meter, metrics.DefaultServerRequestsCounterName)
	if err != nil {
		panic(err)
	}

	_metricSeconds, err = metrics.DefaultSecondsHistogram(meter, metrics.DefaultServerSecondsHistogramName)
	if err != nil {
		panic(err)
	}
}

func main() {
	grpcSrv := grpc.NewServer(
		grpc.Address(":9000"),
		grpc.Middleware(
			metrics.Server(
				metrics.WithSeconds(_metricSeconds),
				metrics.WithRequests(_metricRequests),
			),
		),
	)
	httpSrv := http.NewServer(
		http.Address(":8000"),
		http.Middleware(
			metrics.Server(
				metrics.WithSeconds(_metricSeconds),
				metrics.WithRequests(_metricRequests),
			),
		),
	)

	s := &server{}
	helloworld.RegisterGreeterServer(grpcSrv, s)
	helloworld.RegisterGreeterHTTPServer(httpSrv, s)

	app := kratos.New(
		kratos.Name(Name),
		kratos.Server(
			httpSrv,
			grpcSrv,
		),
	)

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
