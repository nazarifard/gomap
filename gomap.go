package gomap

type Map[K comparable, V any] map[K]V

func New[K comparable, V any](hintSize ...int) Map[K, V] {
	if len(hintSize) == 0 {
		return make(map[K]V)
	}
	return make(map[K]V, hintSize[0])
}

func (m Map[K, V]) Set(key K, value V) {
	m[key] = value
}

func (m Map[K, V]) Update(key K, fn func(V) V) {
	m[key] = fn(m[key]) //old:=m[key]
}

func (m Map[K, V]) Get(key K) (value V, ok bool) {
	value, ok = m[key]
	return
}

func (m Map[K, V]) Len() int {
	return len(m)
}

func (m Map[K, V]) Seq(operationFn func(key K, value V) bool) {
	for k, v := range m {
		next := operationFn(k, v)
		if !next {
			break
		}
	}
}

func (m Map[K, V]) Delete(key K) {
	delete(m, key)
}
