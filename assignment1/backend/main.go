package main

import (
	"app/internal/cmd"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

func main() {

	godotenv.Load()

	rootCmd := &cobra.Command{
		Use: "app",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Usage()
		},
	}

	rootCmd.AddCommand(cmd.Http())
	rootCmd.AddCommand(cmd.Migrate())

	rootCmd.Execute()
}
