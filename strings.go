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

func String[S strings](dict Dict, prop string) (S, error) {
	if someval, exist := dict[prop]; !exist {
		return zero[S](), newPropNotPresentError(prop)
	} else {
		return toString[S](someval, prop)
	}
}

func StringOr[S strings](dict Dict, prop string, defaultVal S) (S, error) {
	if someval, exist := dict[prop]; !exist {
		return defaultVal, nil
	} else {
		return toString[S](someval, prop)
	}
}
