package app

import (
	"github.com/alekseiapa/dockerfile-maker/cmd/dfg/cmd"
)

// Run is the entrypoint for the dfg command
func Run() error {
	cmd := cmd.NewDfgCommand()
	return cmd.Execute()
}
