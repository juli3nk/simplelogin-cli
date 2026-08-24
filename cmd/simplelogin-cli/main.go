package main

import (
	"log"

	"github.com/juli3nk/simplelogin-cli/command"
)

func main() {
	if err := command.NewSimpleLoginCommand().Execute(); err != nil {
		log.Fatal(err)
	}
}
