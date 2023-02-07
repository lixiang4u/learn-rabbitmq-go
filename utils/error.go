package utils

import (
	"fmt"
	"log"
)

func PrintMqFailOnError(err error, msg string) {
	if err != nil {
		log.Println(fmt.Sprintf("%s: %s", msg, err))
	}
}
