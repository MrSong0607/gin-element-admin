package utils

import (
	"fmt"
	"log"
	"os"
)

const TraceID ContextKey = "TraceID"

var (
	logger = log.New(os.Stderr, "[API]", log.LstdFlags|log.Llongfile)
)

type ContextKey string

func LogInfo(v ...any) {
	logger.Output(2, fmt.Sprintln(v...))
}
