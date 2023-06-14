package iter

import "prelude/maybe"

func Map[A, B any](f func(A) B, xs Iterator[A]) Iterator[B] {
	return func() maybe.Maybe[B] {
		return maybe.Fmap(f, xs())
	}
}
