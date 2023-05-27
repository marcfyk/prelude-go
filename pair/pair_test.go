package pair

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	left := 1
	right := "a"
	pair := New(left, right)
	assert.Equal(t, pair, Pair[int, string]{Left: left, Right: right})
}

func TestLeft(t *testing.T) {
	left := 1
	right := "a"
	pair := New(left, right)
	assert.Equal(t, Left(pair), pair.Left)
}

func TestRight(t *testing.T) {
	left := 1
	right := "a"
	pair := New(left, right)
	assert.Equal(t, Right(pair), pair.Right)
}
