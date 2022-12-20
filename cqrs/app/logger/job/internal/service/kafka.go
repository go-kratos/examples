package service

import (
	"context"
	"fmt"
	"github.com/tx7do/kratos-transport/broker"
	svcV1 "kratos-cqrs/api/logger/service/v1"
)

func (s *LoggerJobService) InsertSensorData(ctx context.Context, topic string, headers broker.Headers, msg *[]svcV1.SensorData) error {
	fmt.Println("InsertSensorData() Topic: ", topic)

	if err := s.sensorData.BatchInsertSensorData(context.Background(), msg); err != nil {
		s.log.Debug("InsertSensorData Insert", err.Error())
		return err
	}

	return nil
}

func (s *LoggerJobService) InsertSensor(ctx context.Context, topic string, headers broker.Headers, msg *svcV1.Sensor) error {
	fmt.Println("InsertSensor() Topic: ", topic)

	if err := s.sensor.Create(context.Background(), msg); err != nil {
		s.log.Debug("InsertSensor Insert", err.Error())
		return err
	}

	return nil
}
