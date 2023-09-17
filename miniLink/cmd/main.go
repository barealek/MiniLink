package main

import (
	"fmt"
	"os"

	"github.com/barealek/minilink/pkg/shutdown"
)

func main() {
	// Prepare the exit code and defer the exit
	var exitCode int

	defer func() {
		os.Exit(exitCode)
	}()

	// Run the application
	cleanup, err := run()
	defer cleanup()

	// Handle any startup error
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		exitCode = 1
		return
	}

	// Wait for the application to exit and exit gracefully
	shutdown.Shutdown()
}

func run() (func(), error) {}
