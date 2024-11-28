package cmd

import (
	"app/app/modules"
	"app/config/log"
	"app/database/migrations"
	"app/database/seeds"
	"context"
	"os"
	"time"

	"github.com/spf13/cobra"
)

func Migrate() *cobra.Command {
	cmd := &cobra.Command{
		Use: "db",
	}

	cmd.AddCommand(migrateUp())
	cmd.AddCommand(migrateDown())
	cmd.AddCommand(migrateRefresh())
	cmd.AddCommand(migrateSeed())

	return cmd
}

func migrateUp() *cobra.Command {
	cmd := &cobra.Command{
		Use: "up",
		Run: func(cmd *cobra.Command, args []string) {
			log.Info("Executing model up...")
			db := modules.Get().DB
			for _, ent := range migrations.Entities() {
				if _, err := db.NewCreateTable().Model(ent).Exec(context.Background()); err != nil {
					log.Error("%s", err)
					os.Exit(1)
					return
				}
			}
			log.Info("model up success.")
		},
	}
	return cmd
}

func migrateDown() *cobra.Command {
	cmd := &cobra.Command{
		Use: "down",
		Run: func(cmd *cobra.Command, args []string) {
			log.Info("Executing model down...")
			db := modules.Get().DB
			for _, ent := range migrations.Entities() {
				if _, err := db.NewDropTable().Model(ent).Exec(context.Background()); err != nil {
					log.Error("%s", err)
					os.Exit(1)
					return
				}
			}
			log.Info("model down success.")

		},
	}
	return cmd
}

func migrateRefresh() *cobra.Command {
	cmd := &cobra.Command{
		Use: "refresh",
		Run: func(cmd *cobra.Command, args []string) {
			log.Info("Executing model refresh...")

			db := modules.Get().DB
			for _, ent := range migrations.Entities() {
				if _, err := db.NewDropTable().Model(ent).Exec(context.Background()); err != nil {
					log.Error("%s", err)
					os.Exit(1)
					return
				}
			}

			for _, ent := range migrations.Entities() {
				if _, err := db.NewCreateTable().Model(ent).Exec(context.Background()); err != nil {
					log.Error("%s", err)
					os.Exit(1)
					return
				}
			}

			log.Info("model refresh success.")

		},
	}
	return cmd
}

func migrateSeed() *cobra.Command {
	cmd := &cobra.Command{
		Use: "seed",
		Run: func(cmd *cobra.Command, args []string) {
			log.Info("Executing model seed...")

			mod := modules.Get()
			ctx := context.Background()

			startTime := time.Now()
			if err := seeds.Seeds(mod, ctx); err != nil {
				log.Error("%s", err)
				os.Exit(1)
				return
			}
			endTime := time.Now()
			duration := endTime.Sub(startTime)

			log.Info("model seed success.")
			log.Info("Model seed success. Duration: %s", duration)

		},
	}
	return cmd
}
