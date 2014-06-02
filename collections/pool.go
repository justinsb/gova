package collections

import (
	"io"
	"sync"

	"container/list"

	"github.com/justinsb/gova/log"
)

type HasOwner interface {
	GetOwner() string
}

type HasKey interface {
	GetKey() string
}

type Pool interface {
	Borrow(owner string) Pooled
	Return(pooled Pooled)
}

type Pooled interface {
	io.Closer

	ReturnToPool()
	Value() interface{}
}

type PoolEntry struct {
	pool     Pool
	borrowed interface{}
}

func NewPooled(pool Pool, borrowed interface{}) Pooled {
	pooled := &PoolEntry{}
	pooled.pool = pool
	pooled.borrowed = borrowed
	return pooled
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

type InMemoryPool struct {
	allocated map[interface{}]bool

	mutex sync.Mutex

	available *list.List
}

func (self *InMemoryPool) init() {
	self.allocated = make(map[interface{}]bool)
	self.available = list.New()
}

func NewPool() *InMemoryPool {
	self := &InMemoryPool{}

	self.init()

	return self
}

func (self *InMemoryPool) Add(element interface{}) {
	self.mutex.Lock()
	defer self.mutex.Unlock()

	if self.allocated == nil {
		self.init()
	}

	if self.allocated[element] {
		panic("Elemenet is already allocated from the pool")
	}

	self.available.PushBack(element)
}

func (self *InMemoryPool) Borrow(owner string) Pooled {
	self.mutex.Lock()
	defer self.mutex.Unlock()

	if self.allocated == nil {
		self.init()
	}

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

	return NewPooled(self, element)
}

func (self *InMemoryPool) Return(pooled Pooled) {
	self.mutex.Lock()
	defer self.mutex.Unlock()

	if self.allocated == nil {
		self.init()
	}

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
