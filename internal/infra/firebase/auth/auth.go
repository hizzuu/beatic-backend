package auth

import (
	"context"

	au "firebase.google.com/go/auth"
	"github.com/hizzuu/beatic-backend/internal/infra/firebase"
)

type auth struct {
	client *au.Client
}

type Auth interface {
	GetUser(ctx context.Context, uid string) (*au.UserRecord, error)
	VerifyIDToken(ctx context.Context, token string) (*au.Token, error)
	SetCustomClaims(ctx context.Context, uid string, id int) error
}

func New(ctx context.Context, f firebase.Firebase) (*auth, error) {
	client, err := f.Auth(ctx)
	if err != nil {
		return nil, err
	}

	return &auth{client: client}, nil
}

func (a *auth) GetUser(ctx context.Context, uid string) (*au.UserRecord, error) {
	return a.client.GetUser(ctx, uid)
}

func (a *auth) VerifyIDToken(ctx context.Context, token string) (*au.Token, error) {
	return a.client.VerifyIDToken(ctx, token)
}

func (a *auth) SetCustomClaims(ctx context.Context, uid string, id int64) error {
	claims := map[string]interface{}{"userID": id}
	if err := a.client.SetCustomUserClaims(ctx, uid, claims); err != nil {
		return err
	}

	return nil
}
