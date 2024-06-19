package gomap

type mapIterItem[K comparable, V any] struct {
	Next    bool
	MapItem struct {
		Key   K
		Value V
	}
}

type mapIter[K comparable, V any] struct {
	Ch      chan mapIterItem[K, V]
	MapItem struct {
		Key   K
		Value V
	}
}

func (it *mapIter[K, V]) Next() bool {
	n := <-it.Ch
	it.MapItem = n.MapItem
	return n.Next
}

func (it *mapIter[K, V]) Value() V {
	return it.MapItem.Value
}

func (it *mapIter[K, V]) Key() K {
	return it.MapItem.Key
}
