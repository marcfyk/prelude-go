package iter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIterate(t *testing.T) {
	f := func(n int) int { return n * 2 }
	xs := Iterate(f, 1)
	firstTen := Take(10, xs)
	ys := ToSlice(firstTen)
	assert.Equal(t, []int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512}, ys)
}
