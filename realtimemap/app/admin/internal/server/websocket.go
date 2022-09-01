package server

import (
	"github.com/tx7do/kratos-transport/transport/websocket"

	"github.com/go-kratos/kratos/v2/log"
	"kratos-realtimemap/app/admin/internal/conf"
	"kratos-realtimemap/app/admin/internal/service"
)

// NewWebsocketServer create a websocket server.
func NewWebsocketServer(c *conf.Server, _ log.Logger, s *service.AdminService) *websocket.Server {
	srv := websocket.NewServer(
		websocket.WithAddress(c.Websocket.Addr),
		websocket.WithReadHandle(c.Websocket.Path, s.OnWebsocketMessage),
		websocket.WithConnectHandle(s.OnWebsocketConnect),
	)

	s.SetWebsocketServer(srv)

	return srv
}
