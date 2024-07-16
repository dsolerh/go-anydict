package anydict

type StructBuilder[T any] interface {
	FromDict(d Dict) (T, error)
}

func Struct[S StructBuilder[T], T any](dict Dict, prop string) (T, error) {
	val, err := Value[Dict](dict, prop)
	if err != nil {
		return zero[T](), err
	}

	var s S
	return s.FromDict(val)
}

type StructBuilderFn[T any] func(d Dict) (T, error)

func StructFn[T any](dict Dict, prop string, builder StructBuilderFn[T]) (T, error) {
	val, err := Value[Dict](dict, prop)
	if err != nil {
		return zero[T](), err
	}

	return builder(val)
}
