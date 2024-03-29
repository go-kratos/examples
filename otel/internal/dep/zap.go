package dep

import (
	"fmt"

	"otel/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"go.uber.org/zap"
)

type LogWrapper func(level log.Level, keyvals ...interface{}) error

func (f LogWrapper) Log(level log.Level, keyvals ...interface{}) error {
	return f(level, keyvals...)
}

func NewZapLogger(bc *conf.Bootstrap) (log.Logger, error) {
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{"stderr"}
	logfile := bc.GetLog().GetFilepath()
	if logfile != "" {
		cfg.OutputPaths = append(cfg.OutputPaths, logfile)
	}
	logger, err := cfg.Build()
	if err != nil {
		return nil, err
	}
	return LogWrapper(func(level log.Level, keyvals ...interface{}) error {
		zapLevel := zap.DebugLevel
		switch level {
		case log.LevelDebug:
			zapLevel = zap.DebugLevel
		case log.LevelInfo:
			zapLevel = zap.InfoLevel
		case log.LevelWarn:
			zapLevel = zap.WarnLevel
		case log.LevelError:
			zapLevel = zap.ErrorLevel
		case log.LevelFatal:
			zapLevel = zap.FatalLevel
		}
		var fields []zap.Field
		for i := 0; i < len(keyvals); i += 2 {
			fields = append(fields, zap.String(fmt.Sprintf("%v", keyvals[i]), fmt.Sprintf("%v", keyvals[i+1])))
		}
		logger.Log(zapLevel, "", fields...)
		return nil
	}), nil
}
