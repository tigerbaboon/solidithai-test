package entities

import "time"

type CreateUpdateUnixTimestamp struct {
	CreateUnixTimestamp
	UpdateUnixTimestamp
}

type CreateUnixTimestamp struct {
	CreatedAt int64 `bun:"created_at,notnull,default:EXTRACT(EPOCH FROM NOW())"`
}

type UpdateUnixTimestamp struct {
	UpdatedAt int64 `bun:"updated_at,notnull,default:EXTRACT(EPOCH FROM NOW())"`
}

type SoftDelete struct {
	DeletedAt *time.Time `bun:"deleted_at,soft_delete,nullzero"`
}

func (t *UpdateUnixTimestamp) SetUpdate(ts int64) {
	t.UpdatedAt = ts
}

func (t *UpdateUnixTimestamp) SetUpdateNow() {
	t.SetUpdate(time.Now().Unix())
}
