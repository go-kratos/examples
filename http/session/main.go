package main

import (
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-redis/redis/v8"

	"github.com/go-kratos/examples/helloworld/helloworld"
	"github.com/go-kratos/examples/http/session/sessions"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name = "session"
	// Version is the version of the compiled software.
	// Version = "v1.0.0"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	helloworld.UnimplementedGreeterServer

	store *sessions.RedisStore
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	if tr, ok := transport.FromServerContext(ctx); ok {
		if ht, ok := tr.(*http.Transport); ok {
			// get a session
			session, err := s.store.Get(ht, "session")
			if err != nil {
				return nil, errors.InternalServer("INTERNAL_ERROR", "get session error")
			}

			// modified the value of the key, and save
			session.Values["key"] = "value"
			if err = session.Save(ht); err != nil {
				return nil, errors.InternalServer("INTERNAL_ERROR", "save session error")
			}

			// delete session
			//session.Options.MaxAge = -1
			//if err = session.Save(ht); err != nil {
			//	return nil, errors.InternalServer("INTERNAL_ERROR", "save session error")
			//}
		}
	}

	return &helloworld.HelloReply{Message: fmt.Sprintf("Hello %+v", in.Name)}, nil
}

func main() {
	rdCmd := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})
	store, err := sessions.NewRedisStore(rdCmd, []byte("secret"))
	store.SetMaxAge(10 * 24 * 3600)
	if err != nil {
		log.Fatal(err)
	}

	httpSrv := http.NewServer(
		http.Address(":8000"),
	)

	s := &server{store: store}
	helloworld.RegisterGreeterHTTPServer(httpSrv, s)

	app := kratos.New(
		kratos.Name(Name),
		kratos.Server(
			httpSrv,
		),
	)

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
