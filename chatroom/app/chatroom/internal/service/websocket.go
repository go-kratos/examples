package service

import (
	"encoding/json"
	"fmt"
	"github.com/tx7do/kratos-transport/transport/websocket"
	v1 "kratos-chatroom/api/chatroom/v1"
)

func (s *ChatRoomService) SetWebsocketServer(ws *websocket.Server) {
	s.ws = ws
}

func (s *ChatRoomService) OnWebsocketMessage(connectionId string, message *websocket.Message) error {
	s.log.Infof("[%s] Payload: %s\n", connectionId, string(message.Body))

	var proto v1.WebsocketProto

	if err := json.Unmarshal(message.Body, &proto); err != nil {
		s.log.Error("Error unmarshalling proto json %v", err)
		return nil
	}

	switch proto.EventId {
	case "chat":
		chatMsg := proto.Payload
		fmt.Println("chat message:", chatMsg)
		_ = s.OnChatMessage(connectionId, &chatMsg)
	}

	return nil
}

func (s *ChatRoomService) OnChatMessage(connectionId string, msg *string) error {
	s.BroadcastToWebsocketClient("chat", msg)
	return nil
}

func (s *ChatRoomService) OnWebsocketConnect(connectionId string, register bool) {
	if register {
		fmt.Printf("%s connected\n", connectionId)
	} else {
		fmt.Printf("%s disconnect\n", connectionId)
	}
}

func (s *ChatRoomService) BroadcastToWebsocketClient(eventId string, payload interface{}) {
	if payload == nil {
		return
	}

	bufPayload, _ := json.Marshal(&payload)

	var proto v1.WebsocketProto
	proto.EventId = eventId
	proto.Payload = string(bufPayload)

	bufProto, _ := json.Marshal(&proto)

	var msg websocket.Message
	msg.Body = bufProto

	s.ws.Broadcast(&msg)
}
