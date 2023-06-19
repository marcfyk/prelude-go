package iter

import "prelude/maybe"

func ContainsBy[A any](f func(A) bool, xs Iterator[A]) bool {
	return maybe.IsJust(FindBy(f, xs))
}

func Contains[A comparable](x A, xs Iterator[A]) bool {
	return maybe.IsJust(Find(x, xs))
}
