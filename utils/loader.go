package utils

import (
	"flag"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

const PROJECT_NAME = "go-envload"

func GetEnvValue(key string) string {

	if flag.Lookup("test.v") == nil {
		// When "go run"
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatalln("Error loading .env file (loader) ", err)
		}
		return os.Getenv(key)
	} else {
		// When "go test"
		path, _ := os.Getwd()
		err := godotenv.Load(strings.Split(path, PROJECT_NAME)[0] + PROJECT_NAME + "\\.env.test")
		if err != nil {
			log.Fatalln("Error loading .env.test file (loader) ", err)
		}

		return os.Getenv(key)
	}
}
