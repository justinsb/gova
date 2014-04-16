package collections

type Iterator interface {
	Next() interface{}
	HasNext() bool
}
