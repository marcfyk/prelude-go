package iter

import "prelude/maybe"

type Iterator[A any] func() maybe.Maybe[A]

func Empty[A any]() Iterator[A] {
	return maybe.Nothing[A]
}

func Single[A any](value A) Iterator[A] {
	return FromSlice([]A{value})
}
