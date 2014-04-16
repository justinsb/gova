package log

import (
	"encoding/json"
	"fmt"
)

type Logger struct {
	appenders []Appender
}

func (self *Logger) log(priority Priority, msg string, params ...interface{}) {
	e := &Event{}

	n := len(params)
	if n != 0 {
		last := params[n-1]
		if err, ok := last.(error); ok {
			e.Error = err
		}
	}
	e.Params = params
	e.Message = msg
	e.Priority = priority

	for _, appender := range self.appenders {
		appender.doAppend(e)
	}
}

func (self *Logger) Fatal(msg string, params ...interface{}) {
	self.log(PriorityFatal, msg, params...)
}

func (self *Logger) Error(msg string, params ...interface{}) {
	self.log(PriorityError, msg, params...)
}

func (self *Logger) Warn(msg string, params ...interface{}) {
	self.log(PriorityWarn, msg, params...)
}

func (self *Logger) Info(msg string, params ...interface{}) {
	self.log(PriorityInfo, msg, params...)
}

func (self *Logger) Debug(msg string, params ...interface{}) {
	self.log(PriorityDebug, msg, params...)
}

func (self *Logger) AddAppender(logger Appender) {
	// TODO: mutex
	self.appenders = append(self.appenders, logger)
}

func AsJson(o interface{}) string {
	if o == nil {
		return "nil"
	}
	bytes, err := json.Marshal(o)
	if err != nil {
		return fmt.Sprintf("[Error: %v]", err)
	}

	return string(bytes)
}
