package server

import (
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc/encoding"
	svcV1 "kratos-cqrs/api/logger/service/v1"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/kratos-transport/broker"
	"github.com/tx7do/kratos-transport/transport/kafka"

	"kratos-cqrs/app/logger/job/internal/conf"
	"kratos-cqrs/app/logger/job/internal/service"
)

// NewKafkaServer create a kafka server.
func NewKafkaServer(c *conf.Server, _ log.Logger, svc *service.LoggerJobService) *kafka.Server {
	ctx := context.Background()

	srv := kafka.NewServer(
		kafka.WithAddress(c.Kafka.Addrs),
		kafka.WithCodec(encoding.GetCodec("json")),
	)

	registerKafkaSubscribers(ctx, srv, svc)

	return srv
}

func registerKafkaSubscribers(ctx context.Context, srv *kafka.Server, svc *service.LoggerJobService) {
	registerInsertSensorDataHandler(ctx, srv, svc)
	registerInsertSensorHandler(ctx, srv, svc)
}

func registerInsertSensorDataHandler(ctx context.Context, srv *kafka.Server, svc *service.LoggerJobService) {
	binder := func() broker.Any { return &svcV1.SensorData{} }

	fnc := func(ctx context.Context, event broker.Event) error {
		var msg *[]*svcV1.SensorData

		switch t := event.Message().Body.(type) {
		case []byte:
			msg = &[]*svcV1.SensorData{}
			if err := json.Unmarshal(t, msg); err != nil {
				return err
			}
		case string:
			msg = &[]*svcV1.SensorData{}
			if err := json.Unmarshal([]byte(t), msg); err != nil {
				return err
			}
		case *[]*svcV1.SensorData:
			msg = t
		default:
			return fmt.Errorf("unsupported type: %T", t)
		}

		if err := svc.InsertSensorData(ctx, event.Topic(), event.Message().Headers, msg); err != nil {
			return err
		}

		return nil
	}

	_ = srv.RegisterSubscriber(ctx,
		"logger.sensor.ts",
		"sensor_logger",
		false,
		fnc,
		binder,
	)
}

func registerInsertSensorHandler(ctx context.Context, srv *kafka.Server, svc *service.LoggerJobService) {
	binder := func() broker.Any { return &svcV1.Sensor{} }

	fnc := func(ctx context.Context, event broker.Event) error {
		var msg *svcV1.Sensor

		switch t := event.Message().Body.(type) {
		case []byte:
			msg = &svcV1.Sensor{}
			if err := json.Unmarshal(t, msg); err != nil {
				return err
			}
		case string:
			msg = &svcV1.Sensor{}
			if err := json.Unmarshal([]byte(t), msg); err != nil {
				return err
			}
		case *svcV1.Sensor:
			msg = t
		default:
			return fmt.Errorf("unsupported type: %T", t)
		}

		if err := svc.InsertSensor(ctx, event.Topic(), event.Message().Headers, msg); err != nil {
			return err
		}

		return nil
	}

	_ = srv.RegisterSubscriber(ctx,
		"logger.sensor.instance",
		"sensor",
		false,
		fnc,
		binder,
	)
}
