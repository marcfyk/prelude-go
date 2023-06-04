package result

import (
	"prelude/either"
	"prelude/fn"
)

type Result[A any] struct {
	value either.Either[error, A]
}

func Ok[A any](value A) Result[A] {
	return Result[A]{
		value: either.Right[error](value),
	}
}

func Err[A any](err error) Result[A] {
	return Result[A]{
		value: either.Left[error, A](err),
	}
}

func IsOk[A any](result Result[A]) bool {
	return either.IsRight(result.value)
}

func IsErr[A any](result Result[A]) bool {
	return either.IsLeft(result.value)
}

func Match[A, B any](errFn func(e error) B, okFn func(A) B, result Result[A]) B {
	return either.Match(errFn, okFn, result.value)
}

func UnwrapErr[A any](result Result[A]) error {
	return either.UnwrapLeft(result.value)
}

func UnwrapErrOr[A any](result Result[A], fallback error) error {
	return either.UnwrapLeftOr(result.value, fallback)
}

func UnwrapOk[A any](result Result[A]) A {
	return either.UnwrapRight(result.value)
}

func UnwrapOkOr[A any](result Result[A], fallback A) A {
	return either.UnwrapRightOr(result.value, fallback)
}

func Fmap[A, B any](f func(A) B, result Result[A]) Result[B] {
	return Match(
		Err[B],
		fn.Fmap(Ok[B], f),
		result,
	)
}

func Apply[A, B any](f Result[func(A) B], result Result[A]) Result[B] {
	return Match(
		Err[B],
		fn.Flip(fn.Curry(Fmap[A, B]))(result),
		f,
	)
}

func Join[A any](result Result[Result[A]]) Result[A] {
	return Match(
		Err[A],
		fn.Id[Result[A]],
		result,
	)
}

func Bind[A, B any](f func(A) Result[B], result Result[A]) Result[B] {
	return Join(Fmap(f, result))
}
