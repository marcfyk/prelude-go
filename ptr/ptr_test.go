package ptr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRef(t *testing.T) {
	value := 1
	assert.Equal(t, &value, Ref(value))
}

func TestDeref(t *testing.T) {
	value := 1
	pointer := &value
	assert.Equal(t, value, Deref(pointer))
}

func TestNil(t *testing.T) {
	assert.Nil(t, Nil[int]())
}

func TestIsNil(t *testing.T) {
	assert.True(t, IsNil[int](nil))
	assert.False(t, IsNil(Ref(1)))
}

func TestNotNil(t *testing.T) {
	assert.True(t, NotNil(Ref(1)))
	assert.False(t, NotNil[int](nil))
}
