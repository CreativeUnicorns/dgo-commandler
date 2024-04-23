// File: utils/logging.go
// Package: utils
// Description: This file contains the setup for structured logging.

package utils

import (
	"log/slog"
	"os"
)

// setupLogger configures and returns a new slog.Logger with structured logging.
func SetupLogger() *slog.Logger {
	// Creating a TextHandler to log in key=value text format to standard error.
	handler := slog.NewTextHandler(os.Stderr, nil)

	// Create a new logger with the configured handler.
	return slog.New(handler)
}

// Global Logger for use throughout the commandler package.
var Logger = SetupLogger()
