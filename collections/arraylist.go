package collections

import "reflect"

type ArrayList struct {
	MutableCollection

	items []interface{}
}

func NewArrayList() *ArrayList {
	return &ArrayList{items: make([]interface{}, 0)}
}

func (self *ArrayList) Iterator() Iterator {
	it := NewSequenceIterator(self)
	return &it
}

func (self *ArrayList) Size() int {
	return len(self.items)
}

func (self *ArrayList) At(index int) interface{} {
	return self.items[index]
}

func (self *ArrayList) Add(item interface{}) {
	self.items = append(self.items, item)
}

func FromSlice(slice interface{}) Sequence {
	switch reflect.TypeOf(slice).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(slice)

		v := NewArrayList()

		for i := 0; i < s.Len(); i++ {
			v.Add(s.Index(i).Interface())
		}
		return v

	default:
		panic("Expected type slice in FromSlice")
	}
}
