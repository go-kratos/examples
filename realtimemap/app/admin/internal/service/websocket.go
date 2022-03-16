package service

import (
	"encoding/json"
	"fmt"
	"github.com/tx7do/kratos-transport/transport/websocket"
	v1 "kratos-realtimemap/api/admin/v1"
	"kratos-realtimemap/app/admin/internal/pkg/data"
)

func (s *AdminService) SetWebsocketServer(ws *websocket.Server) {
	s.ws = ws
}

type WebsocketProto struct {
	EventId string      `protobuf:"bytes,1,opt,name=event_id,json=eventId,proto3" json:"eventId,omitempty"`
	Payload interface{} `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (s *AdminService) OnWebsocketMessage(connectionId string, message *websocket.Message) (*websocket.Message, error) {
	s.log.Infof("[%s] Payload: %s\n", connectionId, string(message.Body))

	var proto v1.WebsocketProto

	if err := json.Unmarshal(message.Body, &proto); err != nil {
		s.log.Error("Error unmarshalling proto json %v", err)
		return nil, nil
	}

	switch proto.EventId {
	case "viewport":
		var msg v1.Viewport
		if err := json.Unmarshal([]byte(proto.Payload), &msg); err != nil {
			s.log.Error("Error unmarshalling payload json %v", err)
			return nil, nil
		}

		_ = s.OnWsSetViewport(connectionId, &msg)
	}

	return nil, nil
}

func (s *AdminService) OnWsSetViewport(connectionId string, msg *v1.Viewport) error {
	s.viewports[connectionId] = msg
	return nil
}

func (s *AdminService) OnWebsocketConnect(connectionId string, register bool) {
	if register {
	} else {
		delete(s.viewports, connectionId)
	}
}

func (s *AdminService) BroadcastToWebsocketClient(eventId string, payload interface{}) {
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

func (s *AdminService) BroadcastVehiclePosition(positions data.PositionArray) {
	s.BroadcastToWebsocketClient("positions", positions)
}

func (s *AdminService) BroadcastVehicleTurnoverNotification(turnovers data.TurnoverArray) {
	for _, turnover := range turnovers {
		var str string
		if turnover.Status {
			str = fmt.Sprintf("%s from %s entered the zone %s",
				turnover.VehicleId, turnover.OrganizationName, turnover.GeofenceName)
		} else {
			str = fmt.Sprintf("%s from %s left the zone %s",
				turnover.VehicleId, turnover.OrganizationName, turnover.GeofenceName)
		}
		s.BroadcastToWebsocketClient("notification", str)
	}
}
