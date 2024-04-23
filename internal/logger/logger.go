package logger

import "log"

// LogInfo logs informational messages to the standard log.
func LogInfo(messages ...interface{}) {
	log.Println(messages...)
}

// LogError logs error messages to the standard log.
func LogError(messages ...interface{}) {
	log.Println(messages...)
}
