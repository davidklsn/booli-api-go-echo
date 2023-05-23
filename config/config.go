package config

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadENV() error {
	goEnv := os.Getenv("GO_ENV")

	if goEnv != "production" {
		err := godotenv.Load()
		if err != nil {
			return err
		}
	}

	return nil
}
