package iter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	f := func(n int) bool { return n%2 == 0 }
	xs := FromSlice([]int{1, 2, 3, 4, 5})
	filtered := Filter(f, xs)
	ys := ToSlice(filtered)
	assert.Equal(t, []int{2, 4}, ys)
}
