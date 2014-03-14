package services

import (
	//	"github.com/justinsb/gova/collections"
	"github.com/justinsb/gova/errors"
	"github.com/justinsb/slf4g/log"

	"sync"
	//	"fmt"
	//	"io"
)

type ServiceMap struct {
	factory  fnServiceFactory
	services map[string]*ServiceSlot
	changes  chan ServiceConfig

	mutex sync.Mutex
}

type ServiceSlot struct {
	service interface{}
	channel chan<- ServiceConfig
}

type fnServiceFactory func(key string, channel <-chan ServiceConfig) interface{}

func NewServiceMap(factory fnServiceFactory) *ServiceMap {
	self := &ServiceMap{}
	self.factory = factory
	self.services = make(map[string]*ServiceSlot)
	self.changes = make(chan ServiceConfig)

	go func() {
		//		foundKeys := make(map[string]bool)

		for config := range self.changes {
			key := config.Key()
			slot := self.getSlot(key)
			if slot == nil {
				continue
			}

			slot.channel <- config
		}
	}()

	return self
}

func (self *ServiceMap) Channel() chan<- ServiceConfig {
	return self.changes
}

func (self *ServiceMap) getSlot(key string) *ServiceSlot {
	self.mutex.Lock()
	defer self.mutex.Unlock()

	slot := self.services[key]
	if slot == nil {
		channel := make(chan ServiceConfig)
		service := self.factory(key, channel)
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

func (self *ServiceMap) Stop() errors.ErrorList {
	log.Error("TODO: Implement!!!")
	return errors.NoErrors()
}

func (self *ServiceMap) remove(key string, tombstone ServiceConfig) {
	slot := self.getSlot(key)
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

func (self *ServiceMap) ReplaceConfigs(configs []ServiceConfig, fnTombstone func(string) ServiceConfig) {
	keys := make(map[string]bool)
	for _, key := range self.keys() {
		keys[key] = true
	}

	for _, config := range configs {
		key := config.Key()

		delete(keys, key)

		slot := self.getSlot(key)
		if slot == nil {
			continue
		}

		slot.channel <- config
	}

	for key, _ := range keys {
		tombstone := fnTombstone(key)

		self.remove(key, tombstone)
	}
}
