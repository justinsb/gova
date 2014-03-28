package log

type Appender interface {
	doAppend(e *Event)
}
