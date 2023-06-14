package iter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	f := func(n int) int { return n + 1 }
	xs := FromSlice([]int{3, 1, 2})
	mapped := Map(f, xs)
	ys := ToSlice(mapped)
	assert.Equal(t, []int{f(3), f(1), f(2)}, ys)
}
