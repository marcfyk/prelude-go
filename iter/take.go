package iter

import (
	"prelude/maybe"
	"prelude/pair"
)

func TakeWhile[A any](f func(A) bool, xs Iterator[A]) Iterator[A] {
	shouldTake := true
	return func() maybe.Maybe[A] {
		if shouldTake {
			g := func(x A) maybe.Maybe[A] {
				if f(x) {
					return maybe.Just(x)
				} else {
					shouldTake = false
					return maybe.Nothing[A]()
				}
			}
			return maybe.Bind(g, xs())
		} else {
			return maybe.Nothing[A]()
		}
	}
}

func Take[A any](limit uint, xs Iterator[A]) Iterator[A] {
	enumerated := Enumerate(xs)
	taken := TakeWhile(func(p pair.Pair[uint, A]) bool { return p.Left < limit }, enumerated)
	return Map(pair.Right[uint, A], taken)
}
