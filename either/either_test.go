package either

import (
	"prelude/maybe"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeft(t *testing.T) {
	x := 1
	assert.Equal(
		t,
		Either[int, any]{
			left:  maybe.Just(x),
			right: maybe.Nothing[any](),
		},
		Left[int, any](x),
	)
}

func TestRight(t *testing.T) {
	x := 1
	assert.Equal(
		t,
		Either[any, int]{
			left:  maybe.Nothing[any](),
			right: maybe.Just(x),
		},
		Right[any](x),
	)
}

func TestIsLeft(t *testing.T) {
	assert.True(t, IsLeft(Left[int, any](1)))
	assert.False(t, IsLeft(Right[any]("abc")))
}

func TestIsRight(t *testing.T) {
	assert.True(t, IsRight(Right[any]("abc")))
	assert.False(t, IsRight(Left[int, any](1)))
}

func TestMatch(t *testing.T) {
	f := func(n int) int { return n + 1 }
	g := func(n int) int { return n - 1 }
	x := 1
	assert.Equal(t, f(x), Match(f, g, Left[int, int](x)))
	assert.Equal(t, g(x), Match(f, g, Right[int](x)))

	invalidEither := Either[int, int]{
		left:  maybe.Just(1),
		right: maybe.Just(1),
	}
	assert.PanicsWithError(
		t,
		errInvalidEither[int, int]{value: invalidEither}.Error(),
		func() { Match(f, g, invalidEither) },
	)
}

func TestUnwrapLeft(t *testing.T) {
	assert.Equal(t, 1, UnwrapLeft(Left[int, any](1)))
	assert.PanicsWithError(
		t,
		errInvalidUnwrapLeft[any, int]{value: Right[any](1)}.Error(),
		func() {
			UnwrapLeft(Right[any](1))
		},
	)
}

func TestUnwrapLeftOr(t *testing.T) {
	assert.Equal(t, 1, UnwrapLeftOr(Left[int, any](1), 2))
	assert.Equal(t, 2, UnwrapLeftOr(Right[any](1), 2))
}

func TestUnwrapRight(t *testing.T) {
	assert.Equal(t, 1, UnwrapRight(Right[any](1)))
	assert.PanicsWithError(
		t,
		errInvalidUnwrapRight[int, any]{value: Left[int, any](1)}.Error(),
		func() {
			UnwrapRight(Left[int, any](1))
		},
	)
}

func TestUnwrapRightOr(t *testing.T) {
	assert.Equal(t, 1, UnwrapRightOr(Right[any](1), 2))
	assert.Equal(t, 2, UnwrapRightOr(Left[int, any](1), 2))
}

func TestFmap(t *testing.T) {
	f := func(n int) int { return n + 1 }
	x := 1
	assert.Equal(t, Left[int, int](x), Fmap(f, Left[int, int](x)))
	assert.Equal(t, Right[any](f(x)), Fmap(f, Right[any](x)))
}

func TestApply(t *testing.T) {
	f := func(n int) int { return n + 1 }
	assert.Equal(t, Left[int, int](1), Apply(Left[int, func(int) int](1), Left[int, int](2)))
	assert.Equal(t, Left[int, int](1), Apply(Left[int, func(int) int](1), Right[int](2)))
	assert.Equal(t, Left[int, int](2), Apply(Right[int](f), Left[int, int](2)))
	assert.Equal(t, Right[int](f(2)), Apply(Right[int](f), Right[int](2)))
}

func TestJoin(t *testing.T) {
	assert.Equal(t, Right[any](1), Join(Right[any](Right[any](1))))
	assert.Equal(t, Left[int, int](2), Join(Right[int](Left[int, int](2))))
	assert.Equal(t, Left[int, int](3), Join(Left[int, Either[int, int]](3)))
}

func TestBind(t *testing.T) {
	f := func(n int) Either[int, int] { return Right[int](n + 1) }
	g := func(n int) Either[int, int] { return Left[int, int](n - 1) }
	assert.Equal(t, f(1), Bind(f, Right[int](1)))
	assert.Equal(t, g(1), Bind(g, Right[int](1)))
	assert.Equal(t, Left[int, int](1), Bind(f, Left[int, int](1)))
	assert.Equal(t, Left[int, int](1), Bind(g, Left[int, int](1)))
}
