package bootstrap

import (
	"os"

	"github.com/sirupsen/logrus"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"

	aliyunLogger "github.com/go-kratos/kratos/contrib/log/aliyun/v2"
	fluentLogger "github.com/go-kratos/kratos/contrib/log/fluent/v2"
	logrusLogger "github.com/go-kratos/kratos/contrib/log/logrus/v2"
	tencentLogger "github.com/go-kratos/kratos/contrib/log/tencent/v2"
	zapLogger "github.com/go-kratos/kratos/contrib/log/zap/v2"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"

	"kratos-gorm-example/gen/api/go/common/conf"
)

type LoggerType string

const (
	LoggerTypeStd     LoggerType = "std"
	LoggerTypeFluent  LoggerType = "fluent"
	LoggerTypeLogrus  LoggerType = "logrus"
	LoggerTypeZap     LoggerType = "zap"
	LoggerTypeAliyun  LoggerType = "aliyun"
	LoggerTypeTencent LoggerType = "tencent"
)

// NewLoggerProvider 创建一个新的日志记录器提供者
func NewLoggerProvider(cfg *conf.Logger, serviceInfo *ServiceInfo) log.Logger {
	l := NewLogger(cfg)

	return log.With(
		l,
		"service.id", serviceInfo.Id,
		"service.name", serviceInfo.Name,
		"service.version", serviceInfo.Version,
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"trace_id", tracing.TraceID(),
		"span_id", tracing.SpanID(),
	)
}

// NewLogger 创建一个新的日志记录器
func NewLogger(cfg *conf.Logger) log.Logger {
	if cfg == nil {
		return NewStdLogger()
	}

	switch LoggerType(cfg.Type) {
	default:
		fallthrough
	case LoggerTypeStd:
		return NewStdLogger()
	case LoggerTypeFluent:
		return NewFluentLogger(cfg)
	case LoggerTypeZap:
		return NewZapLogger(cfg)
	case LoggerTypeLogrus:
		return NewLogrusLogger(cfg)
	case LoggerTypeAliyun:
		return NewAliyunLogger(cfg)
	case LoggerTypeTencent:
		return NewTencentLogger(cfg)
	}
}

// NewStdLogger 创建一个新的日志记录器 - Kratos内置，控制台输出
func NewStdLogger() log.Logger {
	l := log.NewStdLogger(os.Stdout)
	return l
}

// NewZapLogger 创建一个新的日志记录器 - Zap
func NewZapLogger(cfg *conf.Logger) log.Logger {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	jsonEncoder := zapcore.NewJSONEncoder(encoderConfig)

	lumberJackLogger := &lumberjack.Logger{
		Filename:   cfg.Zap.Filename,
		MaxSize:    int(cfg.Zap.MaxSize),
		MaxBackups: int(cfg.Zap.MaxBackups),
		MaxAge:     int(cfg.Zap.MaxAge),
	}
	writeSyncer := zapcore.AddSync(lumberJackLogger)

	var lvl = new(zapcore.Level)
	if err := lvl.UnmarshalText([]byte(cfg.Zap.Level)); err != nil {
		return nil
	}

	core := zapcore.NewCore(jsonEncoder, writeSyncer, lvl)
	logger := zap.New(core).WithOptions()

	wrapped := zapLogger.NewLogger(logger)

	return wrapped
}

// NewLogrusLogger 创建一个新的日志记录器 - Logrus
func NewLogrusLogger(cfg *conf.Logger) log.Logger {
	loggerLevel, err := logrus.ParseLevel(cfg.Logrus.Level)
	if err != nil {
		loggerLevel = logrus.InfoLevel
	}

	var loggerFormatter logrus.Formatter
	switch cfg.Logrus.Formatter {
	default:
		fallthrough
	case "text":
		loggerFormatter = &logrus.TextFormatter{
			DisableColors:    cfg.Logrus.DisableColors,
			DisableTimestamp: cfg.Logrus.DisableTimestamp,
			TimestampFormat:  cfg.Logrus.TimestampFormat,
		}
		break
	case "json":
		loggerFormatter = &logrus.JSONFormatter{
			DisableTimestamp: cfg.Logrus.DisableTimestamp,
			TimestampFormat:  cfg.Logrus.TimestampFormat,
		}
		break
	}

	logger := logrus.New()
	logger.Level = loggerLevel
	logger.Formatter = loggerFormatter

	wrapped := logrusLogger.NewLogger(logger)
	return wrapped
}

// NewFluentLogger 创建一个新的日志记录器 - Fluent
func NewFluentLogger(cfg *conf.Logger) log.Logger {
	wrapped, err := fluentLogger.NewLogger(cfg.Fluent.Endpoint)
	if err != nil {
		panic("create fluent logger failed")
		return nil
	}
	return wrapped
}

// NewAliyunLogger 创建一个新的日志记录器 - Aliyun
func NewAliyunLogger(cfg *conf.Logger) log.Logger {
	wrapped := aliyunLogger.NewAliyunLog(
		aliyunLogger.WithProject(cfg.Aliyun.Project),
		aliyunLogger.WithEndpoint(cfg.Aliyun.Endpoint),
		aliyunLogger.WithAccessKey(cfg.Aliyun.AccessKey),
		aliyunLogger.WithAccessSecret(cfg.Aliyun.AccessSecret),
	)
	return wrapped
}

// NewTencentLogger 创建一个新的日志记录器 - Tencent
func NewTencentLogger(cfg *conf.Logger) log.Logger {
	wrapped, err := tencentLogger.NewLogger(
		tencentLogger.WithTopicID(cfg.Tencent.TopicId),
		tencentLogger.WithEndpoint(cfg.Tencent.Endpoint),
		tencentLogger.WithAccessKey(cfg.Tencent.AccessKey),
		tencentLogger.WithAccessSecret(cfg.Tencent.AccessSecret),
	)
	if err != nil {
		panic(err)
		return nil
	}
	return wrapped
}
