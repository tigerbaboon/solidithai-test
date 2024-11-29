package entities

import "github.com/uptrace/bun"

type AccessTokenEntity struct {
	bun.BaseModel `bun:"table:access_tokens"`

	UserID string `bun:"user_id,notnull,unique"`
	Token  string `bun:"token,type:text,notnull"`

	CreateUpdateUnixTimestamp
}
