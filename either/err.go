package either

import "fmt"

type errInvalidUnwrapLeft[A, B any] struct {
	value Either[A, B]
}

func (e errInvalidUnwrapLeft[A, B]) Error() string {
	return fmt.Sprintf("unwrap left value from either: %T", e.value)
}

type errInvalidUnwrapRight[A, B any] struct {
	value Either[A, B]
}

func (e errInvalidUnwrapRight[A, B]) Error() string {
	return fmt.Sprintf("unwrap right value from either: %T", e.value)
}

type errInvalidEither[A, B any] struct {
	value Either[A, B]
}

func (e errInvalidEither[A, B]) Error() string {
	return fmt.Sprintf("invalid either: %+v", e.value)
}
