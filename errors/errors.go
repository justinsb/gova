package errors

type ErrorList struct {
	errors []error
}

func NoErrors() ErrorList {
	return NewErrorList()
}

func NewErrorList() ErrorList {
	self := ErrorList{}
	self.errors = nil
	return self
}

func (self *ErrorList) IsEmpty() bool {
	if self.errors == nil || len(self.errors) == 0 {
		return true
	}
	return false
}

func (self *ErrorList) Any() error {
	if self.errors == nil || len(self.errors) == 0 {
		return nil
	}
	return self.errors[0]
}

func (self *ErrorList) AddAll(errors ErrorList) {
	if !errors.IsEmpty() {
		if self.errors == nil {
			self.errors = make([]error, 0)
		}

		self.errors = append(self.errors, errors.errors...)
	}
}

func (self *ErrorList) Add(err error) {
	if err != nil {
		if self.errors == nil {
			self.errors = make([]error, 0)
		}

		self.errors = append(self.errors, err)
	}
}
