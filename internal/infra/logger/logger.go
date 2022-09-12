package logger

import (
	"context"
	"fmt"

	"github.com/hizzuu/beatic-backend/conf"
	"github.com/hizzuu/beatic-backend/graph/model"
	"github.com/hizzuu/beatic-backend/internal/util/environment"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger interface {
	Infof(ctx context.Context, template string, args ...interface{})
	Errorf(ctx context.Context, template string, args ...interface{})
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

	return &logger{logger: l}, nil
}

func (l *logger) Infof(ctx context.Context, template string, args ...interface{}) {
	l.logger.With(trace(ctx)...).Info(fmt.Sprintf(template, args...))
}

func (l *logger) Errorf(ctx context.Context, templete string, args ...interface{}) {
	l.logger.Error(fmt.Sprintf(templete, args...))
}

func trace(ctx context.Context) []zapcore.Field {
	if !environment.IsProd() {
		return nil
	}

	id := conf.C.Credentials.GCP.ProjectID
	t, _ := ctx.Value(model.TracerCtxKey).(*model.Trace)
	return []zapcore.Field{
		zap.String("logging.googleapis.com/trace", fmt.Sprintf("projects/%s/traces/%s", id, t.TraceID)),
		zap.String("logging.googleapis.com/spanId", fmt.Sprintf("projects/%s/traces/%s", id, t.SpanID)),
		zap.String("logging.googleapis.com/trace_sampled", fmt.Sprintf("projects/%s/traces/%t", id, t.Sampled)),
	}
}
