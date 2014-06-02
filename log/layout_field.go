package log

import (
	"bytes"
	"fmt"
	"path"
	"reflect"
	"runtime"
	"strings"
	"time"
)

type LayoutField interface {
	format(e *Event) string
}

type TimestampLayoutField struct {
}

func (self *TimestampLayoutField) format(e *Event) string {
	now := time.Now()

	return now.Format(time.StampMilli)
}

type PriorityLayoutField struct {
}

func (self *PriorityLayoutField) format(e *Event) string {
	return e.Priority.String()
}

type LiteralField struct {
	Literal string
}

func (self *LiteralField) format(e *Event) string {
	return self.Literal
}

type MessageField struct {
}

func (self *MessageField) format(e *Event) string {
	format := e.Message
	s := fmt.Sprintf(format, e.Params...)
	return s
}

type ErrorField struct {
}

func (self *ErrorField) format(e *Event) string {
	if e.Error == nil {
		return ""
	}

	var buffer bytes.Buffer

	callers := make([]uintptr, 20)
	n := runtime.Callers(3, callers) // starts in (*Logger).Log or similar
	callers = callers[:n]

	buffer.WriteString("\n")
	errorType := reflect.TypeOf(e.Error)
	buffer.WriteString(errorType.String() + ": ")
	buffer.WriteString(e.Error.Error())
	for _, pc := range callers {
		f := runtime.FuncForPC(pc)
		if !strings.Contains(f.Name(), "/gova/log/") {
			pathname, lineno := f.FileLine(pc)
			filename := path.Base(pathname)

			s := fmt.Sprintf("\n    at %s (%s:%d)", f.Name(), filename, lineno)
			buffer.WriteString(s)
		}
	}

	return buffer.String()
}
