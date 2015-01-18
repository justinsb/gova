package log

import "strings"

type Priority int

const (
	PriorityDebug Priority = 10000
	PriorityInfo  Priority = 20000
	PriorityWarn  Priority = 30000
	PriorityError Priority = 40000
	PriorityFatal Priority = 50000
)

func (priority Priority) String() string {
	switch priority {
	case PriorityDebug:
		return "DEBUG"
	case PriorityInfo:
		return "INFO"
	case PriorityWarn:
		return "WARN"
	case PriorityError:
		return "ERROR"
	case PriorityFatal:
		return "FATAL"

	default:
		return "UNKNOWN"
	}
}

func ParsePriority(s string) (Priority, bool) {
	s = strings.ToLower(s)
	if s == "debug" {
		return PriorityDebug, true
	}
	if s == "info" {
		return PriorityInfo, true
	}
	if s == "warn" {
		return PriorityWarn, true
	}
	if s == "error" {
		return PriorityError, true
	}
	if s == "fatal" {
		return PriorityFatal, true
	}
	return PriorityInfo, false
}

