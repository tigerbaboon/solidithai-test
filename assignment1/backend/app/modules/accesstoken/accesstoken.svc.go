package accesstoken

import (
	"app/app/entities"
	"context"

	"github.com/uptrace/bun"
)

type AccessTokenService struct {
	db *bun.DB
}

func newService(db *bun.DB) *AccessTokenService {
	return &AccessTokenService{
		db: db,
	}
}

func (s *AccessTokenService) Create(ctx context.Context, userID string, token string) error {

	m := entities.AccessTokenEntity{
		UserID: userID,
		Token:  token,
	}

	_, err := s.db.NewInsert().Model(&m).On("CONFLICT (user_id) DO UPDATE").Set("token = EXCLUDED.token").Exec(ctx)

	return err
}

func (s *AccessTokenService) Verify(ctx context.Context, userID string, token string) bool {

	ex, err := s.db.NewSelect().Model(&entities.AccessTokenEntity{}).Where("user_id = ? and token = ?", userID, token).Exists(ctx)
	if err != nil {
		return false
	}

	return ex
}

func (s *AccessTokenService) Delete(ctx context.Context, userID string) error {

	_, err := s.db.NewDelete().Model(&entities.AccessTokenEntity{}).Where("user_id = ?", userID).Exec(ctx)

	return err
}
