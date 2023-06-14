package iter

import (
	"prelude/maybe"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmpty(t *testing.T) {
	xs := Empty[int]()
	assert.Equal(t, maybe.Nothing[int](), xs())
}

func TestSingle(t *testing.T) {
	xs := Single(1)
	assert.Equal(t, maybe.Just(1), xs())
	assert.Equal(t, maybe.Nothing[int](), xs())
}
