package gomap

type mapIterItem[K comparable, V any] struct {
	next    bool
	mapItem struct {
		Key   K
		Value V
	}
}

type mapIter[K comparable, V any] struct {
	ch      chan mapIterItem[K, V]
	mapItem struct {
		Key   K
		Value V
	}
}

func (it *mapIter[K, V]) Next() bool {
	n := <-it.ch
	it.mapItem = n.mapItem
	return n.next
}

func (it *mapIter[K, V]) Value() V {
	return it.mapItem.Value
}

func (it *mapIter[K, V]) Key() K {
	return it.mapItem.Key
}
