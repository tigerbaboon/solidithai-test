package accesstoken

import (
	"github.com/uptrace/bun"
)

type AccessTokenModule struct {
	Svc *AccessTokenService
}

func New(db *bun.DB) *AccessTokenModule {

	svc := newService(db)
	return &AccessTokenModule{
		Svc: svc,
	}
}
