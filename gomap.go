package gomap

type Map[K comparable, V any] map[K]V

func New[K comparable, V any]() Map[K, V] {
	return make(map[K]V)
}

func (m Map[K, V]) Set(key K, value V) {
	m[key] = value
}

func (m Map[K, V]) Get(key K) (value V, ok bool) {
	value, ok = m[key]
	return
}

func (m Map[K, V]) Len() int {
	return len(m)
}

func (m Map[K, V]) Range(operationFn func(key K, value V) bool) {
	for k, v := range m {
		next := operationFn(k, v)
		if !next {
			break
		}
	}
}

// type MapIterator[K comparable, V any] interface {
// 	Next() bool
// 	Value() V
// 	Key() K
// }

// func (m Map[K, V]) Iterator() *mapIter[K, V] {
// 	it := mapIter[K, V]{
// 		ch: make(chan mapIterItem[K, V]),
// 	}
// 	go func() {
// 		defer close(it.ch)
// 		for k, v := range m {
// 			it.ch <- mapIterItem[K, V]{true, struct {
// 				Key   K
// 				Value V
// 			}{Key: k, Value: v}}
// 		}

// 		it.ch <- mapIterItem[K, V]{false, struct {
// 			Key   K
// 			Value V
// 		}{}}
// 	}()
// 	return &it
// }
