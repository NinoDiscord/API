package utils

import (
	"github.com/sirupsen/logrus"
	"os"
)

// ValidateEnv makes ensures all environment variables are correct
func ValidateEnv() {
	if os.Getenv("GO_ENV") == "" {
		logrus.Fatal("Environment variable `GO_ENV` was not defined.")
	}

	if os.Getenv("MASTER_KEY") == "" {
		logrus.Fatal("Missing master key as `MASTER_KEY` environment variable.")
		os.Exit(1)
	}

	logrus.Info("Environment variables look alright so far...")
}
