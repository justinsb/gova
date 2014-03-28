package log

import (
	"bytes"
)

type SimpleLayout struct {
	fields []LayoutField
}

func NewSimpleLayout() Layout {
	self := &SimpleLayout{}
	self.fields = []LayoutField{
		&TimestampLayoutField{},
		&LiteralField{Literal: " "},
		&PriorityLayoutField{},
		&LiteralField{Literal: " "},
		&MessageField{},
		&LiteralField{Literal: " "},
		&ErrorField{},
	}
	return self
}

func (self *SimpleLayout) doLayout(e *Event) string {
	var b bytes.Buffer
	for _, field := range self.fields {
		s := field.format(e)
		b.Write([]byte(s))
	}
	return b.String()
}
