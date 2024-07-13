package anydict

func zero[T any]() T { return *new(T) }

func iszero[T comparable](v T) bool { return v == zero[T]() }
