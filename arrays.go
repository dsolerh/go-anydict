package anydict

func safeCast[T any](val any) (T, bool) {
	tval, ok := val.(T)
	return tval, ok
}

func transformArr[T any](arr []any) ([]T, error) {
	rval := make([]T, 0)
	for i := 0; i < len(arr); i++ {
		v, ok := safeCast[T](arr[i])
		if !ok {
			return nil, newInvalidTypeAtIndex(i, v)
		}
		rval = append(rval, v)
	}
	return rval, nil
}

func toArray[T any](someval any, prop string) ([]T, error) {
	switch arr := someval.(type) {
	case []T:
		return arr, nil
	case []any:
		return transformArr[T](arr)
	default:
		return nil, newPropNotOfTypeError(prop, ([]T)(nil))
	}
}

func Array[T any](dict Dict, prop string) ([]T, error) {
	if someval, exist := dict[prop]; !exist {
		return nil, newPropNotPresentError(prop)
	} else {
		return toArray[T](someval, prop)
	}
}

func ArrayOr[T any](dict Dict, prop string, defaultVar []T) ([]T, error) {
	if someval, exist := dict[prop]; !exist {
		return defaultVar, nil
	} else {
		return toArray[T](someval, prop)
	}
}
