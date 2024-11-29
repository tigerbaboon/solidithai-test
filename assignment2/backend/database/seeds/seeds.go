package seeds

import (
	"app/app/modules"
	"context"
)

func Seeds(mod *modules.Modules, ctx context.Context) error {

	if err := adminSeed(ctx, mod); err != nil {
		return err
	}

	return nil
}
