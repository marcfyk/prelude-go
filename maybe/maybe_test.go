package maybe

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJust(t *testing.T) {
	x := 1
	assert.Equal(t, Maybe[int]{value: &x}, Just(x))
}

func TestNothing(t *testing.T) {
	assert.Equal(t, Maybe[any]{value: nil}, Nothing[any]())
}

func TestIsJust(t *testing.T) {
	assert.True(t, IsJust(Just(1)))
	assert.False(t, IsJust(Nothing[any]()))
}

func TestIsNothing(t *testing.T) {
	assert.True(t, IsNothing(Nothing[any]()))
	assert.False(t, IsNothing(Just(1)))
}

func TestMatch(t *testing.T) {
	f := func(n int) int { return n + 1 }
	g := func() int { return -1 }
	assert.Equal(t, f(1), Match(f, g, Just(1)))
	assert.Equal(t, g(), Match(f, g, Nothing[int]()))
}

func TestUnwrap(t *testing.T) {
	assert.Equal(t, 1, Unwrap(Just(1)))
	assert.PanicsWithError(t, errInvalidUnwrap[any]{value: Nothing[any]()}.Error(), func() {
		Unwrap(Nothing[any]())
	})
}

func TestUnwrapOr(t *testing.T) {
	assert.Equal(t, 1, UnwrapOr(Just(1), 2))
	assert.Equal(t, 2, UnwrapOr(Nothing[any](), 2))
}

func TestFmap(t *testing.T) {
	f := func(n int) int { return n + 1 }

	assert.Equal(t, Just(f(1)), Fmap(f, Just(1)))
	assert.Equal(t, Nothing[int](), Fmap(f, Nothing[int]()))
}

func TestApply(t *testing.T) {
	f := func(n int) int { return n + 1 }
	assert.Equal(t, Just(f(1)), Apply(Just(f), Just(1)))
	assert.Equal(t, Nothing[int](), Apply(Just(f), Nothing[int]()))
	assert.Equal(t, Nothing[int](), Apply(Nothing[func(int) int](), Just(1)))
	assert.Equal(t, Nothing[int](), Apply(Nothing[func(int) int](), Nothing[int]()))
}

func TestJoin(t *testing.T) {
	assert.Equal(t, Just(1), Join(Just(Just(1))))
	assert.Equal(t, Nothing[any](), Join(Just(Nothing[any]())))
	assert.Equal(t, Nothing[any](), Join(Nothing[Maybe[any]]()))
}

func TestBind(t *testing.T) {
	f := func(n int) Maybe[int] { return Just(n + 1) }
	g := func(_ int) Maybe[int] { return Nothing[int]() }
	assert.Equal(t, f(1), Bind(f, Just(1)))
	assert.Equal(t, g(1), Bind(g, Just(1)))
	assert.Equal(t, Nothing[int](), Bind(f, Nothing[int]()))
	assert.Equal(t, Nothing[int](), Bind(g, Nothing[int]()))
}
