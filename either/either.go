package either

import (
	"prelude/fn"
	"prelude/maybe"
)

type Either[A, B any] struct {
	left  maybe.Maybe[A]
	right maybe.Maybe[B]
}

func Left[A, B any](left A) Either[A, B] {
	return Either[A, B]{
		left:  maybe.Just(left),
		right: maybe.Nothing[B](),
	}
}

func Right[A, B any](right B) Either[A, B] {
	return Either[A, B]{
		left:  maybe.Nothing[A](),
		right: maybe.Just(right),
	}
}

func IsLeft[A, B any](either Either[A, B]) bool {
	return maybe.IsJust(either.left) && maybe.IsNothing(either.right)
}

func IsRight[A, B any](either Either[A, B]) bool {
	return maybe.IsJust(either.right) && maybe.IsNothing(either.left)
}

func Match[A, B, C any](leftFn func(A) C, rightFn func(B) C, either Either[A, B]) C {
	if IsLeft(either) {
		return leftFn(maybe.Unwrap(either.left))
	} else if IsRight(either) {
		return rightFn(maybe.Unwrap(either.right))
	} else {
		panic(errInvalidEither[A, B]{value: either})
	}
}

func UnwrapLeft[A, B any](either Either[A, B]) A {
	if !IsLeft(either) {
		panic(errInvalidUnwrapLeft[A, B]{value: either})
	}
	return maybe.Unwrap(either.left)
}

func UnwrapLeftOr[A, B any](either Either[A, B], fallback A) A {
	return Match(
		fn.Id[A],
		fn.Const[A, B](fallback),
		either,
	)
}

func UnwrapRight[A, B any](either Either[A, B]) B {
	if !IsRight(either) {
		panic(errInvalidUnwrapRight[A, B]{value: either})
	}
	return maybe.Unwrap(either.right)
}

func UnwrapRightOr[A, B any](either Either[A, B], fallback B) B {
	return Match(
		fn.Const[B, A](fallback),
		fn.Id[B],
		either,
	)
}

func Fmap[A, B, C any](f func(B) C, either Either[A, B]) Either[A, C] {
	return Match(
		Left[A, C],
		fn.Fmap(Right[A, C], f),
		either,
	)
}

func Apply[A, B, C any](f Either[A, func(B) C], either Either[A, B]) Either[A, C] {
	return Match(
		Left[A, C],
		fn.Flip(fn.Curry(Fmap[A, B, C]))(either),
		f,
	)
}

func Join[A, B any](either Either[A, Either[A, B]]) Either[A, B] {
	return Match(
		Left[A, B],
		fn.Id[Either[A, B]],
		either,
	)
}

func Bind[A, B, C any](f func(B) Either[A, C], either Either[A, B]) Either[A, C] {
	return Join(Fmap(f, either))
}
