package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/kratos-transport/transport/kafka"

	"kratos-realtimemap/app/admin/internal/conf"
	"kratos-realtimemap/app/admin/internal/service"
)

// NewKafkaServer create a kafka server.
func NewKafkaServer(c *conf.Server, _ log.Logger, s *service.AdminService) *kafka.Server {
	//ctx := context.Background()

	srv := kafka.NewServer(
		kafka.WithAddress(c.Kafka.Addrs),
	)

	s.SetKafkaBroker(srv)

	return srv
}
