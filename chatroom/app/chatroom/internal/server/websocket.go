package server

import (
	"errors"

	"github.com/go-kratos/kratos/v2/encoding"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/kratos-transport/transport/websocket"

	v1 "kratos-chatroom/api/chatroom/v1"
	"kratos-chatroom/app/chatroom/internal/conf"
	"kratos-chatroom/app/chatroom/internal/service"
)

// NewWebsocketServer create a websocket server.
func NewWebsocketServer(c *conf.Server, _ log.Logger, svc *service.ChatRoomService) *websocket.Server {
	srv := websocket.NewServer(
		websocket.WithAddress(c.Websocket.Addr),
		websocket.WithPath(c.Websocket.Path),
		websocket.WithConnectHandle(svc.OnWebsocketConnect),
		websocket.WithCodec(encoding.GetCodec("json")),
	)

	svc.SetWebsocketServer(srv)

	srv.RegisterMessageHandler(websocket.MessageType(v1.MessageType_Chat),
		func(sessionId websocket.SessionID, payload websocket.MessagePayload) error {
			switch t := payload.(type) {
			case *v1.ChatMessage:
				return svc.OnChatMessage(sessionId, t)
			default:
				return errors.New("invalid payload type")
			}
		},
		func() websocket.Any { return &v1.ChatMessage{} },
	)

	return srv
}
