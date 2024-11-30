package websocket

import (
	"app/app/modules/user"

	"github.com/uptrace/bun"
)

type WebsocketModule struct {
	Client       *WebsocketClient
	WebsocketHub *WebsocketHub
	Ctl          *WebSocketController
}

func New(db *bun.DB, user *user.UserModule) *WebsocketModule {

	websocketHub := websocketHub(db, user)

	return &WebsocketModule{
		Client:       client(websocketHub),
		WebsocketHub: websocketHub,
		Ctl:          newController(websocketHub),
	}
}
