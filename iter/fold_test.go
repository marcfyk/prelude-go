package iter

import (
	"prelude/maybe"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFoldl(t *testing.T) {
	f := func(acc int, n int) int { return acc - n }
	xs := FromSlice([]int{3, 2, 5})
	folded := Foldl(f, 0, xs)
	assert.Equal(t, f(f(f(0, 3), 2), 5), folded)
}

func TestReduce(t *testing.T) {
	f := func(acc int, n int) int { return acc - n }
	xs := FromSlice([]int{3, 2, 5})
	reduced := Reduce(f, xs)
	assert.Equal(t, maybe.Just(f(f(3, 2), 5)), reduced)

	ys := FromSlice([]int{})
	reduced = Reduce(f, ys)
	assert.Equal(t, maybe.Nothing[int](), reduced)
}
