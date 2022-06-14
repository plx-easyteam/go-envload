package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	log.Println(":: Hello go-envload ::")


	log.Println(GetEnvMsg())
}

func GetEnvMsg() string{
	e := godotenv.Load()
	
	if e != nil {
		log.Fatal("Error loading .env file (main)")
	}

	return os.Getenv("MSG")
}