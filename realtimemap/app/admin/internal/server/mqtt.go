package server

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/kratos-transport/transport/mqtt"
	"kratos-realtimemap/app/admin/internal/conf"
	"kratos-realtimemap/app/admin/internal/service"
)

// NewMQTTServer create a mqtt server.
func NewMQTTServer(c *conf.Server, _ log.Logger, svc *service.AdminService) *mqtt.Server {
	ctx := context.Background()

	srv := mqtt.NewServer(
		mqtt.WithAddress([]string{c.Mqtt.Addr}),
		mqtt.WithCodec("json"),
	)

	_ = srv.RegisterSubscriber(ctx,
		"/hfp/v2/journey/ongoing/vp/bus/#",
		registerSensorDataHandler(svc.TransitPostTelemetry),
		hfpEventCreator,
	)

	svc.SetMqttBroker(srv)

	return srv
}
