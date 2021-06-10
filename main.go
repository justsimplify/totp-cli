package main

import (
	"log"

	"github.com/justsimplify/totp-cli/cmd"
)

func main() {
	err := cmd.Run()
	if err != nil {
		log.Fatal(err.Error())
	}
}
