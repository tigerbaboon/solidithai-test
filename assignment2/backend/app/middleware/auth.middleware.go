package middleware

import (
	"app/app/base"
	"app/app/base/messages"
	"app/app/modules"
	"app/app/util/jwt"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(mod *modules.Modules) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			base.Unauthorized(ctx, messages.Unauthorized, nil)
			ctx.Abort()
			return
		}

		authParts := strings.Split(authHeader, " ")
		if len(authParts) != 2 || authParts[0] != "Bearer" {
			base.Unauthorized(ctx, messages.Unauthorized, nil)
			ctx.Abort()
			return
		}

		token := authParts[1]
		data, err := jwt.Verify(token)
		if err != nil {
			base.Unauthorized(ctx, messages.Unauthorized, nil)
			ctx.Abort()
			return
		}

		if !mod.Accesstoken.Svc.Verify(ctx, data.UserID, token) {
			base.Unauthorized(ctx, messages.Unauthorized, nil)
			ctx.Abort()
			return
		}

		ctx.Set("userID", data.UserID)

		ctx.Next()
	}
}
