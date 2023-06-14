package iter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTakeWhile(t *testing.T) {
	xs := FromSlice([]int{1, 2, 3, 4, 5, 4, 3, 2, 1})
	f := func(n int) bool { return n < 3 }
	taken := TakeWhile(f, xs)
	ys := ToSlice(taken)
	assert.Equal(t, []int{1, 2}, ys)
}

func TestTake(t *testing.T) {
	xs := FromSlice([]int{5, 4, 3, 2, 1})
	taken := Take(2, xs)
	ys := ToSlice(taken)
	assert.Equal(t, []int{5, 4}, ys)
}
