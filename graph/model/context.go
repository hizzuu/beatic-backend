package model

type contextKey string

const (
	TracerCtxKey contextKey = "tracer"
)

type Trace struct {
	TraceID string
	SpanID  string
	Sampled bool
}
