package joiner

import (
	"bytes"
	"fmt"
	"github.com/justinsb/gova/collections"
	"reflect"
)

type Joiner struct {
	separator  string
	skipNulls  bool
	useForNull *string
}

func On(separator string) Joiner {
	self := Joiner{}
	self.separator = separator
	return self
}

func (self Joiner) SkipNulls() Joiner {
	var ret Joiner
	ret = self
	ret.skipNulls = true
	ret.useForNull = nil
	return ret
}

func (self Joiner) UseForNull(useForNull string) Joiner {
	var ret Joiner
	ret = self
	ret.skipNulls = true
	ret.useForNull = &useForNull
	return ret
}

func (self Joiner) joinIterable(items collections.Iterable) string {
	var buffer bytes.Buffer

	n := 0
	for iterator := items.Iterator(); iterator.HasNext(); {
		v := iterator.Next()

		var s string
		if v == nil {
			if self.skipNulls {
				continue
			} else if self.useForNull != nil {
				s = *self.useForNull
			} else {
				panic("Nil pointer detected")
			}
		} else {
			if stringer, ok := v.(fmt.Stringer); ok {
				s = stringer.String()
			} else {
				s = fmt.Sprint(v)
			}
		}

		if n != 0 {
			buffer.WriteString(self.separator)
		}

		buffer.WriteString(s)
		n = n + 1
	}

	return buffer.String()
}

func (self Joiner) Join(items interface{}) string {
	vItems := reflect.ValueOf(items)

	switch vItems.Kind() {
	case reflect.Slice, reflect.Array:
		n := vItems.Len()
		dest := make([]interface{}, n, n)
		for i := 0; i < n; i++ {
			dest[i] = vItems.Index(i).Interface()
		}
		iterable := collections.NewSliceSequence(dest)
		return self.joinIterable(iterable)

	default:
		panic("Unhandled type in Join")
	}

	return ""
}
