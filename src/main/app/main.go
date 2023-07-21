package main

import (
	"log"

	"github.com/sommerfeld-io/semver/commands"
)

func init() {
	log.SetPrefix("[semver] ")
}

func main() {
	err := commands.Execute()

	if err != nil {
		log.Fatal(err)
	}
}
