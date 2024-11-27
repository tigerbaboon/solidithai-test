package cmd

import (
	"app/app/modules"
	"app/routes"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func Http() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "http",
		Short: "Start the HTTP server",
		Run: func(cmd *cobra.Command, args []string) {
			r := gin.Default()
			r.GET("/ping", func(c *gin.Context) {
				c.JSON(200, gin.H{
					"message": "pong",
				})
			})

			mod := modules.Get()

			routes.Router(r, mod)

			port := viper.GetString("PORT")
			r.Run(fmt.Sprintf(":%s", port))
		},
	}

	return cmd
}
