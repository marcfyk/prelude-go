package iter

import "prelude/maybe"

type Iterator[A any] interface {
	Next() maybe.Maybe[A]
}
