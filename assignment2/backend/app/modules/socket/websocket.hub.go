package socket

import (
	"sync"

	"github.com/uptrace/bun"
)

type WebsocketHub struct {
	db *bun.DB
}

func hub(db *bun.DB) *WebsocketHub {
	return &WebsocketHub{
		db: db,
	}
}

type Hub struct {
	register   chan *Client
	unregister chan *Client
	mu         sync.Mutex
}

func (h *WebsocketHub) Start() {
	// Start the websocket hub

}
