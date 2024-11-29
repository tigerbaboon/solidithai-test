package auth

import (
	"app/app/base"
	"app/app/base/messages"
	"app/app/helper"
	authdto "app/app/modules/auth/dto"
	"app/config/log"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authSvc *AuthService
}

func newController(authService *AuthService) *AuthController {
	return &AuthController{
		authSvc: authService,
	}
}

func (c *AuthController) Login(ctx *gin.Context) {
	req := authdto.LoginRequest{}
	if err := ctx.Bind(&req); err != nil {
		log.Error(err.Error())
		base.BadRequest(ctx, messages.InvalidRequest, base.Galidator.Validator(req).DecryptErrors(err))
		return
	}

	data, mserr, err := c.authSvc.Login(ctx, req)
	if err != nil {
		ms := messages.InternalError
		if mserr {
			ms = err.Error()
		}
		log.Error(err.Error())
		base.InternalServerError(ctx, ms, nil)
		return
	}

	base.Success(ctx, data)
}

func (c *AuthController) Logout(ctx *gin.Context) {
	user, err := helper.GetAuthorzied(ctx)
	if err != nil {
		log.Error(err.Error())
		base.InternalServerError(ctx, messages.InternalError, nil)
		return
	}

	err = c.authSvc.Logout(ctx, *user)
	if err != nil {
		log.Error(err.Error())
		base.InternalServerError(ctx, messages.InternalError, nil)
		return
	}

	base.Success(ctx, nil)
}
