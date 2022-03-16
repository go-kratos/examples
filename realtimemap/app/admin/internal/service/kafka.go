package service

import (
	"github.com/tx7do/kratos-transport/broker"
	"google.golang.org/protobuf/proto"
)

func (s *AdminService) SetKafkaBroker(b broker.Broker) {
	s.kb = b
}

func (s *AdminService) sendToQueue(topic string, payload proto.Message) error {
	sendData, err := proto.Marshal(payload)
	if err != nil {
		s.log.Fatal("queue message marshaling error: ", err)
		return nil
	}

	var msg broker.Message
	msg.Body = sendData
	return s.kb.Publish(topic, &msg)
}
