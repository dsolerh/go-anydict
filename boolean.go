package anydict

// Boolean tries to get prop(string) from `dict`(Dict) and return it's
// value as `bool` otherwise an error.
//   - If prop is not present returns error *PropNotPresentError.
//   - If prop is not boll returns error *PropNotOfTypeError.
//   - Otherwise returns the value of prop as boll
func Boolean(dict Dict, prop string) (bool, error) {
	if someval, exist := dict[prop]; !exist {
		return false, newPropNotPresentError(prop)
	} else if val, ok := someval.(bool); ok {
		return val, nil
	} else {
		return false, newPropNotOfTypeError(prop, false)
	}
}

// BooleanOr works like Boolean but if the prop is not present it returns defaultVal
func BooleanOr(dict Dict, prop string, defaultVal bool) (bool, error) {
	if someval, exist := dict[prop]; !exist {
		return defaultVal, nil
	} else if val, ok := someval.(bool); ok {
		return val, nil
	} else {
		return false, newPropNotOfTypeError(prop, false)
	}
}
