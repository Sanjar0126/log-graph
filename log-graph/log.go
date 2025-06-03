package loggraph

import (
	"fmt"
	"time"
)

type LogLevel string

const (
	INFO  LogLevel = "INFO"
	DEBUG LogLevel = "DEBUG"
	ERROR LogLevel = "ERROR"
)

func Logger(level LogLevel, message string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("%s - %s - %s\n", timestamp, level, message)
}
