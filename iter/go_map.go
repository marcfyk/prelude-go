package iter

import (
	"prelude/maybe"
	"prelude/pair"
)

type MapIterator[K comparable, V any] struct {
	iterator SliceIterator[pair.Pair[K, V]]
}

func (m *MapIterator[K, V]) Next() maybe.Maybe[pair.Pair[K, V]] {
	return m.iterator.Next()
}

func FromMap[K comparable, V any](xs map[K]V) Iterator[pair.Pair[K, V]] {
	data := make([]pair.Pair[K, V], len(xs), len(xs))
	index := 0
	for k, v := range xs {
		data[index] = pair.New(k, v)
		index++
	}
	return &MapIterator[K, V]{
		iterator: SliceIterator[pair.Pair[K, V]]{
			data:  data,
			index: 0,
		},
	}
}

func ToMap[K comparable, V any](xs Iterator[pair.Pair[K, V]]) map[K]V {
	data := make(map[K]V)
	for x := xs.Next(); maybe.IsJust(x); x = xs.Next() {
		p := maybe.Unwrap(x)
		data[p.Left] = p.Right
	}
	return data
}
