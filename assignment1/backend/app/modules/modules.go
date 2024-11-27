package modules

import (
	"app/app/modules/user"
	"app/config"

	"github.com/uptrace/bun"
)

type Modules struct {
	DB   *bun.DB
	User *user.UserModule
}

func Get() *Modules {

	config.Init()
	db := config.Database()

	user := user.New(db)

	return &Modules{
		DB:   db,
		User: user,
	}
}
