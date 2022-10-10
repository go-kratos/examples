package server

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-kratos/kratos/v2/encoding"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/kratos-transport/broker"
	"github.com/tx7do/kratos-transport/transport/mqtt"
	"kratos-realtimemap/api/hfp"

	"kratos-realtimemap/app/admin/internal/conf"
	"kratos-realtimemap/app/admin/internal/service"
)

// NewMQTTServer create a mqtt server.
func NewMQTTServer(c *conf.Server, _ log.Logger, svc *service.AdminService) *mqtt.Server {
	ctx := context.Background()

	srv := mqtt.NewServer(
		mqtt.WithAddress([]string{c.Mqtt.Addr}),
		mqtt.WithCodec(encoding.GetCodec("json")),
	)

	registerTransitPostTelemetryHandler(ctx, srv, svc)

	svc.SetMqttBroker(srv)

	return srv
}

func registerTransitPostTelemetryHandler(ctx context.Context, srv *mqtt.Server, svc *service.AdminService) {
	fnc := func(ctx context.Context, event broker.Event) error {
		var msg *hfp.Event = nil

		switch t := event.Message().Body.(type) {
		case []byte:
			msg = &hfp.Event{}
			if err := json.Unmarshal(t, msg); err != nil {
				return err
			}
		case string:
			msg = &hfp.Event{}
			if err := json.Unmarshal([]byte(t), msg); err != nil {
				return err
			}
		case *hfp.Event:
			msg = t
		default:
			return fmt.Errorf("unsupported type: %T", t)
		}

		if err := svc.TransitPostTelemetry(ctx, event.Topic(), event.Message().Headers, msg); err != nil {
			return err
		}

		return nil
	}

	_ = srv.RegisterSubscriber(ctx,
		"/hfp/v2/journey/ongoing/vp/bus/#",
		fnc,
		func() broker.Any {
			return &hfp.Event{}
		})
}
