package services

import (
	"github.com/justinsb/gova/collections"
	"github.com/justinsb/gova/errors"
	"github.com/justinsb/slf4g/log"

	"fmt"
	"io"
)

type Service interface {
	Start() errors.ErrorList
	Stop() errors.ErrorList
	Update(config ServiceConfig) errors.ErrorList

	ManageResource(owned io.Closer)
}

type ServiceMap struct {
	factory  fnServiceFactory
	services map[string]Service
}

// +gen
type ServiceConfig interface {
	Key() string
}

type fnServiceFactory func(config ServiceConfig) Service

func NewServiceMap(factory fnServiceFactory) *ServiceMap {
	self := &ServiceMap{}
	self.factory = factory
	self.services = make(map[string]Service)
	return self
}

func (self *ServiceMap) Update(configs ServiceConfigs) errors.ErrorList {
	e := errors.NewErrorList()

	foundKeys := make(map[string]bool)

	for _, config := range configs {
		key := config.Key()
		foundKeys[key] = true

		service := self.services[key]
		if service == nil {
			service = self.factory(config)
			if service == nil {
				log.Warn("Could not create service for key: %v", key)
				e.Add(fmt.Errorf("Could not create service for key: %v", key))
			} else {
				self.services[key] = service

				e.AddAll(service.Start())
			}
		} else {
			e.AddAll(service.Update(config))
		}
	}

	for key, service := range self.services {
		if !foundKeys[key] {
			errs := service.Stop()

			e.AddAll(errs)
			if !errs.IsEmpty() {
				log.Warn("Error stopping service: %v", service, errs)
			} else {
				delete(self.services, key)
			}
		}
	}

	return e
}

func StopServices(services collections.Sequence) errors.ErrorList {
	e := errors.NewErrorList()

	for it := services.Iterator(); it.HasNext(); {
		serviceMap := it.Next().(ServiceMap)
		for _, service := range serviceMap.services {
			e.AddAll(service.Stop())
		}
	}
	return e
}

type ResourceManager struct {
	io.Closer

	resources []io.Closer
}

func (self *ResourceManager) ManageResource(owned io.Closer) {
	self.resources = append(self.resources, owned)
}

func (self *ResourceManager) Close() error {
	var errs errors.ErrorList
	for i, resource := range self.resources {
		if resource == nil {
			continue
		}

		err := resource.Close()
		if err != nil {
			errs.Add(err)
		}
		self.resources[i] = nil
	}

	return errs.Any()
}
