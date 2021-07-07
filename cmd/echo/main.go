package main

import (
	"log"

	_ "github.com/erumble/go-echo/cmd/echo/commands"
	"github.com/erumble/go-echo/pkg/cli"
)

func main() {
	if err := cli.Execute(); err != nil {
		log.Fatal(err)
	}
}
