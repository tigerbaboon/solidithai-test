package seeds

import (
	"app/app/modules"
	userdto "app/app/modules/user/dto"
	"context"
)

func adminSeed(ctx context.Context, mod *modules.Modules) error {

	admin := userdto.CreateUserRequest{
		FirstName: "Admin",
		LastName:  "Admin",
		Email:     "admin@mail.com",
		Password:  "123456789",
	}

	_, _, err := mod.User.Svc.Create(ctx, admin)
	if err != nil {
		return err
	}

	return nil
}
