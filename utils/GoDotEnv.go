package utils

import (
	"os"

	"github.com/joho/godotenv"
)

func GoDotEnv(key string) string {
	env := make(chan string, 1)
	if os.Getenv("GO_ENV") != "PRODUCTION" {
		godotenv.Load(".env")
		env <- os.Getenv(key)
	} else {
		env <- os.Getenv(key)
	}
	return <-env
}
