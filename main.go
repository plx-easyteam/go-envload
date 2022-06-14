package main

import (
	"go-envload/utils"
	"log"
)

func main() {
	log.Println(":: Hello go-envload ::")


	log.Println(GetEnvMsg())
}

func GetEnvMsg() string{
	return utils.GetEnvMsg("")
}