package main

import (
	"go-envload/calc"
	"go-envload/utils"
	"log"
)

func main() {
	log.Println(":: Hello go-envload ::")


	calc.Add(2, 3)
	log.Println(HelloMsg())
}

func HelloMsg() string{
	return utils.GetEnvValue("HELLO")
}