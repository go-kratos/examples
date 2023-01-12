package main

import (
	"context"
	"fmt"
	"os"

	"github.com/go-kratos/kratos/contrib/polaris/v2"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	polaris2 "github.com/polarismesh/polaris-go"

	"github.com/go-kratos/examples/helloworld/helloworld"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	helloworld.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	if s, ok := kratos.FromContext(ctx); ok {
		return &helloworld.HelloReply{Message: fmt.Sprintf("Welcome %+v!", s.Metadata())}, nil
	}
	return &helloworld.HelloReply{Message: fmt.Sprintf("Welcome %+v!", in.Name)}, nil
}

func main() {
	logger := log.NewStdLogger(os.Stdout)
	log := log.NewHelper(logger)

	sdk, err := polaris2.NewSDKContextByAddress("127.0.0.1:8091")
	if err != nil {
		log.Fatal(err)
	}

	p := polaris.New(sdk)

	pc, err := p.Config(polaris.WithConfigFile(polaris.File{Name: "config.yaml", Group: "test"}))
	if err != nil {
		log.Fatal(err)
	}

	cfg := config.New(config.WithSource(pc))

	if err := cfg.Load(); err != nil {
		log.Fatal(err)
	}

	type conf struct {
		Name string `yaml:"name"`
		HTTP struct {
			Addr string `yaml:"addr"`
		}
		GRPC struct {
			Addr string `yaml:"addr"`
		}
	}

	c := conf{}

	if err := cfg.Scan(&c); err != nil {
		log.Fatal(err)
	}

	httpSrv := http.NewServer(
		http.Address(c.HTTP.Addr),
		http.Middleware(
			recovery.Recovery(),
			logging.Server(logger),
		),
	)
	grpcSrv := grpc.NewServer(
		grpc.Address(c.GRPC.Addr),
		grpc.Middleware(
			recovery.Recovery(),
			logging.Server(logger),
		),
	)

	s := &server{}
	helloworld.RegisterGreeterServer(grpcSrv, s)
	helloworld.RegisterGreeterHTTPServer(httpSrv, s)

	app := kratos.New(
		kratos.Name(c.Name),
		kratos.Server(
			grpcSrv,
			httpSrv,
		),
		kratos.Registrar(p.Registry(polaris.WithRegistryTTL(40))),
		kratos.Metadata(map[string]string{
			"az": "bj3",
		}),
	)

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
