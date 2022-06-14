package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	log.Println(":: Hello go-envload ::")
	e := godotenv.Load()

	if e != nil {
		log.Fatal("Error loading .env file (main)")
	}

	envMsg := os.Getenv("MSG")
	log.Println(envMsg)
}