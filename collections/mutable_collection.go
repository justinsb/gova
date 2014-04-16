package collections

type MutableCollection interface {
	Collection
	Add(item interface{})
}
