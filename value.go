package anydict

import "fmt"

// Value returns the value of a given 'prop' string in a given 'dict' Dict
// the Dict type is an alias over a map[string]any
// the main usage of this function is to get the value out of the dict already casted to the desired type
// Example:
//
//	dict := Dict{"prop": int(1)}
//	val, err := Value[int](dict, "prop")
//	if err != nil {
//		panic("there should not be an error")
//	}
//	if val != 1 {
//		panic("val should be equal 1")
//	}
//
// This function will return the value if the value exist and it's of the appropriate type (if the type is
// any it will return the value no matter the type, it's not the intended usage)
// The function will return an error if:
//   - the 'prop' is not defined (empty string).
//   - the value does not exist.
//   - the value cannot be represented by the expected return type
//
// For more info check the test cases at value_test.go
func Value[V any](dict Dict, prop string) (V, error) {
	if val, ok := dict[prop].(V); !ok {
		if prop == "" {
			return val, fmt.Errorf("invalid empty property")
		}
		if _, exist := dict[prop]; exist {
			return val, fmt.Errorf("prop %s is not present in the dict", prop)
		}
		return val, fmt.Errorf("prop %s is not of type %T", prop, val)
	} else {
		return val, nil
	}
}
