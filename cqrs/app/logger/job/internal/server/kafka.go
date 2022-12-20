package server

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/tx7do/kratos-transport/transport/kafka"

	"kratos-cqrs/app/logger/job/internal/conf"
	"kratos-cqrs/app/logger/job/internal/service"
)

// NewKafkaServer create a kafka server.
func NewKafkaServer(c *conf.Server, _ log.Logger, svc *service.LoggerJobService) *kafka.Server {
	ctx := context.Background()

	srv := kafka.NewServer(
		kafka.WithAddress(c.Kafka.Addrs),
		kafka.WithCodec("json"),
	)

	registerKafkaSubscribers(ctx, srv, svc)

	return srv
}

func registerKafkaSubscribers(ctx context.Context, srv *kafka.Server, svc *service.LoggerJobService) {
	_ = srv.RegisterSubscriber(ctx,
		"logger.sensor.ts",
		"sensor_logger",
		false,
		registerSensorDataHandler(svc.InsertSensorData),
		sensorDataCreator,
	)

	_ = srv.RegisterSubscriber(ctx,
		"logger.sensor.instance",
		"sensor",
		false,
		registerSensorHandler(svc.InsertSensor),
		sensorCreator,
	)
}
