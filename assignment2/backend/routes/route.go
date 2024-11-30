package routes

import (
	"app/app/modules"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

func Router(app *gin.Engine, mod *modules.Modules) {
	app.Use(otelgin.Middleware(viper.GetString("APP_NAME")))

	app.Use(cors.New(cors.Config{
		AllowAllOrigins:        true,
		AllowMethods:           []string{"*"},
		AllowHeaders:           []string{"*"},
		AllowCredentials:       true,
		AllowWildcard:          true,
		AllowBrowserExtensions: true,
		AllowWebSockets:        true,
	}))

	api(app.Group("/api"), mod)
}
