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

// NewRequest constructs a new instance of Request
func (l *Logger) NewRequest(payload string) Request {
	l.counterRequest++
	var request = Request{logger: l}

	var pc, _, _, _ = runtime.Caller(1)
	request.endpoint = runtime.FuncForPC(pc).Name()
	request.endpoint = request.endpoint[strings.LastIndex(request.endpoint, ".")+1:]

	var index = strings.Index(payload, "trace_id:")
	if index >= 0 {
		request.traceID = payload[index+10:]
		request.traceID = request.traceID[:strings.Index(request.traceID, `"`)]
	}

	if l.logEverything {
		fmt.Printf(
			l.format, time.Now().Unix(), l.host, request.traceID, request.endpoint, "REQUEST",
			strings.TrimSpace(strings.ReplaceAll(payload, `trace_id:"`+request.traceID+`"`, "")))
	}

	return request
}
