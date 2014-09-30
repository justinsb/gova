package services

import (
	"github.com/justinsb/gova/assert"
	"github.com/justinsb/gova/collections"
	"github.com/justinsb/gova/errors"
	"github.com/justinsb/gova/log"

	"sync"
)

type ServiceMap struct {
	factory     fnServiceFactory
	keyFunction fnKeyFunction

	services map[string]*ServiceSlot
	changes  chan interface{}

	mutex sync.Mutex
}

type ServiceSlot struct {
	service interface{}
	channel chan<- interface{}
}

type fnServiceFactory func(config interface{}, channel <-chan interface{}) interface{}
type fnKeyFunction func(config interface{}) string

func NewServiceMap(keyFunction fnKeyFunction, factory fnServiceFactory) *ServiceMap {
	self := &ServiceMap{}
	self.factory = factory
	self.keyFunction = keyFunction

	self.services = make(map[string]*ServiceSlot)
	self.changes = make(chan interface{})

	go func() {
		//		foundKeys := make(map[string]bool)

		for config := range self.changes {
			slot := self.getSlot(config)
			if slot == nil {
				continue
			}

			slot.channel <- config
		}
	}()

	return self
}

func (self *ServiceMap) Channel() chan<- interface{} {
	return self.changes
}

func (self *ServiceMap) getSlot(config interface{}) *ServiceSlot {
	key := self.keyFunction(config)

	self.mutex.Lock()
	defer self.mutex.Unlock()

	slot := self.services[key]
	if slot == nil {
		channel := make(chan interface{})
		service := self.factory(config, channel)
		if service == nil {
			log.Warn("Could not create service for key: %v", key)
			return nil
		}

		slot = &ServiceSlot{}
		slot.service = service
		slot.channel = channel

		self.services[key] = slot
	}

	return slot

}

func (self *ServiceMap) keys() []string {
	self.mutex.Lock()
	defer self.mutex.Unlock()

	keys := make([]string, 0, len(self.services))
	for key, _ := range self.services {
		keys = append(keys, key)
	}
	return keys
}

//func (self *ServiceMap) Update(configs ServiceConfigs) errors.ErrorList {
//	e := errors.NewErrorList()
//
//	foundKeys := make(map[string]bool)
//
//	for _, config := range configs {
//		key := config.Key()
//		foundKeys[key] = true
//
//		service := self.services[key]
//		if service == nil {
//			service = self.factory(config)
//			if service == nil {
//				log.Warn("Could not create service for key: %v", key)
//				e.Add(fmt.Errorf("Could not create service for key: %v", key))
//			} else {
//				self.services[key] = service
//
//				e.AddAll(service.Start())
//			}
//		} else {
//			e.AddAll(service.Update(config))
//		}
//	}
//
//	for key, service := range self.services {
//		if !foundKeys[key] {
//			errs := service.Stop()
//
//			e.AddAll(errs)
//			if !errs.IsEmpty() {
//				log.Warn("Error stopping service: %v", service, errs)
//			} else {
//				delete(self.services, key)
//			}
//		}
//	}
//
//	return e
//}

func (self *ServiceMap) Snapshot() []interface{} {
	self.mutex.Lock()
	defer self.mutex.Unlock()

	services := make([]interface{}, 0, len(self.services))
	for _, slot := range self.services {
		services = append(services, slot.service)
	}
	return services
}

func (self *ServiceMap) Stop() errors.ErrorList {
	log.Error("ServiceMap::Stop not implemented")
	return errors.NoErrors()
}

func (self *ServiceMap) remove(tombstone interface{}) {
	key := self.keyFunction(tombstone)

	slot := self.getSlot(tombstone)
	if slot == nil {
		log.Warn("Tried to remove service, but service not found: %v", key)
		return
	}

	self.mutex.Lock()
	defer self.mutex.Unlock()

	slot.channel <- tombstone
	close(slot.channel)

	delete(self.services, key)
}

func (self *ServiceMap) ReplaceConfigs(configs collections.Sequence, fnTombstone func(string) interface{}) {
	keys := make(map[string]bool)
	for _, key := range self.keys() {
		keys[key] = true
	}

	for iterator := configs.Iterator(); iterator.HasNext(); {
		config := iterator.Next()
		key := self.keyFunction(config)

		keys[key] = false

		slot := self.getSlot(config)
		if slot == nil {
			continue
		}

		slot.channel <- config
	}

	for key, v := range keys {
		if !v {
			continue
		}

		tombstone := fnTombstone(key)

		assert.That(key == self.keyFunction(tombstone))

		self.remove(tombstone)
	}
}
