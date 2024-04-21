package dep

import "github.com/google/wire"

var ProviderSet = wire.NewSet(NewDB, NewRedis, NewMeter, NewTracer, NewTracerProvider,
	NewProducer, NewConsumerGroup, NewTextMapPropagator, NewMeterProvider)
