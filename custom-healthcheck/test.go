package main

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/health/grpc_health_v1"
)

func main() {
	// init grpc conn
	conn, err := grpc.Dial(":9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	healthClient := grpc_health_v1.NewHealthClient(conn)
	checkReq := &grpc_health_v1.HealthCheckRequest{
		Service: "helloworld.Greeter",
	}
	ctx := context.Background()
	checkResp, err := healthClient.Check(ctx, checkReq)
	if err != nil {
		panic(err)
	}
	log.Info("checkResp: ", checkResp)

	watchReq := &grpc_health_v1.HealthCheckRequest{
		Service: "helloworld.Greeter",
	}
	watchResp, err := healthClient.Watch(ctx, watchReq)
	if err != nil {
		panic(err)
	}
	// loop watch
	for {
		resp, err := watchResp.Recv()
		if err != nil {
			panic(err)
		}
		log.Info("watchResp: ", resp)
	}
}
