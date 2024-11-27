package user

import (
	"app/app/base/messages"
	"app/app/entities"
	userdto "app/app/modules/user/dto"
	"app/app/util/hashing"
	"context"
	"errors"
	"strings"

	"github.com/uptrace/bun"
)

type UserService struct {
	db *bun.DB
}

func newService(db *bun.DB) *UserService {
	return &UserService{
		db: db,
	}
}

func (s *UserService) Create(ctx context.Context, req userdto.CreateUserRequest) (*entities.UserEntity, bool, error) {
	m := entities.UserEntity{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
	}

	var mserr bool
	err := s.db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {

		_, err := tx.NewInsert().Model(&m).Exec(ctx)
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key value") {
				mserr = true
				return errors.New(messages.EmailAlreadyExists)
			}
		}

		password, err := hashing.HashPassword(req.Password)
		if err != nil {
			return err
		}

		userPass := entities.UserPasswordEntity{
			UserID:   m.ID,
			Password: string(password),
		}

		_, err = tx.NewInsert().Model(&userPass).Exec(ctx)

		return err
	})

	return &m, mserr, err
}

func (s *UserService) List(ctx context.Context, req userdto.GetUserListRequest) ([]entities.UserEntity, int, error) {

	m := []entities.UserEntity{}
	err := s.db.NewSelect().Model(&m).Scan(ctx)
	if err != nil {
		return nil, 0, err
	}

	return m, len(m), nil

	// m := []userdto.UserResponse{}

	// var (
	// 	offset = (req.Page - 1) * req.Size
	// 	limit  = req.Size
	// )

	// query := s.db.NewSelect().Model((*entities.UserEntity)(nil)).Column("id", "first_name", "last_name", "email")

	// if req.Search != "" {
	// 	search := fmt.Sprint("%" + strings.ToLower(req.Search) + "%")
	// 	query.WhereGroup(" AND ", func(q *bun.SelectQuery) *bun.SelectQuery {
	// 		return q.WhereOr(`LOWER(first_name) Like ?`, search).
	// 			WhereOr(`LOWER(last_name) Like ?`, search).
	// 			WhereOr(`LOWER(email) Like ?`, search)
	// 	})
	// }

	// count, err := query.Count(ctx)
	// if count == 0 {
	// 	return m, 0, err
	// }

	// order := fmt.Sprintf("%s %s", req.SortBy, req.OrderBy)
	// err = query.Offset(offset).Limit(limit).Order(order).Scan(ctx, &m)

	// return m, count, err
}
