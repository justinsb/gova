package assert

import (
	"bytes"
	"fmt"
	"path"
	"reflect"
	"runtime"

	"github.com/justinsb/slf4g/log"
)

func That(predicate bool) {
	// TODO: Turn off assertions?
	if predicate {
		return
	}

	Fail("Expected condition was false")
}

func Fail(message string) {
	log.Error("Failed assertion: %v", message)

	var buffer bytes.Buffer

	callers := make([]uintptr, 20)
	skip := 1
	n := runtime.Callers(skip, callers)
	callers = callers[:n]

	buffer.WriteString("Assertion failed")
	buffer.WriteString("\n")
	for _, pc := range callers {
		f := runtime.FuncForPC(pc)
		//if !strings.Contains(f.Name(), "/slf4g/")
		{
			pathname, lineno := f.FileLine(pc)
			filename := path.Base(pathname)

			s := fmt.Sprintf("\n    at %s (%s:%d)", f.Name(), filename, lineno)
			buffer.WriteString(s)
		}
	}

	log.Error(buffer.String())

	// TODO: Should we panic?
	panic("Assertion failed: " + buffer.String())
}

func Equal(left interface{}, right interface{}) {
	// TODO: Turn off assertions?
	if reflect.DeepEqual(left, right) {
		return
	}

	message := fmt.Sprintf("Not equal: %v != %v", left, right)
	Fail(message)
}
