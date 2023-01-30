package stream

type MapStream[K, T comparable] struct {
	elements map[K]T
}

func (m *MapStream[K, T]) Unwrap() map[K]T {
	return m.elements
}

func newMapStream[K, T comparable]() *MapStream[K, T] {
	return &MapStream[K, T]{elements: make(map[K]T)}
}

func (m *MapStream[K, T]) Len() int {
	return len(m.elements)
}

func (m *MapStream[K, T]) KeyStream() *ListStream[K] {
	ls := newListStream[K]()
	ls.Append(m.Keys()...)
	return ls
}

func (m *MapStream[K, T]) Keys() []K {
	keys := make([]K, 0, m.Len())
	m.Range(func(k K, _ T) {
		keys = append(keys, k)
	})
	return keys
}

func (m *MapStream[K, T]) Get(k K) (T, bool) {
	v, ok := m.elements[k]
	return v, ok
}

func (m *MapStream[K, T]) Set(k K, v T) {
	m.elements[k] = v
}

func (m *MapStream[K, T]) Delete(k K) *MapStream[K, T] {
	delete(m.elements, k)
	return m
}

func (m *MapStream[K, T]) RangeItem(rangeF func(element T)) *MapStream[K, T] {
	RangeMap[K, T](m.elements, func(key K, v T) {
		rangeF(v)
	})
	return m
}
func (m *MapStream[K, T]) Range(rangeF func(key K, element T)) *MapStream[K, T] {
	RangeMap(m.elements, rangeF)
	return m
}

func (m *MapStream[K, T]) List() *ListStream[T] {
	ls := newListStream[T]()
	m.RangeItem(func(element T) {
		ls.Append(element)
	})
	return ls
}

func (m *MapStream[K, T]) Filter(conditionF func(element T) bool) *MapStream[K, T] {
	filterM := make(map[K]T, len(m.elements))
	m.Range(func(key K, v T) {
		filterM[key] = v
	})
	m.elements = filterM
	return m
}
