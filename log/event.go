package log

type Event struct {
	Priority Priority
	Message  string
	Error    error
	Params   []interface{}
}
