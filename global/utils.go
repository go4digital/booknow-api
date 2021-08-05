package global

import (
	"os"

	log "github.com/go4digital/booknow-api/logger"
	"github.com/joho/godotenv"
)

func Getenv(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Error(err)
	}
	return os.Getenv(key)
}
