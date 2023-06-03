package maybe

import "fmt"

type errInvalidUnwrap[A any] struct {
	value Maybe[A]
}

func (e errInvalidUnwrap[A]) Error() string {
	return fmt.Sprintf("unwrap value from maybe: %T", e.value)
}
