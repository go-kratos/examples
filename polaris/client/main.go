package main

import (
	"context"
	"log"
	"time"

	"github.com/go-kratos/kratos/contrib/polaris/v2"
	"github.com/go-kratos/kratos/v2"
	md "github.com/go-kratos/kratos/v2/metadata"
	"github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	polaris2 "github.com/polarismesh/polaris-go"

	"github.com/go-kratos/examples/helloworld/helloworld"
)

func main() {
	sdk, err := polaris2.NewSDKContextByAddress("127.0.0.1:8091")
	if err != nil {
		log.Fatal(err)
	}

	p := polaris.New(sdk)

	// new grpc client
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///helloworld"),
		grpc.WithDiscovery(p.Registry()),
		grpc.WithNodeFilter(p.NodeFilter(polaris.WithService("helloworld"))),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	gClient := helloworld.NewGreeterClient(conn)

	// new http client
	hConn, err := http.NewClient(
		context.Background(),
		http.WithMiddleware(
			recovery.Recovery(),
			metadata.Client(),
		),
		http.WithEndpoint("discovery:///helloworld"),
		http.WithDiscovery(p.Registry()),
		http.WithNodeFilter(p.NodeFilter()),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer hConn.Close()
	hClient := helloworld.NewGreeterHTTPClient(hConn)

	for {
		time.Sleep(time.Second)
		callGRPC(gClient)
		callHTTP(hClient)
	}
}

func callGRPC(client helloworld.GreeterClient) {
	reply, err := client.SayHello(context.Background(), &helloworld.HelloRequest{Name: "kratos"})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("[grpc] SayHello %+v\n", reply)
}

type mock struct {
}

func (m mock) ID() string {
	return "123"
}

func (m mock) Name() string {
	return "kratos"
}

func (m mock) Version() string {
	return "1.0.0"
}

func (m mock) Metadata() map[string]string {
	return map[string]string{}
}

func (m mock) Endpoint() []string {
	return []string{}
}

func callHTTP(client helloworld.GreeterHTTPClient) {

	ctx := md.NewClientContext(kratos.NewContext(context.Background(), mock{}), md.Metadata{"az": "bj1"})
	reply, err := client.SayHello(ctx, &helloworld.HelloRequest{Name: "kratos"})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("[http] SayHello %s\n", reply.Message)
}
