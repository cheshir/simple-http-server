package main

import (
	"fmt"
	"net/url"
	"strings"
	"time"
)

const (
	colorYellow = "\x1b[1;33m"
	colorGreen  = "\x1b[1;32m"
	resetStyles = "\x1b[0m"
)

type RequestLogger struct {
	fieldColor  string
	valueColour string
}

func (l *RequestLogger) Log(URL *url.URL, method string, payload []byte) {
	line := &strings.Builder{}
	l.writeLine(line, "Time", time.Now().String())
	l.writeLine(line, "URL", URL.String())
	l.writeLine(line, "Method", method)
	l.writeLine(line, "Body", string(payload))

	fmt.Println(line.String())
}

func (l *RequestLogger) writeLine(builder *strings.Builder, field, value string) {
	builder.WriteString(l.fieldColor)
	builder.WriteString(field)
	builder.WriteString(":\t")
	builder.WriteString(l.valueColour)
	builder.WriteString(value)
	builder.WriteString(resetStyles)
	builder.WriteByte('\n')
}
