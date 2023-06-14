package iter

import "prelude/pair"

func Enumerate[A any](xs Iterator[A]) Iterator[pair.Pair[uint, A]] {
	indexes := Iterate(func(i uint) uint { return i + 1 }, uint(0))
	return Zip(indexes, xs)
}
