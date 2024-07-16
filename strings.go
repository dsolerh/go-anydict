package anydict

type strings interface {
	~string | ~[]byte
}

func toString[S strings](someval any, prop string) (S, error) {
	switch val := someval.(type) {
	case S:
		return val, nil
	case string:
		return S(val), nil
	case []byte:
		return S(val), nil
	default:
		var zero S
		return zero, newPropNotOfTypeError(prop, zero)
	}
}

// String returns an S (~string | ~[]byte) or an error.
//   - If prop is not found in the dict an error (*PropNotPresentError)
//   - If prop is not of a type compatible with ~string | ~[]byte it'll return an error (*PropNotOfTypeError)
//   - Otherwise it'll return the value in the specified type.
func String[S strings](dict Dict, prop string) (S, error) {
	if someval, exist := dict[prop]; !exist {
		return zero[S](), newPropNotPresentError(prop)
	} else {
		return toString[S](someval, prop)
	}
}

// StringOr works in the same way as String, but if the prop is not present
// it'll return defaultVal.
func StringOr[S strings](dict Dict, prop string, defaultVal S) (S, error) {
	if someval, exist := dict[prop]; !exist {
		return defaultVal, nil
	} else {
		return toString[S](someval, prop)
	}
}
