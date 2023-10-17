// Package main is naraku command entrypoint.
package main

import (
	"os"

	"github.com/go-spectest/naraku/cmd"
)

// osExit is wrapper for  os.Exit(). It's for unit test.
var osExit = os.Exit //nolint

func main() {
	osExit(cmd.Execute())
}
