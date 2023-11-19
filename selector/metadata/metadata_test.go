package metadata

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/middleware"
	"testing"
	"time"

	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/metadata"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/selector"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/hashicorp/consul/api"

	"github.com/go-kratos/examples/helloworld/helloworld"
)

func TestMetadataNodeFilter(t *testing.T) {
	region := "bj"
	cli, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		t.Fatal(err)
	}

	discovery := consul.New(cli, consul.WithHealthCheck(false), consul.WithHeartbeat(false))

	// region bj
	bjapp := runServer("bj", discovery)
	// region hz
	hzapp := runServer("hz", discovery)

	go bjapp.Run()
	go hzapp.Run()

	time.Sleep(5 * time.Second)

	cc, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithDiscovery(discovery),
		grpc.WithEndpoint("discovery:///helloworld"),
		grpc.WithNodeFilter(filter(region)),
		grpc.WithMiddleware(func(handler middleware.Handler) middleware.Handler {
			return func(ctx context.Context, req interface{}) (interface{}, error) {
				reply, err := handler(ctx, req)
				if p, ok := selector.FromPeerContext(ctx); ok {
					t.Log(p.Node.Address(), p.Node.Metadata())
				}
				return reply, err
			}
		}),
	)
	if err != nil {
		t.Fatal(err)
	}

	client := helloworld.NewGreeterClient(cc)

	ctx := context.Background()
	// priority same region request
	reply, err := client.SayHello(ctx, &helloworld.HelloRequest{Name: "kratos"})
	if err != nil {
		t.Fatal(err)
	}
	peer, ok := selector.FromPeerContext(ctx)
	if ok && peer.Node.Metadata()["region"] != region {
		t.Fatal("node filter result error", peer.Node.Address())
	}
	t.Log(reply)

	// ctx specified availability region
	ctx1 := metadata.NewClientContext(context.Background(), metadata.Metadata{"region": []string{"hz"}})
	reply, err = client.SayHello(ctx1, &helloworld.HelloRequest{Name: "kratos"})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(reply)

	peer, ok = selector.FromPeerContext(ctx1)
	if ok && peer.Node.Metadata()["region"] != "hz" {
		t.Fatal("node filter result error", peer.Node.Address())
	}

	// stop bj app test cross region degradation
	if err = bjapp.Stop(); err != nil {
		t.Fatal(err)
	}

	// ctx specified availability region
	ctx2 := context.Background()
	reply, err = client.SayHello(ctx2, &helloworld.HelloRequest{Name: "kratos"})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(reply)

	peer, ok = selector.FromPeerContext(ctx2)
	if ok && peer.Node.Metadata()["region"] != "hz" {
		t.Fatal("node filter result error", peer.Node.Address())
	}
}

type server struct {
	helloworld.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, request *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	return &helloworld.HelloReply{Message: "hello"}, nil
}

func runServer(region string, registrar registry.Registrar) *kratos.App {
	grpcSrv := grpc.NewServer(
		grpc.Address(fmt.Sprintf(":0")), // rand port
		grpc.Middleware(
			recovery.Recovery(),
			logging.Server(log.GetLogger()),
		),
	)
	s := &server{}
	helloworld.RegisterGreeterServer(grpcSrv, s)

	return kratos.New(
		kratos.ID(region),
		kratos.Name("helloworld"),
		kratos.Metadata(map[string]string{
			"region": region, // region info
		}),
		kratos.Server(grpcSrv),
		kratos.Registrar(registrar),
	)
}

func filter(region string) selector.NodeFilter {
	return func(ctx context.Context, nodes []selector.Node) []selector.Node {
		if v, ok := metadata.FromClientContext(ctx); ok {
			region = v.Get("region") // if a region is specified in the request, use specified in the request
		}

		newNodes := make([]selector.Node, 0, len(nodes))
		for _, node := range nodes {
			if node.Metadata()["region"] == region {
				newNodes = append(newNodes, node)
			}
		}

		if len(newNodes) != 0 {
			return newNodes
		}

		return nodes
	}
}
