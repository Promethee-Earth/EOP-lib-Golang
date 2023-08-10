package log

import (
	"fmt"
	"strings"
	"time"
)

type Request struct {
	logger   *Logger
	endpoint string
	traceID  string
}

// Response logs a response payload.
func (r Request) Response(payload string) {
	if r.logger.logEverything {
		fmt.Printf(r.logger.format,
			time.Now().Unix(), r.logger.host, r.traceID, r.endpoint, "RESPONSE", payload)
	}
}

// Info logs an informative message for debugging purposes.
func (r Request) Info(msg string) {
	if r.logger.logEverything {
		fmt.Printf(r.logger.format,
			time.Now().Unix(), r.logger.host, r.traceID, r.endpoint, "INFO", msg)
	}
}

// Warning logs a warning message.
func (r Request) Warning(msg string) {
	fmt.Printf(r.logger.format,
		time.Now().Unix(), r.logger.host, r.traceID, r.endpoint, "WARNING", msg)
}

// Error logs an error message.
func (r Request) Error(values ...any) {
	r.logger.counterError++
	fmt.Printf(r.logger.format,
		time.Now().Unix(), r.logger.host, r.traceID, r.endpoint, "ERROR",
		strings.TrimSpace(fmt.Sprintln(values...)))
}

// DontPanic recovers from a panic and catches the error
func (r Request) DontPanic() {
	var err = recover()
	if err != nil {
		r.logger.counterError++
		fmt.Printf(r.logger.format,
			time.Now().Unix(), r.logger.host, r.traceID, r.endpoint, "FATAL", err)
	}
}
