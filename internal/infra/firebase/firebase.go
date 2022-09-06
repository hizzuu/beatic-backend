package firebase

import (
	"context"

	fb "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/hizzuu/beatic-backend/conf"
	"google.golang.org/api/option"
)

type firebase struct {
	app *fb.App
}

type Firebase interface {
	Auth(ctx context.Context) (*auth.Client, error)
}

func New(ctx context.Context) (*firebase, error) {
	app, err := fb.NewApp(ctx, nil, option.WithCredentialsJSON([]byte(conf.C.Credentials.Firebase.SecretKey)))
	if err != nil {
		return nil, err
	}

	return &firebase{app: app}, nil
}

func (f *firebase) Auth(ctx context.Context) (*auth.Client, error) {
	return f.app.Auth(ctx)
}
