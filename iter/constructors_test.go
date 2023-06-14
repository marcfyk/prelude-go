package iter

import (
	"prelude/maybe"
	"prelude/pair"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFromSlice(t *testing.T) {
	xs := []int{3, 1, 2}
	ys := FromSlice(xs)

	assert.Equal(t, maybe.Just(3), ys())
	assert.Equal(t, maybe.Just(1), ys())
	assert.Equal(t, maybe.Just(2), ys())
	assert.Equal(t, maybe.Nothing[int](), ys())
}

func TestFromSliceCopy(t *testing.T) {
	xs := []int{3, 1, 2}
	ys := FromSlice(xs)
	xs = append(xs, 4)
	xs[0] = 5

	assert.Equal(t, maybe.Just(3), ys())
	assert.Equal(t, maybe.Just(1), ys())
	assert.Equal(t, maybe.Just(2), ys())
	assert.Equal(t, maybe.Nothing[int](), ys())
}

func TestToSlice(t *testing.T) {
	xs := FromSlice([]int{3, 1, 2})
	assert.Equal(t, []int{3, 1, 2}, ToSlice(xs))
}

func TestFromMap(t *testing.T) {
	xs := map[int]int{
		1: 10,
		2: 20,
		3: 30,
	}
	ys := FromMap(xs)
	elements := make(map[pair.Pair[int, int]]struct{})
	for i := 0; i < len(xs); i++ {
		p := ys()
		assert.True(t, maybe.IsJust(p))
		elements[maybe.Unwrap(p)] = struct{}{}
	}
	assert.Equal(t, maybe.Nothing[pair.Pair[int, int]](), ys())
	assert.Equal(t, len(xs), len(elements))
	_, ok := elements[pair.New(1, 10)]
	assert.True(t, ok)
	_, ok = elements[pair.New(2, 20)]
	assert.True(t, ok)
	_, ok = elements[pair.New(3, 30)]
	assert.True(t, ok)
}

func TestFromMapCopy(t *testing.T) {
	xs := map[int]int{
		1: 10,
		2: 20,
		3: 30,
	}
	ys := FromMap(xs)
	xs[4] = 40
	delete(xs, 2)
	elements := make(map[pair.Pair[int, int]]struct{})
	for i := 0; i < len(xs); i++ {
		p := ys()
		assert.True(t, maybe.IsJust(p))
		elements[maybe.Unwrap(p)] = struct{}{}
	}
	assert.Equal(t, maybe.Nothing[pair.Pair[int, int]](), ys())
	assert.Equal(t, len(xs), len(elements))
	_, ok := elements[pair.New(1, 10)]
	assert.True(t, ok)
	_, ok = elements[pair.New(2, 20)]
	assert.True(t, ok)
	_, ok = elements[pair.New(3, 30)]
	assert.True(t, ok)
}
