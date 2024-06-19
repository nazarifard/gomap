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

func (m Map[K, V]) Iterator() *mapIter[K, V] {
	it := mapIter[K, V]{
		Ch: make(chan mapIterItem[K, V]),
	}
	go func() {
		defer close(it.Ch)
		for k, v := range m {
			it.Ch <- mapIterItem[K, V]{true, struct {
				Key   K
				Value V
			}{Key: k, Value: v}}
		}

		it.Ch <- mapIterItem[K, V]{false, struct {
			Key   K
			Value V
		}{}}
	}()
	return &it
}
