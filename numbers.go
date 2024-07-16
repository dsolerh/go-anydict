package anydict

type integers interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// IntegerSafeConvert tries to convert val from the ISrc type to IDst (both `integers`).
// If the value can be represented with no loss the val in IDst type is returned,
// otherwise an error (*InvalidConversionError) is returned.
func IntegerSafeConvert[ISrc, IDst integers](val ISrc) (IDst, error) {
	val2 := IDst(val)
	if ISrc(val2) != val {
		return IDst(0), newInvalidConversionError(val, val2)
	}
	return val2, nil
}

func toInteger[I integers](someval any, prop string) (I, error) {
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

// Integer tries to get prop from dict.
//   - If prop is not present an error (*PropNotPresentError) is returned.
//   - If prop cannot be represented as an integer (~int | ~int8 | ~int16 | ~int32 | ~int64) an error (*InvalidConversionError) is returned.
//   - Otherwise the value is returned with the proper type.
func Integer[I integers](dict Dict, prop string) (I, error) {
	if someval, exist := dict[prop]; !exist {
		return 0, newPropNotPresentError(prop)
	} else {
		return toInteger[I](someval, prop)
	}
}

// IntegerOr works in the same way as Integer, but if the prop is not present
// it'll return defaultVal.
func IntegerOr[I integers](dict Dict, prop string, defaultVal I) (I, error) {
	if someval, exist := dict[prop]; !exist {
		return defaultVal, nil
	} else {
		return toInteger[I](someval, prop)
	}
}

type floats interface {
	~float32 | ~float64
}

// FloatSafeConvert tries to convert val from the ISrc type to IDst (both `integers`).
// If the value can be represented with no loss the val in IDst type is returned,
// otherwise an error (*InvalidConversionError) is returned.
func FloatSafeConvert[FSrc, FDst floats](val FSrc) (FDst, error) {
	val2 := FDst(val)
	if FSrc(val2) != val {
		return FDst(0), newInvalidConversionError(val, val2)
	}
	return val2, nil
}

func toFloat[F floats](someval any, prop string) (F, error) {
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

// Float tries to get prop from dict.
//   - If prop is not present an error (*PropNotPresentError) is returned.
//   - If prop cannot be represented as an integer (~int | ~int8 | ~int16 | ~int32 | ~int64) an error (*InvalidConversionError) is returned.
//   - Otherwise the value is returned with the proper type.
func Float[F floats](dict Dict, prop string) (F, error) {
	if someval, exist := dict[prop]; !exist {
		return 0, newPropNotPresentError(prop)
	} else {
		return toFloat[F](someval, prop)
	}
}

// FloatOr works in the same way as Float, but if the prop is not present
// it'll return defaultVal.
func FloatOr[F floats](dict Dict, prop string, defaultVal F) (F, error) {
	if someval, exist := dict[prop]; !exist {
		return defaultVal, nil
	} else {
		return toFloat[F](someval, prop)
	}
}
