package dep

import (
	"context"

	"otel/internal/conf"

	"github.com/go-kratos/kratos/v2/middleware/metrics"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/propagation"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.24.0"
	"go.opentelemetry.io/otel/trace"

	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
)

func NewMeter(bc *conf.Bootstrap, provider metric.MeterProvider) (metric.Meter, error) {
	meta := bc.Metadata

	return provider.Meter(meta.Name), nil
}

func NewMeterProvider(bc *conf.Bootstrap) (metric.MeterProvider, error) {
	meta := bc.Metadata
	metricConf := bc.Otel.Metric
	exporter, err := prometheus.New()
	if err != nil {
		return nil, err
	}

	if metricConf.EnableExemplar {
		err = metrics.EnableOTELExemplar()
		if err != nil {
			return nil, err
		}
	}

	provider := sdkmetric.NewMeterProvider(
		sdkmetric.WithResource(
			resource.NewWithAttributes(
				semconv.SchemaURL,
				semconv.ServiceNameKey.String(meta.Name),
				attribute.String("environment", meta.Env.String()),
			),
		),
		sdkmetric.WithReader(exporter),
		sdkmetric.WithView(
			metrics.DefaultSecondsHistogramView(metrics.DefaultServerSecondsHistogramName),
		),
	)
	otel.SetMeterProvider(provider)
	return provider, nil
}

func NewTracerProvider(ctx context.Context, bc *conf.Bootstrap, textMapPropagator propagation.TextMapPropagator) (trace.TracerProvider, error) {
	meta := bc.Metadata
	traceConf := bc.Otel.Trace

	opts := []otlptracehttp.Option{otlptracehttp.WithEndpoint(traceConf.Endpoint)}
	if traceConf.Insecure {
		opts = append(opts, otlptracehttp.WithInsecure())
	}
	client := otlptracehttp.NewClient(opts...)
	exp, err := otlptrace.New(ctx, client)
	if err != nil {
		return nil, err
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exp),
		sdktrace.WithResource(
			resource.NewWithAttributes(
				semconv.SchemaURL,
				semconv.ServiceNameKey.String(meta.Name),
				attribute.String("environment", meta.Env.String()),
			),
		),
	)

	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(textMapPropagator)

	return tp, nil
}

func NewTracer(bc *conf.Bootstrap, tp trace.TracerProvider) (trace.Tracer, error) {
	meta := bc.Metadata
	return tp.Tracer(meta.Name), nil
}

func NewTextMapPropagator() propagation.TextMapPropagator {
	return propagation.NewCompositeTextMapPropagator(
		tracing.Metadata{},
		propagation.Baggage{},
		propagation.TraceContext{},
	)
}
