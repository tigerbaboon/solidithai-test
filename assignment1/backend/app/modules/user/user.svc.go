package user

import (
	"app/app/base/messages"
	"app/app/entities"
	userdto "app/app/modules/user/dto"
	"app/app/util/hashing"
	"context"
	"database/sql"
	"errors"
	"fmt"
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

	fmt.Println(mserr)

	return &m, mserr, err
}

func (s *UserService) Update(ctx context.Context, id userdto.GetUserByIDRequest, req userdto.UpdateUserRequest) (*entities.UserEntity, bool, error) {

	exUser, err := s.db.NewSelect().Model((*entities.UserEntity)(nil)).Where("id = ?", id.ID).Exists(ctx)
	if err != nil {
		return nil, false, err
	}

	if !exUser {
		return nil, false, errors.New(messages.UserNotFound)
	}

	m := entities.UserEntity{
		ID:        id.ID,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
	}

	m.SetUpdateNow()

	_, err = s.db.NewUpdate().Model(&m).
		Set("first_name = ?", m.FirstName).
		Set("last_name = ?", m.LastName).
		Set("email = ?", m.Email).
		WherePK().
		OmitZero().
		Returning("*").
		Exec(ctx)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value") {
			return nil, true, errors.New(messages.EmailAlreadyExists)
		}
		return nil, false, err
	}

	return &m, false, nil
}

func (s *UserService) Delete(ctx context.Context, id userdto.GetUserByIDRequest) (*entities.UserEntity, bool, error) {

	exUser, err := s.db.NewSelect().Model((*entities.UserEntity)(nil)).Where("id = ?", id.ID).Exists(ctx)
	if err != nil {
		return nil, false, err
	}

	if !exUser {
		return nil, false, errors.New(messages.UserNotFound)
	}

	var mserr bool
	err = s.db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {

		_, err := tx.NewDelete().Model(&entities.UserEntity{}).Where("id = ?", id.ID).Exec(ctx)
		if err != nil {
			return err
		}

		_, err = tx.NewDelete().Model(&entities.UserPasswordEntity{}).Where("user_id = ?", id.ID).Exec(ctx)

		return err
	})

	return nil, mserr, err
}

func (s *UserService) Get(ctx context.Context, id userdto.GetUserByIDRequest) (*userdto.UserResponse, error) {

	m := userdto.UserResponse{}

	err := s.db.NewSelect().Model((*entities.UserEntity)(nil)).
		Column("id", "first_name", "last_name", "email").
		Where("id = ?", id.ID).
		Scan(ctx, &m)
	if err != nil {
		return nil, err
	}

	return &m, nil
}

func (s *UserService) ExistByEmail(ctx context.Context, email string) (*entities.UserEntity, error) {

	m := entities.UserEntity{}

	err := s.db.NewSelect().Model(&m).
		Where("email = ?", email).
		Scan(ctx)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
	}

	return &m, nil
}

func (s *UserService) Info(ctx context.Context, userID string) (*userdto.UserResponse, error) {

	m := userdto.UserResponse{}

	err := s.db.NewSelect().Model((*entities.UserEntity)(nil)).
		Column("id", "first_name", "last_name", "email").
		Where("id = ?", userID).
		Scan(ctx, &m)
	if err != nil {
		return nil, err
	}

	return &m, nil
}

func (s *UserService) List(ctx context.Context, req userdto.GetUserListRequest) ([]userdto.UserResponse, int, error) {

	m := []userdto.UserResponse{}

	var (
		offset = (req.Page - 1) * req.Size
		limit  = req.Size
	)

	query := s.db.NewSelect().Model((*entities.UserEntity)(nil)).
		Column("id", "first_name", "last_name", "email")

	if req.Search != "" {
		search := fmt.Sprint("%" + strings.ToLower(req.Search) + "%")
		query.WhereGroup(" AND ", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.WhereOr(`LOWER(first_name) Like ?`, search).
				WhereOr(`LOWER(last_name) Like ?`, search).
				WhereOr(`LOWER(email) Like ?`, search)
		})
	}

	count, err := query.Count(ctx)
	if count == 0 {
		return m, 0, err
	}

	order := fmt.Sprintf("%s %s", req.SortBy, req.OrderBy)
	err = query.Offset(offset).Limit(limit).Order(order).Scan(ctx, &m)

	return m, count, err
}

func (s *UserService) UpdatePassword(ctx context.Context, id userdto.GetUserByIDRequest, req userdto.UpdatePasswordRequest) (*entities.UserPasswordEntity, bool, error) {

	exUser, err := s.db.NewSelect().Model((*entities.UserEntity)(nil)).Where("id = ?", id.ID).Exists(ctx)
	if err != nil {
		return nil, false, err
	}

	if !exUser {
		return nil, false, errors.New(messages.UserNotFound)
	}

	password, err := hashing.HashPassword(req.Password)
	if err != nil {
		return nil, false, err
	}

	userPass := entities.UserPasswordEntity{
		UserID:   id.ID,
		Password: string(password),
	}

	userPass.SetUpdateNow()

	_, err = s.db.NewUpdate().Model(&userPass).
		Set("password = ?", userPass.Password).
		Set("updated_at = ?", userPass.UpdatedAt).
		Where("user_id = ?", userPass.UserID).
		Exec(ctx)

	return nil, false, err
}

func (s *UserService) ExistPassword(ctx context.Context, userID string) (*entities.UserPasswordEntity, error) {

	m := entities.UserPasswordEntity{}

	err := s.db.NewSelect().Model(&m).
		Where("user_id = ?", userID).
		Scan(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
	}

	return &m, nil
}
