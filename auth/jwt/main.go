package main

import (
	"context"
	"log"

	"github.com/go-kratos/examples/helloworld/helloworld"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	jwtv5 "github.com/golang-jwt/jwt/v5"
)

type server struct {
	helloworld.UnimplementedGreeterServer

	hc helloworld.GreeterClient
}

func (s *server) SayHello(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	return &helloworld.HelloReply{Message: "hello from service"}, nil
}

func main() {
	testKey := "testKey"
	httpSrv := http.NewServer(
		http.Address(":8000"),
		http.Middleware(
			jwt.Server(func(token *jwtv5.Token) (interface{}, error) {
				return []byte(testKey), nil
			}),
		),
	)
	grpcSrv := grpc.NewServer(
		grpc.Address(":9000"),
		grpc.Middleware(
			jwt.Server(func(token *jwtv5.Token) (interface{}, error) {
				return []byte(testKey), nil
			}),
		),
	)
	serviceTestKey := "serviceTestKey"
	con, _ := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("dns:///127.0.0.1:9001"),
		grpc.WithMiddleware(
			jwt.Client(func(token *jwtv5.Token) (interface{}, error) {
				return []byte(serviceTestKey), nil
			}),
		),
	)
	s := &server{
		hc: helloworld.NewGreeterClient(con),
	}
	helloworld.RegisterGreeterServer(grpcSrv, s)
	helloworld.RegisterGreeterHTTPServer(httpSrv, s)
	app := kratos.New(
		kratos.Name("helloworld"),
		kratos.Server(
			httpSrv,
			grpcSrv,
		),
	)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
