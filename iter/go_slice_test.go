package iter

import (
	"prelude/maybe"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFromSlice(t *testing.T) {
	xs := []int{3, 1, 2}
	ys := FromSlice(xs)

	assert.Equal(t, maybe.Just(3), ys.Next())
	assert.Equal(t, maybe.Just(1), ys.Next())
	assert.Equal(t, maybe.Just(2), ys.Next())
	assert.Equal(t, maybe.Nothing[int](), ys.Next())
}

func TestFromSliceCopy(t *testing.T) {
	xs := []int{3, 1, 2}
	ys := FromSlice(xs)
	xs = append(xs, 4)
	xs[0] = 5

	assert.Equal(t, maybe.Just(3), ys.Next())
	assert.Equal(t, maybe.Just(1), ys.Next())
	assert.Equal(t, maybe.Just(2), ys.Next())
	assert.Equal(t, maybe.Nothing[int](), ys.Next())
}

func TestToSlice(t *testing.T) {
	xs := FromSlice([]int{3, 1, 2})
	assert.Equal(t, []int{3, 1, 2}, ToSlice(xs))
}
