package iter

import "prelude/maybe"

func FindBy[A any](f func(A) bool, xs Iterator[A]) maybe.Maybe[A] {
	return Filter(f, xs)()
}

func Find[A comparable](x A, xs Iterator[A]) maybe.Maybe[A] {
	return FindBy(func(y A) bool { return y == x }, xs)
}
