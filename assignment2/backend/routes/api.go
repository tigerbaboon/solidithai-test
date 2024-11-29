package routes

import (
	"app/app/middleware"
	"app/app/modules"

	"github.com/gin-gonic/gin"
)

func api(r *gin.RouterGroup, mod *modules.Modules) {

	md := middleware.AuthMiddleware(mod)

	auth := r.Group("/auth")
	{
		auth.POST("/login", mod.Auth.Ctl.Login)
		auth.POST("/logout", md, mod.Auth.Ctl.Logout)
	}

	users := r.Group("/users", md)
	{
		users.POST("", mod.User.Ctl.Create)
		users.PUT("/:id", mod.User.Ctl.Update)
		users.DELETE("/:id", mod.User.Ctl.Delete)
		users.GET("/:id", mod.User.Ctl.Get)
		users.GET("", mod.User.Ctl.List)

		users.GET("/info", mod.User.Ctl.Info)
		users.PATCH("/:id/password", mod.User.Ctl.UpdatePassword)
	}

	ws := r.Group("/ws")
	{
		ws.GET("")
	}
}
