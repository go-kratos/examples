package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/kratos-transport/transport/websocket"

	"kratos-chatroom/app/chatroom/internal/conf"
	"kratos-chatroom/app/chatroom/internal/service"
)

// NewWebsocketServer create a websocket server.
func NewWebsocketServer(c *conf.Server, _ log.Logger, svc *service.ChatRoomService) *websocket.Server {
	srv := websocket.NewServer(
		websocket.WithAddress(c.Websocket.Addr),
		websocket.WithReadHandle(c.Websocket.Path, svc.OnWebsocketMessage),
		websocket.WithConnectHandle(svc.OnWebsocketConnect),
	)

	svc.SetWebsocketServer(srv)

	return srv
}
