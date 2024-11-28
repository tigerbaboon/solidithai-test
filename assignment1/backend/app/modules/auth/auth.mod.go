package auth

import (
	"app/app/modules/accesstoken"
	"app/app/modules/user"

	"github.com/uptrace/bun"
)

type AuthModule struct {
	Ctl *AuthController
	Svc *AuthService
}

func New(db *bun.DB, user *user.UserModule, accesstoken *accesstoken.AccessTokenModule) *AuthModule {

	svc := newService(db, user, accesstoken)
	return &AuthModule{
		Ctl: newController(svc),
		Svc: svc,
	}
}
