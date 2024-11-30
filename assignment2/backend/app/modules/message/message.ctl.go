package message

import (
	"app/app/base"
	"app/app/base/messages"
	messagedto "app/app/modules/message/dto"
	"app/config/log"

	"github.com/gin-gonic/gin"
)

type MessageController struct {
	messageSvc *MessageService
}

func newController(messageService *MessageService) *MessageController {
	return &MessageController{
		messageSvc: messageService,
	}
}

func (c *MessageController) Create(ctx *gin.Context) {
	req := messagedto.CreateMessageRequest{}
	if err := ctx.Bind(&req); err != nil {
		log.Error(err.Error())
		base.BadRequest(ctx, messages.InvalidRequest, nil)
		return
	}

	data, mserr, err := c.messageSvc.Create(ctx, req)
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

func (c *MessageController) List(ctx *gin.Context) {
	req := messagedto.GetMessageListRequest{}
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

	data, count, err := c.messageSvc.List(ctx, req)
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
