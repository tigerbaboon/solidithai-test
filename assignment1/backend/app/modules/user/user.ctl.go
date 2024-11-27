package user

import (
	"app/app/base"
	"app/app/base/messages"
	userdto "app/app/modules/user/dto"
	"app/app/util/log"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userSvc *UserService
}

func newController(userService *UserService) *UserController {
	return &UserController{
		userSvc: userService,
	}
}

func (c *UserController) Create(ctx *gin.Context) {

	req := userdto.CreateUserRequest{}

	if err := ctx.Bind(&req); err != nil {
		log.Error(err.Error())
		base.BadRequest(ctx, messages.InvalidRequest, base.Galidator.Validator(req).DecryptErrors(err))
		return
	}

	data, mserr, err := c.userSvc.Create(ctx, req)
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

func (c *UserController) List(ctx *gin.Context) {
	req := userdto.GetUserListRequest{}
	if err := ctx.Bind(&req); err != nil {
		log.Error(err.Error())
		base.BadRequest(ctx, messages.InvalidRequest, nil)
		return
	}

	if req.Page == 0 {
		req.Page = base.DEFAULT_PAGINATION.Page
	}

	if req.Page == 0 {
		req.Page = base.DEFAULT_PAGINATION.Size
	}

	if req.OrderBy == "" {
		req.OrderBy = base.DEFAULT_PAGINATION.OrderBy
	}

	if req.SortBy == "" {
		req.SortBy = base.DEFAULT_PAGINATION.SortBy
	}

	data, count, err := c.userSvc.List(ctx, req)
	if err != nil {
		log.Error(err.Error())
		base.InternalServerError(ctx, messages.InternalError, nil)
		return
	}

	base.Paginate(ctx, data, &base.ResponsePaginate{
		Page:  int64(req.Page),
		Size:  int64(req.Size),
		Total: int64(count),
	})
}