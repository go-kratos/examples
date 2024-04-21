package dep

import (
	"fmt"
	"otel/internal/conf"

	"github.com/IBM/sarama"
	"github.com/dnwe/otelsarama"
	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

func NewProducer(bc *conf.Bootstrap, logger log.Logger, tracerProvider trace.TracerProvider, textMapPropagator propagation.TextMapPropagator) (sarama.AsyncProducer, func(), error) {
	kafkaConf := bc.Data.Kafka
	config := sarama.NewConfig()
	config.Version = sarama.V2_5_0_0
	// So we can know the partition and offset of messages.
	config.Producer.Return.Successes = true

	producer, err := sarama.NewAsyncProducer(kafkaConf.BrokerList, config)
	if err != nil {
		return nil, nil, fmt.Errorf("starting Sarama producer: %w", err)
	}

	// Wrap instrumentation
	producer = otelsarama.WrapAsyncProducer(config, producer,
		otelsarama.WithTracerProvider(tracerProvider),
		otelsarama.WithPropagators(textMapPropagator),
	)

	lh := log.NewHelper(logger)
	// We will log to STDOUT if we're not able to produce messages.
	go func() {
		for err := range producer.Errors() {
			lh.Error("Failed to write message:", err)
		}
	}()
	go func() {
		for range producer.Successes() {
		}
	}()

	return producer, func() {
		producer.Close()
	}, nil
}

func NewConsumerGroup(bc *conf.Bootstrap) (sarama.ConsumerGroup, func(), error) {
	kafkaConf := bc.Data.Kafka

	config := sarama.NewConfig()
	config.Version = sarama.V2_5_0_0
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	consumerGroup, err := sarama.NewConsumerGroup(kafkaConf.BrokerList, kafkaConf.GroupId, config)
	if err != nil {
		return nil, nil, fmt.Errorf("starting consumer group: %w", err)
	}

	return consumerGroup, func() {
		consumerGroup.Close()
	}, nil
}
