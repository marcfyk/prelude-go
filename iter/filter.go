package iter

import "prelude/maybe"

func Filter[A any](f func(A) bool, xs Iterator[A]) Iterator[A] {
	return func() maybe.Maybe[A] {
		x := xs()
		for maybe.IsJust(x) {
			elem := maybe.Unwrap(x)
			if f(elem) {
				break
			}
			x = xs()
		}
		return x
	}
}
