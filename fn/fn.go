package fn

func Id[A any](value A) A {
	return value
}

func Thunk[A any](value A) func() A {
	return func() A {
		return value
	}
}

func Flip[A, B, C any](f func(A) func(B) C) func(B) func(A) C {
	return func(b B) func(A) C {
		return func(a A) C {
			return f(a)(b)
		}
	}
}

func Curry[A, B, C any](f func(A, B) C) func(A) func(B) C {
	return func(a A) func(B) C {
		return func(b B) C {
			return f(a, b)
		}
	}
}

func Uncurry[A, B, C any](f func(A) func(B) C) func(A, B) C {
	return func(a A, b B) C {
		return f(a)(b)
	}
}

func Fmap[A, B, C any](g func(B) C, f func(A) B) func(A) C {
	return func(a A) C {
		return g(f(a))
	}
}

func Apply[A, B, C any](g func(A) func(B) C, f func(A) B) func(A) C {
	return func(a A) C {
		return g(a)(f(a))
	}
}

func Join[A, B any](f func(A) func(A) B) func(A) B {
	return func(a A) B {
		return f(a)(a)
	}
}

func Bind[A, B, C any](g func(B) func(A) C, f func(A) B) func(a A) C {
	return Join(Fmap(g, f))
}

func Compose[A, B, C any](g func(B) C, f func(A) B) func(A) C {
	return Fmap(g, f)
}

func Const[A, B any](a A) func(b B) A {
	return func(_ B) A {
		return a
	}
}
