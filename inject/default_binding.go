package inject

import "reflect"

type DefaultBinding struct {
	t reflect.Type
}

func NewDefaultBinding(t reflect.Type) *DefaultBinding {
	self := &DefaultBinding{}
	self.t = t
	return self
}

func (self *DefaultBinding) Get() (interface{}, error) {
	v := reflect.New(self.t)
	return v.Interface(), nil
}
