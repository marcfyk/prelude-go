package iter

import (
	"prelude/maybe"
	"prelude/pair"
)

func DropWhile[A any](f func(A) bool, xs Iterator[A]) Iterator[A] {
	shouldDrop := true
	return func() maybe.Maybe[A] {
		if shouldDrop {
			x := xs()
			for maybe.IsJust(x) && f(maybe.Unwrap(x)) {
				x = xs()
			}
			shouldDrop = false
			return x
		} else {
			return xs()
		}
	}
}

func Drop[A any](limit uint, xs Iterator[A]) Iterator[A] {
	enumerated := Enumerate(xs)
	taken := DropWhile(func(p pair.Pair[uint, A]) bool { return p.Left < limit }, enumerated)
	return Map(pair.Right[uint, A], taken)
}
