package iter

import "prelude/maybe"

func Iterate[A any](f func(A) A, initial A) Iterator[A] {
	x := initial
	return func() maybe.Maybe[A] {
		elem := x
		x = f(x)
		return maybe.Just(elem)
	}
}
