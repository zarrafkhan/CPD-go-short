package db

import (
	"log"
)

func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
