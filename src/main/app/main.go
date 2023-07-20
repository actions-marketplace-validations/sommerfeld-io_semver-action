package main

import (
	"log"

	"github.com/sommerfeld-io/semver/cmd"
)

func init() {
	log.SetPrefix("[semver] ")
}

func main() {
	err := cmd.Execute()

	if err != nil {
		log.Fatal(err)
	}
}
