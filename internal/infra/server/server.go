package server

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/hizzuu/beatic-backend/conf"
	"github.com/hizzuu/beatic-backend/graph/generated"
	"github.com/hizzuu/beatic-backend/internal/interface/resolver"
)

type server struct {
	port string
}

func New() *server {
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver.Resolver{}}))
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
		log.Println("========X-Cloud-Trace-Context=======", header)

		next.ServeHTTP(w, r)
	})
}
