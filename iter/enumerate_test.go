package iter

import (
	"prelude/pair"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnumerate(t *testing.T) {
	xs := FromSlice([]int{5, 4, 3, 2, 1})
	enumerated := Enumerate(xs)
	ys := ToSlice(enumerated)
	assert.Equal(t, []pair.Pair[uint, int]{
		{Left: 0, Right: 5},
		{Left: 1, Right: 4},
		{Left: 2, Right: 3},
		{Left: 3, Right: 2},
		{Left: 4, Right: 1},
	}, ys)
}
