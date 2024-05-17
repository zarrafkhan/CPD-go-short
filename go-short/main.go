package main

import (
	I "example/go-short/internals"
	"fmt"
)

func main() {
	fmt.Println("Start shorturl in pure golang")
	if I.Start_Server() < 0 {
		defer I.Disc(I.Client)
	}
}
