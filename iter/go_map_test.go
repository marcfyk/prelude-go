package iter

import (
	"prelude/maybe"
	"prelude/pair"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFromMap(t *testing.T) {
	xs := map[int]int{
		1: 10,
		2: 20,
		3: 30,
	}
	ys := FromMap(xs)
	elements := make(map[pair.Pair[int, int]]struct{})
	for i := 0; i < len(xs); i++ {
		p := ys.Next()
		assert.True(t, maybe.IsJust(p))
		elements[maybe.Unwrap(p)] = struct{}{}
	}
	assert.Equal(t, maybe.Nothing[pair.Pair[int, int]](), ys.Next())
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
		p := ys.Next()
		assert.True(t, maybe.IsJust(p))
		elements[maybe.Unwrap(p)] = struct{}{}
	}
	assert.Equal(t, maybe.Nothing[pair.Pair[int, int]](), ys.Next())
	assert.Equal(t, len(xs), len(elements))
	_, ok := elements[pair.New(1, 10)]
	assert.True(t, ok)
	_, ok = elements[pair.New(2, 20)]
	assert.True(t, ok)
	_, ok = elements[pair.New(3, 30)]
	assert.True(t, ok)

}
