package seeds

import (
	"app/app/helper"
	"app/app/modules"
	userdto "app/app/modules/user/dto"
	"context"
	"fmt"
)

func adminSeed(ctx context.Context, mod *modules.Modules) error {

	number := helper.RandomNumber()
	avatar := fmt.Sprintf("https://avatars.githubusercontent.com/u/%v?v=4", number)

	admin := userdto.CreateUserRequest{
		FirstName: "Admin",
		LastName:  "Admin",
		Email:     "admin@mail.com",
		Password:  "123456789",
		Avatar:    avatar,
	}

	_, _, err := mod.User.Svc.Create(ctx, admin)
	if err != nil {
		return err
	}

	return nil
}
