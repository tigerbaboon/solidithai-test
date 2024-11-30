package message

import "github.com/uptrace/bun"

type MessageModule struct {
	Ctl *MessageController
	Svc *MessageService
}

func New(db *bun.DB) *MessageModule {

	svc := newService(db)
	return &MessageModule{
		Ctl: newController(svc),
		Svc: svc,
	}
}
