package auth

import "github.com/aide-family/sovereign/pkg/config"

type User interface {
	GetOpenID() string
	GetName() string
	GetEmail() string
	GetAvatar() string
	GetAPP() config.OAuth2_APP
	GetRaw() []byte
}
