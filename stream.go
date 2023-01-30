package stream

func ListItems[T comparable](items []T) *ListStream[T] {
	ls := newListStream[T]()
	ls.Append(items...)
	return ls
}

func List[T comparable]() *ListStream[T] {
	return newListStream[T]()
}
func Map[K, T comparable]() *MapStream[K, T] {
	return newMapStream[K, T]()
}
func Set[T comparable]() *SetStream[T] {
	return newSetStream[T]()
}
