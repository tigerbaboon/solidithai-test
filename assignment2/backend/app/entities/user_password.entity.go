package entities

import "github.com/uptrace/bun"

type UserPasswordEntity struct {
	bun.BaseModel `bun:"table:user_passwords"`

	UserID   string `bun:"user_id,notnull"`
	Password string `bun:"password,type:text,notnull"`

	CreateUpdateUnixTimestamp
}
