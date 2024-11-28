package entities

import "github.com/uptrace/bun"

type UserEntity struct {
	bun.BaseModel `bun:"table:users"`

	ID        string `bun:",default:gen_random_uuid(),pk"`
	FirstName string `bun:"first_name,notnull"`
	LastName  string `bun:"last_name,notnull"`
	Email     string `bun:"email,notnull"`

	CreateUpdateUnixTimestamp
	SoftDelete
}
