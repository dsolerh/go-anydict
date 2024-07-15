package anydict

type IntegerLike interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

func IntegerSafeConvert[ISrc, IDst IntegerLike](val ISrc) (IDst, error) {
	val2 := IDst(val)
	if ISrc(val2) != val {
		return IDst(0), newInvalidConversionError(val, val2)
	}
	return val2, nil
}

func toInteger[I IntegerLike](someval any, prop string) (I, error) {
	switch val := someval.(type) {
	case I:
		return val, nil
	case int:
		return IntegerSafeConvert[int, I](val)
	case int8:
		return IntegerSafeConvert[int8, I](val)
	case int16:
		return IntegerSafeConvert[int16, I](val)
	case int32:
		return IntegerSafeConvert[int32, I](val)
	case int64:
		return IntegerSafeConvert[int64, I](val)
	default:
		var zero I
		return zero, newPropNotOfTypeError(prop, zero)
	}
}

func Integer[I IntegerLike](dict Dict, prop string) (I, error) {
	if someval, exist := dict[prop]; !exist {
		return 0, newPropNotPresentError(prop)
	} else {
		return toInteger[I](someval, prop)
	}
}

func IntegerOr[I IntegerLike](dict Dict, prop string, defaultVal I) (I, error) {
	if someval, exist := dict[prop]; !exist {
		return defaultVal, nil
	} else {
		return toInteger[I](someval, prop)
	}
}

type FloatLike interface {
	~float32 | ~float64
}

func FloatSafeConvert[FSrc, FDst FloatLike](val FSrc) (FDst, error) {
	val2 := FDst(val)
	if FSrc(val2) != val {
		return FDst(0), newInvalidConversionError(val, val2)
	}
	return val2, nil
}

func toFloat[F FloatLike](someval any, prop string) (F, error) {
	switch val := someval.(type) {
	case F:
		return val, nil
	case float32:
		return FloatSafeConvert[float32, F](val)
	case float64:
		return FloatSafeConvert[float64, F](val)
	default:
		var zero F
		return zero, newPropNotOfTypeError(prop, zero)
	}
}

func Float[F FloatLike](dict Dict, prop string) (F, error) {
	if someval, exist := dict[prop]; !exist {
		return 0, newPropNotPresentError(prop)
	} else {
		return toFloat[F](someval, prop)
	}
}

func FloatOr[F FloatLike](dict Dict, prop string, defaultVal F) (F, error) {
	if someval, exist := dict[prop]; !exist {
		return defaultVal, nil
	} else {
		return toFloat[F](someval, prop)
	}
}
