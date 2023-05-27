package ptr

func Ref[A any](value A) *A {
	return &value
}

func Deref[A any](pointer *A) A {
	return *pointer
}

func Nil[A any]() *A {
	return nil
}

func IsNil[A any](pointer *A) bool {
	return pointer == nil
}

func NotNil[A any](pointer *A) bool {
	return pointer != nil
}
