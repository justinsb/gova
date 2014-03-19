package assert

import (
	"bytes"
	"fmt"
	"path"
	"runtime"

	"github.com/justinsb/slf4g/log"
)

func That(predicate bool) {
	// TODO: Turn off assertions?
	if predicate {
		return
	}

	log.Warn("Failed assertion")

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
