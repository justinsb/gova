package lists

import (
	"github.com/justinsb/gova/collections"
)

func Of(items ...interface{}) collections.Sequence {
	ret := collections.NewArrayList()

	for _, i := range items {
		ret.Add(i)
	}

	return ret
}
