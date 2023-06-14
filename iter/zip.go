package iter

import (
	"prelude/fn"
	"prelude/maybe"
	"prelude/pair"
)

func ZipWith[A, B, C any](f func(A, B) C, xs Iterator[A], ys Iterator[B]) Iterator[C] {
	g := fn.Curry(f)
	return func() maybe.Maybe[C] {
		withX := maybe.Fmap(g, xs())
		withY := maybe.Apply(withX, ys())
		return withY
	}
}

func Zip[A, B any](xs Iterator[A], ys Iterator[B]) Iterator[pair.Pair[A, B]] {
	return ZipWith(pair.New[A, B], xs, ys)
}
