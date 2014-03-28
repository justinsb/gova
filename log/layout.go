package log

type Layout interface {
	doLayout(e *Event) string
}
