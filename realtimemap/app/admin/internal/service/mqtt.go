package service

import (
	"context"
	"github.com/tx7do/kratos-transport/broker"
	"kratos-realtimemap/api/hfp"
	"kratos-realtimemap/app/admin/internal/pkg/data"
)

const MaxPositionHistory = 200

func (s *AdminService) SetMqttBroker(b broker.Broker) {
	s.mb = b
}

func (s *AdminService) TransitPostTelemetry(_ context.Context, topic string, headers broker.Headers, msg *hfp.Event) error {
	//fmt.Println("Topic: ", topic)

	topicInfo := hfp.Topic{}
	topicInfo.Parse(topic)

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

	return nil
}
