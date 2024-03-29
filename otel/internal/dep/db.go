package dep

import (
	"database/sql"

	"otel/internal/conf"

	"github.com/XSAM/otelsql"
	"github.com/redis/go-redis/extra/redisotel/v9"
	"github.com/redis/go-redis/v9"
	"go.opentelemetry.io/otel/metric"
	semconv "go.opentelemetry.io/otel/semconv/v1.24.0"
	"go.opentelemetry.io/otel/trace"
)

func NewDB(c *conf.Data, provider trace.TracerProvider, meterProvider metric.MeterProvider) (*sql.DB, func(), error) {
	driverName, err := otelsql.Register(
		c.Database.Driver,
		otelsql.WithTracerProvider(provider),
		otelsql.WithAttributes(
			semconv.DBSystemMySQL,
		),
		otelsql.WithMeterProvider(meterProvider),
	)
	if err != nil {
		return nil, nil, err
	}

	db, err := sql.Open(driverName, c.Database.Source)
	if err != nil {
		return nil, nil, err
	}

	return db, func() {
		db.Close()
	}, nil
}

func NewRedis(c *conf.Data, provider trace.TracerProvider, meterProvider metric.MeterProvider) (*redis.Client, error) {
	redisClient := redis.NewClient(&redis.Options{
		Network:      c.Redis.Network,
		Addr:         c.Redis.Addr,
		DialTimeout:  c.Redis.ReadTimeout.AsDuration(),
		ReadTimeout:  c.Redis.ReadTimeout.AsDuration(),
		WriteTimeout: c.Redis.WriteTimeout.AsDuration(),
	})

	// Enable tracing instrumentation.
	if err := redisotel.InstrumentTracing(redisClient,
		redisotel.WithTracerProvider(provider),
		redisotel.WithAttributes(semconv.DBSystemRedis),
	); err != nil {
		return nil, err
	}

	// Enable metrics instrumentation.
	if err := redisotel.InstrumentMetrics(redisClient,
		redisotel.WithMeterProvider(meterProvider),
		redisotel.WithAttributes(semconv.DBSystemRedis),
	); err != nil {
		return nil, err
	}
	return redisClient, nil
}
