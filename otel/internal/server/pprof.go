package server

import (
	"context"
	"net"
	"net/http"
	_ "net/http/pprof"
	"sync/atomic"

	"otel/internal/conf"
)

type PprofServer struct {
	started uint32

	conf *conf.Server_Pprof

	listener net.Listener
}

func NewPprof(server *conf.Server) (*PprofServer, error) {
	pprof := &PprofServer{
		conf: server.Pprof,
	}
	addr := server.GetPprof().GetAddr()
	if addr != "" {
		listener, err := net.Listen("tcp", addr)
		if err != nil {
			return nil, err
		}
		pprof.listener = listener
	}
	return pprof, nil
}

func (x *PprofServer) Start(ctx context.Context) error {
	if x.listener == nil {
		return nil
	}
	if !atomic.CompareAndSwapUint32(&x.started, 0, 1) {
		return nil
	}
	defer atomic.CompareAndSwapUint32(&x.started, 1, 0)
	return http.Serve(x.listener, nil)
}

func (x *PprofServer) Stop(ctx context.Context) error {
	if atomic.LoadUint32(&x.started) == 1 {
		return x.listener.Close()
	}
	return nil
}
