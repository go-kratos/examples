package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/tx7do/kratos-transport/transport/websocket"
	v1 "kratos-chatroom/api/chatroom/v1"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewChatRoomService)

type ChatRoomService struct {
	v1.UnimplementedChatRoomServer

	log *log.Helper
	ws  *websocket.Server
}

func NewChatRoomService(logger log.Logger) *ChatRoomService {
	l := log.NewHelper(log.With(logger, "module", "service/chatroom"))
	return &ChatRoomService{
		log: l,
	}
}
