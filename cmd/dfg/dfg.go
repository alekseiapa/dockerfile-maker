package main

import (
	"github.com/alekseiapa/dockerfile-maker/cmd/dfg/app"
	"os"
)

func main() {
	if err := app.Run(); err != nil {
		os.Exit(1)
	}
}
