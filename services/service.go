package services

import (
	//	"github.com/justinsb/gova/collections"
	"github.com/justinsb/gova/errors"
	//	"github.com/justinsb/slf4g/log"

	//	"fmt"
	"io"
)

type Service interface {
	Start() errors.ErrorList
	Stop() errors.ErrorList
	Update(config interface{}) errors.ErrorList

	ManageResource(owned io.Closer)
}

//// +gen
//type ServiceConfig interface {
//	Key() string
//}

//func StopServices(services collections.Sequence) errors.ErrorList {
//	e := errors.NewErrorList()
//
//	for it := services.Iterator(); it.HasNext(); {
//		serviceMap := it.Next().(ServiceMap)
//		for _, service := range serviceMap.services {
//			e.AddAll(service.Stop())
//		}
//	}
//	return e
//}

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

//type ServiceConfigChangeTracker struct {
//}
//
//type TombstoneServiceConfig struct {
//	key string
//}
//
//func (self *TombstoneServiceConfig) Key() string {
//	return self.key
//}
//
//func NewServiceConfigChangeTracker(configsets <-chan []ServiceConfig) <-chan ServiceConfig {
//	out := make(chan ServiceConfig)
//
//	go func() {
//		previous := make(map[string]ServiceConfig)
//
//		for configset := range configsets {
//			next := make(map[string]ServiceConfig)
//			for _, config := range configset {
//				key := config.Key()
//				previousConfig := previous[key]
//				if previousConfig == nil {
//					// New
//					out <- config
//				} else {
//					// We don't bother with change-detection
//					out <- config
//				}
//				next[key] = config
//			}
//
//			for key, config := range previous {
//				_, found := next[key]
//				if !found {
//					// Removed
//					config := &TombstoneServiceConfig{}
//					config.key = key
//					out <- config
//				}
//			}
//
//		}
//		close(out)
//	}()
//}
