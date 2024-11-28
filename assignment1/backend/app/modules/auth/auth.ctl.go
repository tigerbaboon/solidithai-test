package auth

import (
	"app/app/base"
	"app/app/base/messages"
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

func (ctl *AuthController) Login(ctx *gin.Context) {
	req := authdto.LoginRequest{}
	if err := ctx.Bind(&req); err != nil {
		log.Error(err.Error())
		base.BadRequest(ctx, messages.InvalidRequest, base.Galidator.Validator(req).DecryptErrors(err))
		return
	}

	data, mserr, err := ctl.authSvc.Login(ctx, req)
	if err != nil {
		ms := messages.InternalError
		if mserr {
			ms = err.Error()
		}
		log.Error(err.Error())
		base.InternalServerError(ctx, ms, nil)
		return
	}

	base.Created(ctx, data)
}
