package maybe

import (
	"fmt"
	"prelude/fn"
)

type Maybe[A any] struct {
	value *A
}

func Just[A any](value A) Maybe[A] {
	return Maybe[A]{
		value: &value,
	}
}

func Nothing[A any]() Maybe[A] {
	return Maybe[A]{
		value: nil,
	}
}

func IsJust[A any](maybe Maybe[A]) bool {
	return maybe.value != nil
}

func IsNothing[A any](maybe Maybe[A]) bool {
	return maybe.value == nil
}

func Unwrap[A any](maybe Maybe[A]) A {
	if maybe.value == nil {
		panic(fmt.Sprintf("attempting to unwrap value from %T", maybe))
	}
	return *maybe.value
}

func UnwrapOr[A any](maybe Maybe[A], fallback A) A {
	return Match(
		fn.Id[A],
		fn.Thunk(fallback),
		maybe,
	)
}

func Match[A, B any](justFn func(A) B, nothingFn func() B, maybe Maybe[A]) B {
	if IsJust(maybe) {
		return justFn(*maybe.value)
	} else {
		return nothingFn()
	}
}

func Fmap[A, B any](f func(A) B, maybe Maybe[A]) Maybe[B] {
	return Match(
		fn.Fmap(Just[B], f),
		Nothing[B],
		maybe,
	)
}

func Apply[A, B any](f Maybe[func(A) B], maybe Maybe[A]) Maybe[B] {
	return Match(
		func(g func(A) B) Maybe[B] {
			return Match(
				func(value A) Maybe[B] { return Just(g(value)) },
				Nothing[B],
				maybe,
			)
		},
		Nothing[B],
		f,
	)
}

func Join[A any](maybe Maybe[Maybe[A]]) Maybe[A] {
	return Match(fn.Id[Maybe[A]], Nothing[A], maybe)
}

func Bind[A, B any](f func(A) Maybe[B], maybe Maybe[A]) Maybe[B] {
	return Join(Fmap(f, maybe))
}
