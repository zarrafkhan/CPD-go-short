package main

import (
	I "example/go-short/libraries"
)

func main() {
	if I.Start_Server() < 0 {
		defer I.Disc(I.Client)
	}
}
