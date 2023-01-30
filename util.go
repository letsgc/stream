package stream

func RangeSliceItem[T](slice []T, rangeF func(element T)) {
	RangeSlice(slice, func(idx int, element T) { rangeF(element) })
}

func RangeSlice[T](slice []T, rangeF func(idx int, element T)) {
	for idx := range slice {
		element := slice[idx]
		rangeF(idx, element)
	}
}

func RangeMapItem[T](m map[any]T, rangeF func(element T)) {
	RangeMap(m, func(key any, element T) { rangeF(element) })
}

func RangeMap[K, T](m map[K]T, rangeF func(key K, element T)) {
	for key, element := range m {
		rangeF(key, element)
	}
}
