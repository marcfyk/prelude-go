package iter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConcat(t *testing.T) {
	xs := FromSlice([]int{1, 2, 3})
	ys := FromSlice([]int{4, 5})
	zs := Concat(xs, ys)
	assert.Equal(t, []int{1, 2, 3, 4, 5}, ToSlice(zs))
}

func TestConcatEmpty(t *testing.T) {
	xs := Concat(FromSlice([]int{1, 2, 3}), Empty[int]())
	assert.Equal(t, []int{1, 2, 3}, ToSlice(xs))
	ys := Concat(Empty[int](), FromSlice([]int{4, 5}))
	assert.Equal(t, []int{4, 5}, ToSlice(ys))
	zs := Concat(Empty[int](), Empty[int]())
	assert.Equal(t, []int{}, ToSlice(zs))
}
