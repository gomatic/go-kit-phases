package main

import (
	"os"
)

// main is the entrypoint for the command.
func main() { _ = run(os.Args) }

// run is separated from main to make it more straightforward to write unit tests for the command.
func run(args []string) error { return nil }
