package main

import (
	"log"

	"github.com/snowztech/inkssg"
)

func main() {
	if err := inkssg.Build("."); err != nil {
		log.Fatal(err)
	}
}
