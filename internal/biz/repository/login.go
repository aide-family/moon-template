package repository

import (
	"context"

	"golang.org/x/oauth2"

	"github.com/aide-family/sovereign/pkg/api/auth"
)

type LoginRepository interface {
	Login(ctx context.Context, oauthConfig *oauth2.Config, user auth.User) (string, error)
}
