package collections

type ArrayList struct {
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
