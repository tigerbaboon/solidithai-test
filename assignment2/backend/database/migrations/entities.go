package migrations

import "app/app/entities"

func Entities() []any {
	return []any{
		(*entities.UserEntity)(nil),
		(*entities.UserPasswordEntity)(nil),
		(*entities.AccessTokenEntity)(nil),
		(*entities.MessageEntity)(nil),
	}
}
