package global

import (
	"github.com/go4digital/booknow-api/logger"
	"github.com/joho/godotenv"
)

func LoadEnvFile() {
	err := godotenv.Load(".env")
	if err != nil {
		logger.Error(err)
	}
}
