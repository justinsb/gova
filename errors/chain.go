package errors

import (
	"bytes"
)

type ChainedError struct {
	Message string
	Cause   error
}

func RootCause(err error) error {
	chained, ok := err.(ChainedError)
	if ok {
		return RootCause(chained.Cause)
	}
	return err
}

func (e ChainedError) Error() string {
	return e.Message + "\n" + e.Cause.Error()
}

func joinStrings(message ...string) string {
	var buffer bytes.Buffer

	for _, m := range message {
		buffer.WriteString(m)
	}
	return buffer.String()
}

func Chain(err error, message ...string) error {
	e := ChainedError{}
	e.Message = joinStrings(message...)
	e.Cause = err
	return e
}
