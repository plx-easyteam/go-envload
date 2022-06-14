package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnvMsg(msg string) string {
	e := godotenv.Load()

	if e != nil {
		log.Fatal("Error loading .env file (loader)")
	}

	if msg == ""{
		return os.Getenv("HELLO")
	}

	return os.Getenv(msg)
}