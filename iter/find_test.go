package iter

import (
	"prelude/maybe"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindByFound(t *testing.T) {
	f := func(n int) bool { return n%2 == 0 }
	xs := FromSlice([]int{1, 2, 3, 4, 5})
	x := FindBy(f, xs)
	assert.Equal(t, maybe.Just(2), x)

}

func TestFindByNotFound(t *testing.T) {
	g := func(n int) bool { return n == 0 }
	xs := FromSlice([]int{1, 2, 3, 4, 5})
	x := FindBy(g, xs)
	assert.Equal(t, maybe.Nothing[int](), x)
}

func TestFindFound(t *testing.T) {
	xs := FromSlice([]int{1, 2, 3})
	x := Find(2, xs)
	assert.Equal(t, maybe.Just(2), x)
}

func TestFindNotFound(t *testing.T) {
	xs := FromSlice([]int{1, 2, 3})
	x := Find(0, xs)
	assert.Equal(t, maybe.Nothing[int](), x)
}
