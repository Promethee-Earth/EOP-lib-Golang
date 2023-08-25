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
func (r Request) Warning(values ...any) {
	fmt.Printf(r.logger.format,
		time.Now().Unix(), r.logger.host, r.traceID, r.endpoint, "WARNING",
		strings.TrimSpace(fmt.Sprintln(values...)))
}

// Error logs a gRPC error message then returns it.
func (r Request) Error(err error) error {
	r.logger.counterError++
	fmt.Printf(r.logger.format,
		time.Now().Unix(), r.logger.host, r.traceID, r.endpoint, "ERROR", err)
	return err
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
