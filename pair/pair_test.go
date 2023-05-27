package pair

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	left := 1
	right := "a"
	pair := New(left, right)
	assert.Equal(t, Pair[int, string]{Left: left, Right: right}, pair)
}

func TestLeft(t *testing.T) {
	left := 1
	right := "a"
	pair := New(left, right)
	assert.Equal(t, pair.Left, Left(pair))
}

func TestRight(t *testing.T) {
	left := 1
	right := "a"
	pair := New(left, right)
	assert.Equal(t, pair.Right, Right(pair))
}
