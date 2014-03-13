package collections

type Iterator interface {
	Next() interface{}
	HasNext() bool
}

type Iterable interface {
	Iterator() Iterator
}

type Collection interface {
	Iterable
	Size() int
}

type Sequence interface {
	Collection
	At(index int) interface{}
}

type MutableCollection interface {
	Collection
	Add(item interface{})
}

type SequenceIterator struct {
	seq  Sequence
	next int
}

func NewSequenceIterator(seq Sequence) SequenceIterator {
	return SequenceIterator{seq: seq, next: 0}
}

func (self *SequenceIterator) Next() interface{} {
	i := self.next
	self.next = i + 1
	return self.seq.At(i)
}

func (self *SequenceIterator) HasNext() bool {
	i := self.next
	return i < self.seq.Size()
}

type SliceSequence struct {
	items []interface{}
}

func NewSliceSequence(items []interface{}) SliceSequence {
	return SliceSequence{items: items}
}

func (self *SliceSequence) Iterator() Iterator {
	it := NewSequenceIterator(self)
	return &it
}

func (self *SliceSequence) Size() int {
	return len(self.items)
}

func (self *SliceSequence) At(index int) interface{} {
	return self.items[index]
}

func AsSequence(items ...interface{}) Sequence {
	s := NewSliceSequence(items)
	return &s
}
