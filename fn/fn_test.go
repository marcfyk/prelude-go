package fn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestId(t *testing.T) {
	value := 1
	assert.Equal(t, value, Id(value))
}

func TestThunk(t *testing.T) {
	value := 1
	f := Thunk(value)
	assert.Equal(t, value, f())
}

func TestFlip(t *testing.T) {
	f := func(x int) func(int) int {
		return func(y int) int {
			return x - y
		}
	}
	x, y := 1, 2
	assert.Equal(t, f(y)(x), Flip(f)(x)(y))
}

func TestCurry(t *testing.T) {
	f := func(x int, y int) int { return x - y }
	curried := Curry(f)
	x, y := 1, 2
	assert.Equal(t, f(x, y), curried(x)(y))
}

func TestUncurry(t *testing.T) {
	f := func(x int) func(int) int {
		return func(y int) int {
			return x - y
		}
	}
	uncurried := Uncurry(f)
	x, y := 1, 2
	assert.Equal(t, f(x)(y), uncurried(x, y))
}

func TestFmap(t *testing.T) {
	g := func(x int) int { return x * 3 }
	f := func(x int) int { return x + 1 }
	fmapped := Fmap(g, f)
	x := 10
	assert.Equal(t, g(f(x)), fmapped(x))
}

func TestApply(t *testing.T) {
	g := func(x int) func(int) int {
		return func(y int) int {
			return x - y
		}
	}
	f := func(x int) int { return x + 1 }
	applied := Apply(g, f)
	x := 5
	assert.Equal(t, g(x)(f(x)), applied(x))
}

func TestJoin(t *testing.T) {
	f := func(x int) func(int) int {
		return func(y int) int {
			return x - y
		}
	}
	joined := Join(f)
	x := 5
	assert.Equal(t, f(x)(x), joined(x))
}

func TestBind(t *testing.T) {
	g := func(x int) func(int) int {
		return func(y int) int {
			return x - y
		}
	}
	f := func(x int) int { return x + 1 }
	binded := Bind(g, f)
	x := 5
	assert.Equal(t, g(f(x))(x), binded(x))
}

func TestCompose(t *testing.T) {
	g := func(x int) int { return x * 3 }
	f := func(x int) int { return x + 1 }
	fmapped := Fmap(g, f)
	x := 10
	assert.Equal(t, g(f(x)), fmapped(x))
}

func TestConst(t *testing.T) {
	assert.Equal(t, 1, Const[int, int](1)(2))
}
