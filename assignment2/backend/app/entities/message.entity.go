package entities

import "github.com/uptrace/bun"

type MessageEntity struct {
	bun.BaseModel `bun:"table:messages"`

	ID     string `bun:",default:gen_random_uuid(),pk"`
	UserID string `bun:"user_id,notnull"`
	Text   string `bun:"text,notnull"`

	CreateUpdateUnixTimestamp
}
