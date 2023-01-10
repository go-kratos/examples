package service

import (
	"github.com/tx7do/kratos-transport/broker"
	"google.golang.org/protobuf/proto"
)

func (s *AdminService) SetKafkaBroker(b broker.Broker) {
	s.kb = b
}

func (s *AdminService) sendToQueue(topic string, payload proto.Message) error {
	return s.kb.Publish(topic, payload)
}
