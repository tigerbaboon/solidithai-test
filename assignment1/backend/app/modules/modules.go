package modules

import (
	"app/app/modules/accesstoken"
	"app/app/modules/auth"
	"app/app/modules/user"
	"app/config"

	"github.com/uptrace/bun"
)

type Modules struct {
	DB          *bun.DB
	User        *user.UserModule
	Auth        *auth.AuthModule
	Accesstoken *accesstoken.AccessTokenModule
}

func Get() *Modules {

	config.Init()
	db := config.Database()

	user := user.New(db)
	accesstoken := accesstoken.New(db)
	auth := auth.New(db, user, accesstoken)

	return &Modules{
		DB:          db,
		User:        user,
		Auth:        auth,
		Accesstoken: accesstoken,
	}
}
