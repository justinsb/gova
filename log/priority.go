package log

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
