package inject

import (
	"fmt"
	"reflect"

	"github.com/justinsb/gova/log"
)

type Binder struct {
	bindings map[reflect.Type]Binding
}

func NewBinder() *Binder {
	self := &Binder{}
	self.bindings = make(map[reflect.Type]Binding)
	return self
}

type BindPartial struct {
	binder *Binder
	t      reflect.Type
}

func (self *Binder) Bind(prototype interface{}) BindPartial {
	t := reflect.TypeOf(prototype)
	return BindPartial{binder: self, t: t}
}

func (self *Binder) BindType(t reflect.Type) BindPartial {
	return BindPartial{binder: self, t: t}
}

func (self BindPartial) ToInstance(instance interface{}) {
	self.binder.addBinding(self.t, instance)
}

func (self *Binder) AddProvider(fn interface{}) {
	binding := &FunctionBinding{}
	binding.fn = fn
	binding.valFn = reflect.ValueOf(fn)

	if binding.valFn.Type().Kind() != reflect.Func {
		panic("Binding to invalid provider kind")
	}

	numOut := binding.valFn.Type().NumOut()
	if numOut != 1 && numOut != 2 {
		panic("Invalid number of return values from function provider")
	}

	returnType := binding.valFn.Type().Out(0)
	self.bindings[returnType] = binding
}

func (self *Binder) AddSingleton(obj interface{}) {
	t := reflect.TypeOf(obj)
	self.addBinding(t, obj)
}

func (self *Binder) AddDefaultBinding(t reflect.Type) {
	binding := NewDefaultBinding(t)
	self.bindings[t] = binding
}

func (self *Binder) AddDefaultBindingByPointer(p interface{}) {
	t := reflect.TypeOf(p).Elem()
	log.Info("Type is %v", t)
	self.AddDefaultBinding(t)
}

func (self *Binder) addBinding(t reflect.Type, obj interface{}) {
	binding := &SingletonBinding{}
	binding.obj = obj
	log.Debug("Binding type %v to %v", t, obj)
	self.bindings[t] = binding
}

func (self *Binder) Get(t reflect.Type) (interface{}, error) {
	binding, found := self.bindings[t]
	if found {
		return binding.Get()
	}

	// If we requested *T, and we have a binding for T, we can satisfy that
	if t.Kind() == reflect.Ptr {
		tElem := t.Elem()
		binding, found = self.bindings[tElem]
		if found {
			v, err := binding.Get()
			if err != nil {
				return nil, err
			}
			return v, nil
		}
	}

	return nil, fmt.Errorf("No binding for type %v", t)
}

func (self *Binder) CreateInjector() Injector {
	i := &BinderInjector{}
	i.binder = self
	self.addBinding(reflect.TypeOf((*Injector)(nil)).Elem(), i)
	return i
}
