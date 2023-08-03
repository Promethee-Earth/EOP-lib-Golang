package log

import (
	"fmt"
	"net"
	"os"
	"runtime"
	"strings"
	"time"
)

type Logger struct {
	format         string
	host           string
	logEverything  bool
	counterRequest uint64
	counterError   uint64
}

// Property accessor
func (l *Logger) GetRequestCounter() uint64 { return l.counterRequest }
func (l *Logger) GetErrorCounter() uint64   { return l.counterError }

// NewLogger constructs a new instance of Logger
func NewLogger(format string, logEverything bool) Logger {
	var log = Logger{
		format:        format + "\n",
		logEverything: logEverything}

	log.host, _ = os.Hostname()
	if log.host == "" {
		var addrs, _ = net.LookupIP(log.host)
		for _, address := range addrs {
			if ipv4 := address.To4(); ipv4 != nil {
				log.host = ipv4.String()
				break
			}
		}
	}

	return log
}

// Info logs an informative message.
func (l *Logger) Info(values ...any) {
	fmt.Printf(l.format, time.Now().Unix(), l.host, "", "", "INFO",
		strings.TrimSpace(fmt.Sprintln(values...)))
}

// Fatal logs a message then quit the program!
func (l *Logger) Fatal(msg string) {
	fmt.Printf(l.format, time.Now().Unix(), l.host, "", "", "FATAL", msg)
	os.Exit(1)
}

// NewRequest constructs a new instance of request
func (l *Logger) NewRequest(traceID, payload string) request {
	l.counterRequest++

	var pc, _, _, _ = runtime.Caller(1)
	var function = runtime.FuncForPC(pc).Name()
	function = function[strings.LastIndex(function, ".")+1:]

	if l.logEverything {
		fmt.Printf(l.format, time.Now().Unix(), l.host, traceID, function, "REQUEST",
			strings.TrimSpace(strings.ReplaceAll(payload, `trace_id:"`+traceID+`"`, "")))
	}

	return request{
		endpoint: function,
		traceID:  traceID,
		logger:   l}
}
