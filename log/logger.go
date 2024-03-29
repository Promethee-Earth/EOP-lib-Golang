package log

import (
	"fmt"
	"net"
	"os"
	"runtime"
	"strings"
	"time"
)

type logger struct {
	format         string
	serviceName    string
	serverIP       string
	logEverything  bool
	counterRequest uint64
	counterError   uint64
}

// Property accessor
func (l *logger) GetRequestCounter() uint64 { return l.counterRequest }
func (l *logger) GetErrorCounter() uint64   { return l.counterError }

// NewLogger constructs a new instance of logger
func NewLogger(serviceName, format string, logEverything bool) logger {
	var log = logger{
		format:        format + "\n",
		serviceName:   serviceName,
		logEverything: logEverything}

	log.serverIP, _ = os.Hostname()
	if log.serverIP == "" {
		var addrs, _ = net.LookupIP(log.serverIP)
		for _, address := range addrs {
			if ipv4 := address.To4(); ipv4 != nil {
				log.serverIP = ipv4.String()
				break
			}
		}
	}

	return log
}

// Info logs an informative message.
func (l *logger) Info(values ...any) {
	fmt.Printf(l.format, time.Now().Unix(), l.serviceName, l.serverIP, "", "INFO", "",
		strings.TrimSpace(fmt.Sprintln(values...)))
}

// Fatal logs a message then quit the program!
func (l *logger) Fatal(msg string) {
	fmt.Printf(l.format, time.Now().Unix(), l.serviceName, l.serverIP, "", "FATAL", "", msg)
	os.Exit(1)
}

// NewRequest constructs a new instance of request
func (l *logger) NewRequest(traceID string, payload any) request {
	l.counterRequest++

	var pc, _, _, _ = runtime.Caller(1)
	var function = runtime.FuncForPC(pc).Name()
	function = function[strings.LastIndex(function, ".")+1:]

	if l.logEverything {
		fmt.Printf(l.format,
			time.Now().Unix(), l.serviceName, l.serverIP, traceID, "REQUEST", function,
			strings.TrimSpace(strings.ReplaceAll(fmt.Sprintln(payload), `trace_id:"`+traceID+`"`, "")))
	}

	return request{
		endpoint: function,
		traceID:  traceID,
		logger:   l}
}
