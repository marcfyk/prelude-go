package iter

import (
	"prelude/pair"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestZipWith(t *testing.T) {
	f := func(x, y int) int { return x + y }
	xs := FromSlice([]int{1, 2, 3, 4, 5})
	ys := FromSlice([]int{10, 20, 30, 40, 50})
	zipped := ZipWith(f, xs, ys)
	assert.Equal(t, []int{f(1, 10), f(2, 20), f(3, 30), f(4, 40), f(5, 50)}, ToSlice(zipped))
}

func TestZipWithDiffLength(t *testing.T) {
	f := func(x, y int) int { return x + y }
	xs := FromSlice([]int{1, 2, 3})
	ys := FromSlice([]int{10, 20, 30, 40, 50})
	zipped := ZipWith(f, xs, ys)
	assert.Equal(t, []int{f(1, 10), f(2, 20), f(3, 30)}, ToSlice(zipped))

	xs = FromSlice([]int{1, 2, 3, 4, 5})
	ys = FromSlice([]int{10, 20, 30})
	zipped = ZipWith(f, xs, ys)
	assert.Equal(t, []int{f(1, 10), f(2, 20), f(3, 30)}, ToSlice(zipped))
}

func TestZip(t *testing.T) {
	xs := FromSlice([]int{1, 2, 3, 4, 5})
	ys := FromSlice([]int{10, 20, 30, 40, 50})
	zipped := Zip(xs, ys)
	assert.Equal(t, []pair.Pair[int, int]{
		pair.New(1, 10),
		pair.New(2, 20),
		pair.New(3, 30),
		pair.New(4, 40),
		pair.New(5, 50),
	}, ToSlice(zipped))
}

func TestZipDiffLength(t *testing.T) {
	xs := FromSlice([]int{1, 2, 3})
	ys := FromSlice([]int{10, 20, 30, 40, 50})
	zipped := Zip(xs, ys)
	assert.Equal(t, []pair.Pair[int, int]{
		pair.New(1, 10),
		pair.New(2, 20),
		pair.New(3, 30),
	}, ToSlice(zipped))

	xs = FromSlice([]int{1, 2, 3, 4, 5})
	ys = FromSlice([]int{10, 20, 30})
	zipped = Zip(xs, ys)
	assert.Equal(t, []pair.Pair[int, int]{
		pair.New(1, 10),
		pair.New(2, 20),
		pair.New(3, 30),
	}, ToSlice(zipped))
}
