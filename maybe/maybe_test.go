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
	assert.Equal(t, Maybe[int]{value: nil}, Nothing[int]())
}

func TestIsJust(t *testing.T) {
	assert.True(t, IsJust(Just(1)))
	assert.False(t, IsJust(Nothing[int]()))
}

func TestIsNothing(t *testing.T) {
	assert.True(t, IsNothing(Nothing[int]()))
	assert.False(t, IsNothing(Just(1)))
}

func TestUnwrap(t *testing.T) {
	assert.Equal(t, 1, Unwrap(Just(1)))
	assert.Panics(t, func() {
		Unwrap(Nothing[int]())
	})
}

func TestUnwrapOr(t *testing.T) {
	assert.Equal(t, 1, UnwrapOr(Just(1), 2))
	assert.Equal(t, 2, UnwrapOr(Nothing[int](), 2))
}

func TestMatch(t *testing.T) {
	f := func(n int) int { return n + 1 }
	g := func() int { return -1 }

	x := 1
	assert.Equal(t, f(x), Match(f, g, Just(x)))
	assert.Equal(t, g(), Match(f, g, Nothing[int]()))
}

func TestFmap(t *testing.T) {
	f := func(n int) int { return n + 1 }

	x := 1
	assert.Equal(t, Just(f(x)), Fmap(f, Just(x)))
	assert.Equal(t, Nothing[int](), Fmap(f, Nothing[int]()))
}

func TestApply(t *testing.T) {
	f := func(n int) int { return n + 1 }
	x := 1
	assert.Equal(t, Just(f(x)), Apply(Just(f), Just(x)))
	assert.Equal(t, Nothing[int](), Apply(Just(f), Nothing[int]()))
	assert.Equal(t, Nothing[int](), Apply(Nothing[func(int) int](), Just(x)))
	assert.Equal(t, Nothing[int](), Apply(Nothing[func(int) int](), Nothing[int]()))
}

func TestJoin(t *testing.T) {
	assert.Equal(t, Just(1), Join(Just(Just(1))))
	assert.Equal(t, Nothing[int](), Join(Just(Nothing[int]())))
	assert.Equal(t, Nothing[int](), Join(Nothing[Maybe[int]]()))
}

func TestBind(t *testing.T) {
	f := func(n int) Maybe[int] { return Just(n + 1) }
	g := func(_ int) Maybe[int] { return Nothing[int]() }
	x := 1
	assert.Equal(t, f(x), Bind(f, Just(x)))
	assert.Equal(t, g(x), Bind(g, Just(x)))
	assert.Equal(t, Nothing[int](), Bind(f, Nothing[int]()))
	assert.Equal(t, Nothing[int](), Bind(g, Nothing[int]()))
}
