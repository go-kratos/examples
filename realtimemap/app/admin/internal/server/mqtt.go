package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/kratos-transport/transport/mqtt"

	"kratos-realtimemap/app/admin/internal/conf"
	"kratos-realtimemap/app/admin/internal/service"
)

// NewMQTTServer create a mqtt server.
func NewMQTTServer(c *conf.Server, _ log.Logger, s *service.AdminService) *mqtt.Server {
	//ctx := context.Background()

	srv := mqtt.NewServer(
		mqtt.Address(c.Mqtt.Addr),
		mqtt.Subscribe("/hfp/v2/journey/ongoing/vp/bus/#", s.TransitPostTelemetry),
	)

	s.SetMqttBroker(srv)

	return srv
}
