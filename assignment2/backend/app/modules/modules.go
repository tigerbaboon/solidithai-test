package modules

import (
	"app/app/modules/accesstoken"
	"app/app/modules/auth"
	"app/app/modules/message"
	"app/app/modules/user"
	"app/app/modules/websocket"
	"app/config"

	"github.com/uptrace/bun"
)

type Modules struct {
	DB          *bun.DB
	User        *user.UserModule
	Auth        *auth.AuthModule
	Accesstoken *accesstoken.AccessTokenModule
	Websocket   *websocket.WebsocketModule
	Message     *message.MessageModule
}

func Get() *Modules {

	config.Init()
	db := config.Database()

	user := user.New(db)
	accesstoken := accesstoken.New(db)
	auth := auth.New(db, user, accesstoken)
	message := message.New(db)
	websocket := websocket.New(db, user)

	return &Modules{
		DB:          db,
		User:        user,
		Auth:        auth,
		Accesstoken: accesstoken,
		Websocket:   websocket,
		Message:     message,
	}
}
