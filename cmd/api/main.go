package main

import (
	"context"

	"github.com/hizzuu/beatic-backend/internal/infra/logger"
	"github.com/hizzuu/beatic-backend/internal/infra/server"
)

func main() {
	ctx := context.Background()
	l, err := logger.New()
	if err != nil {
		panic(err)
	}

	s := server.New(l)
	l.Errorf(ctx, s.ListenAndServe().Error())
}
