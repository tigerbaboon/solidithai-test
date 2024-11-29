package socket

import "github.com/uptrace/bun"

type UserModule struct {
	Client *WebsocketClient
	Hub    *WebsocketHub
}

func New(db *bun.DB) *UserModule {

	return &UserModule{
		Client: client(db),
		Hub:    hub(db),
	}
}
