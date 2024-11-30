package message

import (
	"app/app/base/messages"
	"app/app/entities"
	messagedto "app/app/modules/message/dto"
	"app/config/log"
	"context"
	"errors"
	"fmt"

	"github.com/uptrace/bun"
)

type MessageService struct {
	db *bun.DB
}

func newService(db *bun.DB) *MessageService {
	return &MessageService{
		db: db,
	}
}

func (s *MessageService) Create(ctx context.Context, req messagedto.CreateMessageRequest) (*entities.MessageEntity, bool, error) {

	exUser, err := s.db.NewSelect().Model(&entities.UserEntity{}).Where("id = ?", req.UserID).Exists(ctx)
	if err != nil {
		log.Error(err.Error())
		return nil, false, err
	}

	if !exUser {
		return nil, true, errors.New(messages.UserNotFound)
	}

	m := entities.MessageEntity{
		UserID: req.UserID,
		Text:   req.Text,
	}

	_, err = s.db.NewInsert().Model(&m).Exec(ctx)

	return nil, false, err
}

func (s *MessageService) List(ctx context.Context, req messagedto.GetMessageListRequest) ([]messagedto.MessageResponse, int, error) {

	m := []messagedto.MessageResponse{}

	var (
		offset = (req.Page - 1) * req.Size
		limit  = req.Size
	)

	query := s.db.NewSelect().TableExpr("messages AS m").
		Column("m.id", "m.text", "m.created_at").
		ColumnExpr("u.id as user_id").
		ColumnExpr("u.first_name").
		ColumnExpr("u.last_name").
		ColumnExpr("u.avatar").
		Join("LEFT JOIN users AS u ON u.id = m.user_id")

	count, err := query.Count(ctx)
	if count == 0 {
		return m, 0, err
	}

	order := fmt.Sprintf("m.%s %s", req.SortBy, req.OrderBy)
	err = query.Offset(offset).Limit(limit).Order(order).Scan(ctx, &m)

	return m, count, err

}
