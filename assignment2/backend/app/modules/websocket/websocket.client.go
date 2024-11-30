package websocket

import (
	"app/app/base"
	"app/app/base/messages"
	userdto "app/app/modules/user/dto"
	"app/config/log"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	ws "github.com/gorilla/websocket"
)

type WebsocketClient struct {
	WebsocketHub *WebsocketHub
}

func client(websocketHub *WebsocketHub) *WebsocketClient {
	return &WebsocketClient{
		WebsocketHub: websocketHub,
	}
}

type Client struct {
	UserID    string
	FirstName string
	LastName  string
	Avatar    string
	Conn      *ws.Conn
	Send      chan Send
	Hub       *Hub
	Context   context.Context
}

func (c *Client) ReadPump() {
	defer func() {
		c.Hub.Unregister <- c
		c.Conn.Close()
	}()
	for {

		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			break
		}

		msg := Message{
			UserID:    c.UserID,
			FirstName: c.FirstName,
			LastName:  c.LastName,
			Avatar:    c.Avatar,
			Text:      string(message),
		}
		c.Hub.Broadcast <- msg
	}
}

func (c *Client) WritePump() {
	defer func() {
		c.Conn.Close()
	}()

	for {
		select {
		case msg, ok := <-c.Send:
			if !ok {
				c.Conn.WriteMessage(ws.CloseMessage, []byte{})
				return
			}
			c.Conn.WriteJSON(msg)
		}
	}
}

var upgrader = ws.Upgrader{
	// ReadBufferSize:  1024,
	// WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (c WebsocketClient) ServeWS(ctx *gin.Context) {

	userID := ctx.Query("userID")

	if userID == "" {
		base.InternalServerError(ctx, messages.InvalidRequest, nil)
		return
	}
	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Error(err.Error())
		base.InternalServerError(ctx, messages.InternalError, nil)
		return
	}

	user, err := c.WebsocketHub.user.Svc.Get(ctx, userdto.GetUserByIDRequest{ID: userID})
	if err != nil {
		conn.Close()
		log.Error(err.Error())
		base.InternalServerError(ctx, messages.InternalError, nil)
		return
	}

	client := &Client{
		UserID:    userID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Avatar:    user.Avatar,
		Conn:      conn,
		Send:      make(chan Send),
		Hub:       c.WebsocketHub.hub,
		Context:   ctx,
	}

	client.Hub.Register <- client

	go client.WritePump()
	go client.ReadPump()

}
