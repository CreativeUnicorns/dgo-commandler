// Package utils provides utility functions and configurations used by the commandler package,
// particularly for logging and other common tasks.
package utils

import (
	"log/slog"
	"os"
)

// SetupLogger configures and returns a new slog.Logger configured for structured logging.
// The logger outputs key=value pairs in text format to the standard error stream.
func SetupLogger() *slog.Logger {
	handler := slog.NewTextHandler(os.Stderr, nil)

	return slog.New(handler)
}

// Logger is a global slog.Logger instance used for logging throughout the commandler package.
// It is initialized with a handler that outputs logs in a structured text format.
var Logger = SetupLogger()
