package data

import (
	"context"
	"strconv"
	"time"

	"otel/internal/pkg/trace/kafka"

	"otel/internal/biz"

	"github.com/IBM/sarama"
	"github.com/dnwe/otelsarama"
	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

type greeterRepo struct {
	data              *Data
	log               *log.Helper
	producer          sarama.AsyncProducer
	consumerGroup     sarama.ConsumerGroup
	tracerProvider    trace.TracerProvider
	textMapPropagator propagation.TextMapPropagator
	tracer            trace.Tracer
}

func (r *greeterRepo) Setup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (r *greeterRepo) Cleanup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (r *greeterRepo) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	r.log.Infow("claim.Partition()", claim.Partition())
	r.log.Infow("session.Claims()", session.Claims())
	for {
		var (
			message *sarama.ConsumerMessage
			ok      bool
		)
		select {
		case message, ok = <-claim.Messages():
			if !ok {
				r.log.Info("chan closed")
				return nil
			}
		case <-session.Context().Done():
			r.log.Info("session.Context().Done()")
			return nil
		}
		session.MarkMessage(message, "")
		_ = kafka.WrapTrace(session.Context(), r.tracer, r.textMapPropagator, message, func(ctx context.Context, message *sarama.ConsumerMessage) error {
			i64, _ := strconv.Atoi(string(message.Key))
			r.log.WithContext(ctx).Infow("handleKeyExpired:key", string(message.Key), "message", message)
			if i64 == 0 {
				return nil
			}
			greeter, err := r.findByID(ctx, int64(i64))
			if err != nil {
				r.log.WithContext(ctx).Error("r.findByID err ", err)
				return err
			}
			err = r.data.redis.Set(ctx, string(message.Key), greeter.Hello, time.Hour).Err()
			if err != nil {
				r.log.WithContext(ctx).Error("redis.Set err ", err)
			}
			return nil
		})
	}
}

// NewGreeterRepo .
func NewGreeterRepo(ctx context.Context, data *Data, logger log.Logger,
	producer sarama.AsyncProducer, consumerGroup sarama.ConsumerGroup,
	tracerProvider trace.TracerProvider, textMapPropagator propagation.TextMapPropagator,
) biz.GreeterRepo {
	lh := log.NewHelper(logger)
	repo := &greeterRepo{
		data:              data,
		log:               lh,
		producer:          producer,
		consumerGroup:     consumerGroup,
		tracerProvider:    tracerProvider,
		tracer:            tracerProvider.Tracer("greeter_repo"),
		textMapPropagator: textMapPropagator,
	}

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-time.After(time.Second):
			}
			err := repo.handleKeyExpired(ctx)
			if err != nil {
				lh.Error("handleKeyExpired error", err)
			}
		}
	}()

	return repo
}

func (r *greeterRepo) produceKeyExpired(ctx context.Context, key string) {
	msg := &sarama.ProducerMessage{
		Topic: "test",
		Key:   sarama.StringEncoder(key),
	}
	r.textMapPropagator.Inject(ctx, otelsarama.NewProducerMessageCarrier(msg))
	r.producer.Input() <- msg
}

func (r *greeterRepo) handleKeyExpired(ctx context.Context) error {
	return r.consumerGroup.Consume(ctx, []string{"test"}, otelsarama.WrapConsumerGroupHandler(r, otelsarama.WithTracerProvider(r.tracerProvider)))
}

func (r *greeterRepo) Save(ctx context.Context, g *biz.Greeter) (*biz.Greeter, error) {
	stmt, err := r.data.db.PrepareContext(ctx, "insert into greeter(hello) values (?)")
	if err != nil {
		return nil, err
	}
	result, err := stmt.Exec(g.Hello)
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	r.produceKeyExpired(ctx, strconv.Itoa(int(id)))
	return g, nil
}

func (r *greeterRepo) FindByID(ctx context.Context, id int64) (*biz.Greeter, error) {
	hello, _ := r.data.redis.Get(ctx, strconv.Itoa(int(id))).Result()
	if hello != "" {
		return &biz.Greeter{
			Hello: hello,
		}, nil
	}
	return r.findByID(ctx, id)
}

func (r *greeterRepo) findByID(ctx context.Context, id int64) (*biz.Greeter, error) {
	stmt, err := r.data.db.PrepareContext(ctx, "select hello from greeter where id = ?")
	if err != nil {
		return nil, err
	}
	row := stmt.QueryRow(id)
	var dst biz.Greeter
	err = row.Scan(&dst.Hello)
	if err != nil {
		return nil, err
	}
	return &dst, nil
}

func (r *greeterRepo) ListByHello(ctx context.Context, hello string) ([]*biz.Greeter, error) {
	stmt, err := r.data.db.PrepareContext(ctx, "select hello from greeter where hello = ?")
	if err != nil {
		return nil, err
	}
	rows, err := stmt.QueryContext(ctx, "%"+hello+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*biz.Greeter
	for rows.Next() {
		var dst biz.Greeter
		err = rows.Scan(&dst.Hello)
		if err != nil {
			return nil, err
		}
		result = append(result, &dst)
	}
	return result, nil
}

func (r *greeterRepo) ListAll(ctx context.Context) ([]*biz.Greeter, error) {
	stmt, err := r.data.db.PrepareContext(ctx, "select hello from greeter")
	if err != nil {
		return nil, err
	}
	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var result []*biz.Greeter
	for rows.Next() {
		var dst biz.Greeter
		err = rows.Scan(&dst.Hello)
		if err != nil {
			return nil, err
		}
		result = append(result, &dst)
	}
	return result, nil
}
