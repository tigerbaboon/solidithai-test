package socket

import (
	"github.com/uptrace/bun"
)

type WebsocketClient struct {
	db *bun.DB
}

func client(db *bun.DB) *WebsocketClient {
	return &WebsocketClient{
		db: db,
	}
}

type Client struct {
	hub      *Hub
	username string
}
