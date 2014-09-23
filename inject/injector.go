package inject

import "reflect"

type Injector interface {
	Get(t reflect.Type) (interface{}, error)
	Inject(p interface{}) error
}

type BinderInjector struct {
	binder *Binder
}

func (self *BinderInjector) Inject(p interface{}) error {
	pType := reflect.TypeOf(p)
	t := pType.Elem()
	v, err := self.Get(t)
	if err != nil {
		return err
	}

	val := reflect.ValueOf(p)
	val.Elem().Set(reflect.ValueOf(v))
	return nil
}

func (self *BinderInjector) Get(t reflect.Type) (interface{}, error) {
	v, err := self.binder.Get(t)
	if err != nil {
		return nil, err
	}

	err = self.injectFields(v)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func (self *BinderInjector) injectFields(v interface{}) error {
	vValue := reflect.ValueOf(v)

	for vValue.Kind() == reflect.Ptr {
		vValue = vValue.Elem()
	}

	if vValue.Kind() == reflect.Struct {
		vType := vValue.Type()
		n := vType.NumField()
		for i := 0; i < n; i++ {
			field := vType.Field(i)
			injectTag := field.Tag.Get("inject")
			if injectTag != "" {
				fieldValue := vValue.Field(i)
				err := self.injectField(&fieldValue, &field)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (self *BinderInjector) injectField(fieldValue *reflect.Value, field *reflect.StructField) error {
	t := field.Type
	v, err := self.Get(t)
	if err != nil {
		return err
	}

	fieldValue.Set(reflect.ValueOf(v))
	return nil
}
