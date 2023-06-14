package iter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDropWhile(t *testing.T) {
	xs := FromSlice([]int{1, 2, 3, 4, 5, 4, 3, 2, 1})
	f := func(n int) bool { return n < 3 }
	dropped := DropWhile(f, xs)
	ys := ToSlice(dropped)
	assert.Equal(t, []int{3, 4, 5, 4, 3, 2, 1}, ys)
}

func TestDrop(t *testing.T) {
	xs := FromSlice([]int{5, 4, 3, 2, 1})
	dropped := Drop(2, xs)
	ys := ToSlice(dropped)
	assert.Equal(t, []int{3, 2, 1}, ys)
}
