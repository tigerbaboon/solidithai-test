package websocket

import (
	"app/app/modules/user"
	"sync"

	"github.com/uptrace/bun"
)

type WebsocketHub struct {
	db   *bun.DB
	hub  *Hub
	user *user.UserModule
}

func websocketHub(db *bun.DB, user *user.UserModule) *WebsocketHub {
	runhub := NewHub()
	go runhub.Run()
	return &WebsocketHub{
		db:   db,
		hub:  runhub,
		user: user,
	}
}

type Hub struct {
	Clients    map[*Client]bool
	Users      map[string]*User
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan Message
	Mu         sync.Mutex
}

type Send struct {
	Type    string  `json:"type"`
	Users   []User  `json:"users"`
	Message Message `json:"message"`
}

type User struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Avatar    string `json:"avatar"`
}

type Message struct {
	UserID    string `json:"userId"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Avatar    string `json:"avatar"`
	Text      string `json:"text"`
}

func NewHub() *Hub {
	return &Hub{
		Clients:    make(map[*Client]bool),
		Users:      make(map[string]*User),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan Message),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.Clients[client] = true
			h.Users[client.UserID] = &User{
				ID:        client.UserID,
				FirstName: client.FirstName,
				LastName:  client.LastName,
				Avatar:    client.Avatar,
			}
			h.HandleActive()
		case client := <-h.Unregister:
			if _, ok := h.Clients[client]; ok {
				delete(h.Clients, client)
				delete(h.Users, client.UserID)
				close(client.Send)
				h.HandleActive()
			}
		case message := <-h.Broadcast:

			h.HandleMessage(message)
		}
	}
}

func (h *Hub) HandleMessage(message Message) {

	for client := range h.Clients {
		temp := Send{
			Type:    "message",
			Message: message,
		}

		client.Send <- temp
	}
}

func (h *Hub) HandleActive() {
	// h.Mu.Lock()
	// defer h.Mu.Unlock()
	users := []User{}
	for _, user := range h.Users {

		users = append(users, *user)
	}

	for client := range h.Clients {
		temp := Send{
			Type:  "users",
			Users: users,
		}

		client.Send <- temp
	}
}
