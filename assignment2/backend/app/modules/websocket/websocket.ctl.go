package websocket

import (
	"app/app/base"

	"github.com/gin-gonic/gin"
)

type WebSocketController struct {
	websocketHub *WebsocketHub
}

func newController(websocketHub *WebsocketHub) *WebSocketController {
	return &WebSocketController{
		websocketHub: websocketHub,
	}
}

func (c *WebSocketController) GetUserActive(ctx *gin.Context) {

	users := []User{}
	for _, user := range c.websocketHub.hub.Users {

		users = append(users, *user)
	}

	base.Success(ctx, users)
}
