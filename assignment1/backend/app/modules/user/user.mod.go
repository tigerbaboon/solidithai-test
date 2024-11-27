package user

import "github.com/uptrace/bun"

type UserModule struct {
	Ctl *UserController
	Svc *UserService
}

func New(db *bun.DB) *UserModule {

	svc := newService(db)
	return &UserModule{
		Ctl: newController(svc),
		Svc: svc,
	}
}
