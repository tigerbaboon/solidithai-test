package routes

import (
	"app/app/modules"

	"github.com/gin-gonic/gin"
)

func api(r *gin.RouterGroup, mod *modules.Modules) {

	users := r.Group("/users")
	{
		users.POST("", mod.User.Ctl.Create)
		users.GET("", mod.User.Ctl.List)
	}
}
