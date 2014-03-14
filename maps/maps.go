package maps

import (
	"reflect"

	"github.com/justinsb/gova/collections"
)

func Values(m interface{}) collections.Sequence {
	vMap := reflect.ValueOf(m)
	ret := collections.NewArrayList()

	keys := vMap.MapKeys()

	// TODO: ret.Reserve(len(keys))

	for _, vKey := range keys {
		vValue := vMap.MapIndex(vKey)

		ret.Add(vValue.Interface)
	}

	return ret
}
