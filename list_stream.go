package stream

type ListStream[T comparable] struct {
	elements []T
}

func newListStream[T comparable]() *ListStream[T] {
	return &ListStream[T]{elements: make([]T, 0)}
}

const (
	Boundary int = -1
)

func (l *ListStream[T]) Unwrap() []T {
	return l.elements
}

func (l *ListStream[T]) Slice(left, right int) *ListStream[T] {
	if left == Boundary {
		left = 0
	}
	if right == Boundary {
		right = len(l.elements)
	}

	l.elements = l.elements[left:right]
	return l
}

func (l *ListStream[T]) Append(elements ...T) *ListStream[T] {
	l.elements = append(l.elements, elements)
	return l
}

func (l *ListStream[T]) Index(i int) T {
	return l.elements[i]
}

func (l *ListStream[T]) Filter(conditionF func(element T) bool) *ListStream[T] {
	newElements := make([]T, 0, len(l.elements))
	l.RangeItem(func(element T) {
		if conditionF(element) {
			newElements = append(newElements, element)
		}
	})
	l.elements = newElements
	return l
}

func (l *ListStream[T]) Map(keyF func(element T) any) *MapStream[any, T] {
	ms := newMapStream[any, T]()
	l.RangeItem(func(element T) {
		ms.Set(keyF(element), element)
	})
	return ms
}

func (l *ListStream[T]) Set() *SetStream[T] {
	ss := newSetStream[T]()
	l.RangeItem(func(element T) {
		ss.Add(element)
	})
	return ss
}

func (l *ListStream[T]) RangeItem(rangeF func(element T)) *ListStream[T] {
	RangeSliceItem(l.elements, rangeF)
	return l
}
func (l *ListStream[T]) Range(rangeF func(idx int, element T)) *ListStream[T] {
	RangeSlice(l.elements, rangeF)
	return l
}
