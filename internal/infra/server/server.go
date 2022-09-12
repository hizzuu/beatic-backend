package server

import (
	"context"
	"net/http"
	"regexp"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/hizzuu/beatic-backend/conf"
	"github.com/hizzuu/beatic-backend/graph/generated"
	"github.com/hizzuu/beatic-backend/graph/model"
	"github.com/hizzuu/beatic-backend/internal/infra/logger"
	"github.com/hizzuu/beatic-backend/internal/interface/resolver"
)

type server struct {
	port string
}

func New(l logger.Logger) *server {
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver.New(l)}))
	http.Handle("/query", tracer(srv))
	http.Handle("/", tracer(playground.Handler("GraphQL playground", "/query")))

	return &server{port: conf.C.Api.Port}
}

func (s *server) ListenAndServe() error {
	return http.ListenAndServe(":"+s.port, nil)
}

func tracer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("X-Cloud-Trace-Context")
		// "TRACE_ID" + "/SPAN_ID" + ;0=TRACE_TRUE
		m := regexp.MustCompile(`([a-f\d]+)?` + `(?:/([a-f\d]+))?` + `(?:;o=(\d))?`).FindStringSubmatch(header)
		traceID, spanID, sampled := m[1], m[2], m[3] == "1"
		next.ServeHTTP(
			w,
			r.WithContext(
				context.WithValue(
					r.Context(),
					model.TracerCtxKey,
					&model.Trace{
						TraceID: traceID,
						SpanID:  spanID,
						Sampled: sampled,
					},
				),
			),
		)
	})
}
