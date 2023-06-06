package iter

import "prelude/maybe"

type SliceIterator[A any] struct {
	data  []A
	index int
}

func (s *SliceIterator[A]) Next() maybe.Maybe[A] {
	if s.index >= len(s.data) || s.index < 0 {
		return maybe.Nothing[A]()
	}
	elem := s.data[s.index]
	s.index++
	return maybe.Just(elem)
}

func FromSlice[A any](xs []A) Iterator[A] {
	ys := make([]A, len(xs), len(xs))
	copy(ys, xs)
	return &SliceIterator[A]{
		data:  ys,
		index: 0,
	}
}

func ToSlice[A any](xs Iterator[A]) []A {
	data := make([]A, 0)
	for x := xs.Next(); maybe.IsJust(x); x = xs.Next() {
		data = append(data, maybe.Unwrap(x))
	}
	return data
}
