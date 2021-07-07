package cli

import (
	"fmt"
)

var version = "dev"

type versionCmd struct{}

// Version is a getter for the version variable.
func Version() string {
	return version
}

func init() {
	var cmd versionCmd
	if _, err := AddCommand(
		&cmd,
		"version",
		"Show app version",
		"Show the version of the application",
	); err != nil {
		// yep, panic. If this fails something is wrong with either the serveCmd struct, or the serveCmd.Execute function
		panic(err)
	}
}

// Execute implements the Commander interace from the jessevdk/go-flags package
// We don't (currnetly) care about positional arguments, so we use an `_ []string` to ignore them.
func (cmd versionCmd) Execute(_ []string) error {
	fmt.Printf("CLI version: %v\n", Version())
	return nil
}
