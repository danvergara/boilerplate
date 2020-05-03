package logger

import (
	"log"
	"os"
)

var (
	// Info is a logger with info format
	Info = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	// Error is a logger that print error through the Stderr
	Error = log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
)
