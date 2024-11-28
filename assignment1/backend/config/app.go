package config

import (
	"app/config/log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func Init() {
	godotenv.Load()
	app()
	log.Init()
}

func app() {
	conf("APP_NAME", "go_app")
	conf("PORT", 8080)
	conf("JWT_SECRET", "secret")
	conf("JWT_DURATION", 24)

	conf("HTTP_JSON_NAMING", "camel_case")
}

func conf[T string | ~int | ~int32 | ~int64 | bool](key string, fallback T) T {
	viper.Set(key, fallback)
	if value, ok := os.LookupEnv(key); ok {
		viper.Set(key, value)
	}
	viper.UnmarshalKey(key, &fallback)
	return fallback
}
