package calc

import (
	"go-envload/utils"
	"log"
)

func Add(a int, b int) int {
	log.Print(WhoAmI())
	return a + b
}

func WhoAmI() string {
	return utils.GetEnvValue("CALC_ADD")
}