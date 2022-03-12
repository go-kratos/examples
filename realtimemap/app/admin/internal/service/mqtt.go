package service

import (
	"encoding/json"
	"github.com/tx7do/kratos-transport/broker"
	"kratos-realtimemap/api/hfp"
	"kratos-realtimemap/app/admin/internal/pkg/data"
)

const MaxPositionHistory = 200

func (s *AdminService) SetMqttBroker(b broker.Broker) {
	s.mb = b
}

func (s *AdminService) TransitPostTelemetry(event broker.Event) error {
	//fmt.Println("Topic: ", event.Topic(), " Payload: ", string(event.Message().Body))

	topicInfo := hfp.Topic{}
	topicInfo.Parse(event.Topic())

	var msg hfp.Event

	if err := json.Unmarshal(event.Message().Body, &msg); err != nil {
		s.log.Error("Error unmarshalling json %v", err)
	} else {

		msg.OperatorId = topicInfo.OperatorId
		msg.VehicleId = topicInfo.GetVehicleUID()

		position := msg.MapToPosition()
		if position != nil {
			s.positionHistory.Update(position)
			turnovers := data.AllOrganizations.Update(position)

			s.BroadcastVehicleTurnoverNotification(turnovers)
			s.BroadcastVehiclePosition(s.positionHistory.GetPositionsHistory(position.VehicleId))
		}

		s.log.Infof("事件类型: %s 交通工具类型: %s 司机ID: %s 车辆ID: %s", topicInfo.EventType, topicInfo.TransportMode, topicInfo.OperatorId, msg.VehicleId)
	}

	return nil
}
