package collections

import (
	"io"
	"sync"

	"container/list"

	"github.com/justinsb/slf4g/log"
)

type Pooled interface {
	io.Closer

	ReturnToPool()
	Value() interface{}
}

type PoolEntry struct {
	pool     *Pool
	borrowed interface{}
}

func (self *PoolEntry) Close() error {
	if self.borrowed != nil {
		self.ReturnToPool()
	}
	return nil
}

func (self *PoolEntry) ReturnToPool() {
	self.pool.Return(self)
	self.borrowed = nil
}

func (self *PoolEntry) Value() interface{} {
	return self.borrowed
}

type Pool struct {
	allocated map[interface{}]bool

	mutex sync.Mutex

	available *list.List
}

func (self *Pool) Init() {
	self.allocated = make(map[interface{}]bool)
	self.available = list.New()
}

func NewPool() *Pool {
	self := &Pool{}

	self.Init()

	return self
}

func (self *Pool) Add(element interface{}) {
	self.mutex.Lock()
	defer self.mutex.Unlock()

	if self.allocated[element] {
		panic("Elemenet is already allocated from the pool")
	}

	self.available.PushBack(element)
}

func (self *Pool) Borrow() Pooled {
	self.mutex.Lock()
	defer self.mutex.Unlock()

	head := self.available.Front()
	if head == nil {
		log.Warn("No elements left in pool")
		return nil
	}

	self.available.Remove(head)

	element := head.Value

	if self.allocated[element] {
		panic("Attempt to do double-allocation")
	}

	self.allocated[element] = true

	return &PoolEntry{pool: self, borrowed: element}
}

func (self *Pool) Return(pooled Pooled) {
	self.mutex.Lock()
	defer self.mutex.Unlock()

	poolEntry := pooled.(*PoolEntry)
	borrowed := poolEntry.borrowed

	if !self.allocated[borrowed] {
		panic("Attempt to return item that was not allocated")
	}

	// Not clear whether we should delete or set to false...
	//	delete(self.allocated, borrowed)
	self.allocated[borrowed] = false

	self.available.PushBack(borrowed)
}
