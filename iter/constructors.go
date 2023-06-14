package iter

import (
	"prelude/maybe"
	"prelude/pair"
)

func FromSlice[A any](xs []A) Iterator[A] {
	ys := make([]A, len(xs), len(xs))
	copy(ys, xs)
	index := 0
	return func() maybe.Maybe[A] {
		if index >= len(ys) || index < 0 {
			return maybe.Nothing[A]()
		}
		elem := ys[index]
		index++
		return maybe.Just(elem)
	}
}

func ToSlice[A any](xs Iterator[A]) []A {
	data := make([]A, 0)
	for x := xs(); maybe.IsJust(x); x = xs() {
		data = append(data, maybe.Unwrap(x))
	}
	return data
}

func FromMap[K comparable, V any](xs map[K]V) Iterator[pair.Pair[K, V]] {
	ys := make([]pair.Pair[K, V], len(xs), len(xs))
	{
		index := 0
		for k, v := range xs {
			ys[index] = pair.New(k, v)
			index++
		}
	}
	index := 0
	return func() maybe.Maybe[pair.Pair[K, V]] {
		if index >= len(ys) || index < 0 {
			return maybe.Nothing[pair.Pair[K, V]]()
		}
		elem := ys[index]
		index++
		return maybe.Just(elem)

	}
}

func ToMap[K comparable, V any](xs Iterator[pair.Pair[K, V]]) map[K]V {
	data := make(map[K]V)
	for x := xs(); maybe.IsJust(x); x = xs() {
		p := maybe.Unwrap(x)
		data[p.Left] = p.Right
	}
	return data
}
