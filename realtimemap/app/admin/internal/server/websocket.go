package server

import (
	"github.com/tx7do/kratos-transport/transport/websocket"

	"github.com/go-kratos/kratos/v2/log"
	"kratos-realtimemap/app/admin/internal/conf"
	"kratos-realtimemap/app/admin/internal/service"
)

// NewWebsocketServer create a websocket server.
func NewWebsocketServer(c *conf.Server, _ log.Logger, svc *service.AdminService) *websocket.Server {
	srv := websocket.NewServer(
		websocket.WithAddress(c.Websocket.Addr),
		websocket.WithPath(c.Websocket.Path),
		websocket.WithConnectHandle(svc.OnWebsocketConnect),
		websocket.WithCodec("json"),
	)

	svc.SetWebsocketServer(srv)

	return srv
}
