package kafka

import (
	"context"
	"fmt"
	"strconv"

	"github.com/IBM/sarama"
	"github.com/dnwe/otelsarama"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/propagation"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	"go.opentelemetry.io/otel/trace"
)

func WrapTrace(ctx context.Context, tracer trace.Tracer, textMapPropagator propagation.TextMapPropagator, message *sarama.ConsumerMessage, fn func(ctx context.Context, message *sarama.ConsumerMessage) error) error {
	// Extract a span context from message to link.
	carrier := otelsarama.NewConsumerMessageCarrier(message)
	parentSpanContext := textMapPropagator.Extract(ctx, carrier)

	// Create a span.
	attrs := []attribute.KeyValue{
		semconv.MessagingSystem("kafka"),
		semconv.MessagingDestinationKindTopic,
		semconv.MessagingDestinationName(message.Topic),
		semconv.MessagingOperationReceive,
		semconv.MessagingMessageID(strconv.FormatInt(message.Offset, 10)),
		semconv.MessagingKafkaSourcePartition(int(message.Partition)),
	}
	opts := []trace.SpanStartOption{
		trace.WithAttributes(attrs...),
		trace.WithSpanKind(trace.SpanKindConsumer),
	}
	newCtx, span := tracer.Start(parentSpanContext, fmt.Sprintf("%s handle", message.Topic), opts...)
	defer span.End()
	return fn(newCtx, message)
}
