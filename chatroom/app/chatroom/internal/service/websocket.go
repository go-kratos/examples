package service

import (
	"github.com/tx7do/kratos-transport/transport/websocket"
	v1 "kratos-chatroom/api/chatroom/v1"
)

func (s *ChatRoomService) SetWebsocketServer(ws *websocket.Server) {
	s.ws = ws
}

func (s *ChatRoomService) OnWebsocketConnect(sessionId websocket.SessionID, register bool) {
	if register {
		s.log.Infof("%s connected\n", sessionId)
	} else {
		s.log.Infof("%s disconnect\n", sessionId)
	}
}

func (s *ChatRoomService) OnChatMessage(sessionId websocket.SessionID, msg *v1.ChatMessage) error {
	s.ws.Broadcast(websocket.MessageType(v1.MessageType_Chat), msg)
	return nil
}
