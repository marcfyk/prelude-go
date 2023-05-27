package pair

type Pair[A, B any] struct {
	Left  A
	Right B
}

func New[A, B any](left A, right B) Pair[A, B] {
	return Pair[A, B]{
		Left:  left,
		Right: right,
	}
}

func Left[A, B any](pair Pair[A, B]) A {
	return pair.Left
}

func Right[A, B any](pair Pair[A, B]) B {
	return pair.Right
}
