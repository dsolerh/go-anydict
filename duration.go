package anydict

import "time"

func Duration(dict Dict, prop string) (time.Duration, error) {
	if someval, exist := dict[prop]; !exist {
		return 0, newPropNotPresentError(prop)
	} else {
		switch d := someval.(type) {
		case time.Duration:
			return d, nil
		case string:
			return time.ParseDuration(d)
		default:
			return 0, newPropNotOfTypeError(prop, time.Duration(0))
		}
	}
}

func DurationOr(dict Dict, prop string, defaultVal time.Duration) (time.Duration, error) {
	if someval, exist := dict[prop]; !exist {
		return defaultVal, nil
	} else {
		switch d := someval.(type) {
		case time.Duration:
			return d, nil
		case string:
			return time.ParseDuration(d)
		default:
			return 0, newPropNotOfTypeError(prop, time.Duration(0))
		}
	}
}
