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

	if os.Getenv("SIGNING_KEY") == "" {
		logrus.Fatal("Missing JWT signing key as `SIGNING_KEY`")
		os.Exit(1)
	}

	logrus.Info("Environment variables look alright so far...")
}
