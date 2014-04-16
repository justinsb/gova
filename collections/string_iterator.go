package collections

type StringIterator interface {
	Iterator

	NextString() string
}
