package anydict

type IntegerLike interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

func Integer[I IntegerLike](dict Dict, prop string) (I, error) {
	return zero[I](), nil
}

func IntegerOr[I IntegerLike](dict Dict, prop string, defaultVal I) (I, error) {
	return defaultVal, nil
}

type FloatLike interface {
	~float32 | ~float64
}

func Float[F FloatLike](dict Dict, prop string) (F, error) {
	return zero[F](), nil
}

func FloatOr[F FloatLike](dict Dict, prop string, defaultVal F) (F, error) {
	return defaultVal, nil
}
