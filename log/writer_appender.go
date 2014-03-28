package log

import (
	"io"
)

type WriterAppender struct {
	writer io.Writer
	layout Layout
}

func (self *WriterAppender) doAppend(e *Event) {
	s := self.layout.doLayout(e)
	s = s + "\n"
	self.writer.Write([]byte(s))
}

func NewWriterAppender(writer io.Writer, layout Layout) Appender {
	s := &WriterAppender{}
	s.writer = writer
	s.layout = layout
	return s
}
