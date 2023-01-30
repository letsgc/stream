package stream

type SetStream[T comparable] struct {
	elements map[T]struct{}
}

func newSetStream[T comparable]() *SetStream[T] {
	return &SetStream[T]{elements: make(map[T]struct{})}
}

func (s *SetStream[T]) Add(element T) {
	s.elements[element] = struct{}{}
}

func (s *SetStream[T]) Remove(element T) *SetStream[T] {
	delete(s.elements, element)
	return s
}

func (s *SetStream[T]) Contain(element T) bool {
	_, ok := s.elements[element]
	return ok
}

func (s *SetStream[T]) Range(rangeF func(element T)) *SetStream[T] {
	RangeMap[T, struct{}](s.elements, func(element T, _ struct{}) {
		rangeF(element)
	})
	return s
}

func (s *SetStream[T]) List() *ListStream[T] {
	ls := newListStream[T]()
	s.Range(func(element T) {
		ls.Append(element)
	})
	return ls
}
