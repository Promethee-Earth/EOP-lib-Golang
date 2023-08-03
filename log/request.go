package log

import (
	"fmt"
	"strings"
	"time"
)

type request struct {
	logger   *Logger
	endpoint string
	traceID  string
}

// Response logs a response payload.
func (r request) Response(payload any) {
	if r.logger.logEverything {
		fmt.Printf(r.logger.format,
			time.Now().Unix(), r.logger.host, r.traceID, r.endpoint, "RESPONSE", payload)
	}
}

// Info logs an informative message for debugging purposes.
func (r request) Info(msg string) {
	if r.logger.logEverything {
		fmt.Printf(r.logger.format,
			time.Now().Unix(), r.logger.host, r.traceID, r.endpoint, "INFO", msg)
	}
}

// Warning logs a warning message.
func (r request) Warning(msg string) {
	fmt.Printf(r.logger.format,
		time.Now().Unix(), r.logger.host, r.traceID, r.endpoint, "WARNING", msg)
}

// Error logs an error message.
func (r request) Error(values ...any) {
	r.logger.counterError++
	fmt.Printf(r.logger.format,
		time.Now().Unix(), r.logger.host, r.traceID, r.endpoint, "ERROR",
		strings.TrimSpace(fmt.Sprintln(values...)))
}

// DontPanic recovers from a panic and catches the error
func (r request) DontPanic() {
	var err = recover()
	if err != nil {
		r.logger.counterError++
		fmt.Printf(r.logger.format,
			time.Now().Unix(), r.logger.host, r.traceID, r.endpoint, "FATAL", err)
	}
}
