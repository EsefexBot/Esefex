package opt

type Option[T any] struct {
	some  bool
	value T
}

func Some[T any](value T) Option[T] {
	return Option[T]{true, value}
}

func None[T any]() Option[T] {
	return Option[T]{some: false}
}

func (o Option[T]) IsSome() bool {
	return o.some
}

func (o Option[T]) IsNone() bool {
	return !o.some
}

func (o Option[T]) Unwrap() T {
	if !o.some {
		panic("Unwrap called on None")
	}

	return o.value
}

func (o Option[T]) UnwrapOr(def T) T {
	if o.some {
		return o.value
	}
	return def
}

func (o Option[T]) UnwrapOrElse(f func() T) T {
	if o.some {
		return o.value
	}
	return f()
}

func (o Option[T]) Expect(msg string) T {
	if o.some {
		return o.value
	}
	panic(msg)
}
