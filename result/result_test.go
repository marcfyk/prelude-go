package result

import (
	"errors"
	"prelude/either"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOk(t *testing.T) {
	assert.Equal(t, Result[int]{value: either.Right[error](1)}, Ok(1))
}

func TestErr(t *testing.T) {
	err := errors.New("some error")
	assert.Equal(t, Result[any]{value: either.Left[error, any](err)}, Err[any](err))
}

func TestIsOk(t *testing.T) {
	assert.True(t, IsOk(Ok(1)))
	assert.False(t, IsOk(Err[int](errors.New("some error"))))
}

func TestIsErr(t *testing.T) {
	assert.True(t, IsErr(Err[int](errors.New("some error"))))
	assert.False(t, IsErr(Ok(1)))
}

func TestMatch(t *testing.T) {
	f := func(e error) int { return len(e.Error()) }
	g := func(n int) int { return n + 1 }
	err := errors.New("some error")
	assert.Equal(t, f(err), Match(f, g, Err[int](err)))
	assert.Equal(t, g(1), Match(f, g, Ok(1)))
}

func TestUnwrapErr(t *testing.T) {
	err := errors.New("some error")
	assert.Equal(t, err, UnwrapErr(Err[any](err)))
	assert.Panics(t, func() { UnwrapErr(Ok(1)) })
}

func TestUnwrapErrOr(t *testing.T) {
	err1 := errors.New("error 1")
	err2 := errors.New("error 2")
	assert.Equal(t, err1, UnwrapErrOr(Err[any](err1), err2))
	assert.Equal(t, err2, UnwrapErrOr(Ok(1), err2))
}

func TestUnwrapOk(t *testing.T) {
	assert.Equal(t, 1, UnwrapOk(Ok(1)))
	assert.Panics(t, func() { UnwrapOk(Err[any](errors.New("some error"))) })
}

func TestUnwrapOkOr(t *testing.T) {
	assert.Equal(t, 1, UnwrapOkOr(Ok(1), 2))
	assert.Equal(t, 2, UnwrapOkOr(Err[int](errors.New("some error")), 2))
}

func TestFmap(t *testing.T) {
	f := func(n int) int { return n + 1 }
	assert.Equal(t, Ok(f(1)), Fmap(f, Ok(1)))
	err := errors.New("some error")
	assert.Equal(t, Err[int](err), Fmap(f, Err[int](err)))
}

func TestApply(t *testing.T) {
	f := func(n int) int { return n + 1 }
	assert.Equal(t, Ok(f(1)), Apply(Ok(f), Ok(1)))
	err1 := errors.New("error 1")
	err2 := errors.New("error 2")
	assert.Equal(t, Err[int](err1), Apply(Err[func(int) int](err1), Err[int](err2)))
	assert.Equal(t, Err[int](err1), Apply(Err[func(int) int](err1), Ok(1)))
	assert.Equal(t, Err[int](err2), Apply(Ok(f), Err[int](err2)))
}

func TestJoin(t *testing.T) {
	err := errors.New("some error")
	assert.Equal(t, Err[any](err), Join(Err[Result[any]](err)))
	assert.Equal(t, Err[any](err), Join(Ok(Err[any](err))))
	assert.Equal(t, Ok(1), Join(Ok(Ok(1))))
}

func TestBind(t *testing.T) {
	err1 := errors.New("error 1")
	err2 := errors.New("error 2")
	f := func(n int) Result[int] { return Ok(n + 1) }
	g := func(_ int) Result[int] { return Err[int](err1) }
	assert.Equal(t, f(1), Bind(f, Ok(1)))
	assert.Equal(t, g(1), Bind(g, Ok(1)))
	assert.Equal(t, Err[int](err2), Bind(f, Err[int](err2)))
	assert.Equal(t, Err[int](err2), Bind(g, Err[int](err2)))
}
