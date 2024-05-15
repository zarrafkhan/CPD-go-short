package main

import (
	"log"
)

func main() {

}

// handle error
func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
