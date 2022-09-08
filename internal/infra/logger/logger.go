package logger

import (
	"context"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger interface {
	Info(ctx context.Context, msg string)
	Infof(ctx context.Context, template string, args ...interface{})
	Error(ctx context.Context, e error)
}

type logger struct {
	logger *zap.Logger
}

var logLevelSeverity = map[zapcore.Level]string{
	zapcore.DebugLevel:  "DEBUG",
	zapcore.InfoLevel:   "INFO",
	zapcore.WarnLevel:   "WARNING",
	zapcore.ErrorLevel:  "ERROR",
	zapcore.DPanicLevel: "CRITICAL",
	zapcore.PanicLevel:  "ALERT",
	zapcore.FatalLevel:  "EMERGENCY",
}

func New() (*logger, error) {
	l, err := zap.Config{
		Level:       zap.NewAtomicLevelAt(zapcore.InfoLevel),
		Development: false,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding: "json",
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:      "severity",
			NameKey:       "logger",
			CallerKey:     "caller",
			StacktraceKey: "stack_trace",
			TimeKey:       "time",
			MessageKey:    "message",
			LineEnding:    zapcore.DefaultLineEnding,
			EncodeTime:    zapcore.RFC3339NanoTimeEncoder,
			EncodeLevel: func(l zapcore.Level, pae zapcore.PrimitiveArrayEncoder) {
				pae.AppendString(logLevelSeverity[l])
			},
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}.Build(zap.AddCallerSkip(1))
	if err != nil {
		return nil, err
	}

	l.With()

	return &logger{logger: l}, nil
}

func (l *logger) Info(ctx context.Context, msg string) {
	l.logger.Sugar().Info(msg)
}

func (l *logger) Infof(ctx context.Context, template string, args ...interface{}) {
	l.logger.Sugar().Infof(template, args...)
}

func (l *logger) Error(ctx context.Context, e error) {
	l.logger.Sugar().Error(e)
}
