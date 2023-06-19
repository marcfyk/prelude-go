package iter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainsByFound(t *testing.T) {
	f := func(n int) bool { return n%2 == 0 }
	xs := FromSlice([]int{1, 2, 3, 4, 5})
	assert.True(t, ContainsBy(f, xs))
}

func TestContainsByNotFound(t *testing.T) {
	g := func(n int) bool { return n == 0 }
	xs := FromSlice([]int{1, 2, 3, 4, 5})
	assert.False(t, ContainsBy(g, xs))
}

func TestContainsFound(t *testing.T) {
	xs := FromSlice([]int{1, 2, 3})
	assert.True(t, Contains(2, xs))
}

func TestContainsNotFound(t *testing.T) {
	xs := FromSlice([]int{1, 2, 3})
	assert.False(t, Contains(5, xs))
}
