package iter

import (
	"prelude/fn"
	"prelude/maybe"
)

func Foldl[A, B any](f func(B, A) B, initial B, xs Iterator[A]) B {
	output := initial
	for x := xs(); maybe.IsJust(x); x = xs() {
		output = f(output, maybe.Unwrap(x))
	}
	return output
}

func Reduce[A any](f func(A, A) A, xs Iterator[A]) maybe.Maybe[A] {
	g := fn.Flip(fn.Curry(f))
	return Foldl(func(acc maybe.Maybe[A], x A) maybe.Maybe[A] {
		return maybe.Fmap(g(x), acc)
	}, xs(), xs)
}
