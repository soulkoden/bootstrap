package bootstrap

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func MustLoadEnv() {
	if err := godotenv.Load(".env"); err != nil && !os.IsNotExist(err) {
		cannotLoadEnvFile(".env", err)
	}

	if err := godotenv.Overload(".env.local"); err != nil && !os.IsNotExist(err) {
		cannotLoadEnvFile(".env.local", err)
	}
}

func cannotLoadEnvFile(filename string, err error) {
	log.Fatalf("cannot load env file %s: %v\n", filename, err)
}
