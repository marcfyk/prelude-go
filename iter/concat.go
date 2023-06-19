package iter

import "prelude/maybe"

func Concat[A any](xs, ys Iterator[A]) Iterator[A] {
	return func() maybe.Maybe[A] {
		return maybe.Match(
			maybe.Just[A],
			func() maybe.Maybe[A] {
				return maybe.Match(
					maybe.Just[A],
					maybe.Nothing[A],
					ys(),
				)
			},
			xs(),
		)
	}
}
