package anydict

import (
	"reflect"
	"testing"
)

func valExecutor[T any](d Dict, prop string) (any, error) {
	val, err := Value[T](d, prop)
	return any(val), err
}

func Test_Value(t *testing.T) {
	dict := Dict{
		"int":  int(23),
		"i32":  int32(32),
		"i64":  int64(64),
		"str":  "some interesting quote",
		"f32":  float32(32.1),
		"f64":  float64(64.1),
		"bool": true,
	}

	type TestCase struct {
		desc        string
		prop        string
		executor    func(Dict, string) (any, error)
		expectedVal any
		expectErr   bool
	}
	testcases := []TestCase{
		{
			desc:        "should get an error: invalid empty prop",
			prop:        "",
			executor:    Value[any],
			expectedVal: nil,
			expectErr:   true,
		},
		{
			desc:        "should get an error: invalid prop type (string -> int)",
			prop:        "int",
			executor:    valExecutor[string],
			expectedVal: nil,
			expectErr:   true,
		},
		{
			desc:        "should get an error: invalid prop type (int64 -> int)",
			prop:        "int",
			executor:    valExecutor[int64],
			expectedVal: nil,
			expectErr:   true,
		},
		{
			desc:        "should get an error: invalid prop type (int8 -> int64)",
			prop:        "i64",
			executor:    valExecutor[int8],
			expectedVal: nil,
			expectErr:   true,
		},
		{
			desc:        "should get the prop as any",
			prop:        "int",
			executor:    Value[any],
			expectedVal: dict["int"],
			expectErr:   false,
		},
		{
			desc:        "should get an int",
			prop:        "int",
			executor:    valExecutor[int],
			expectedVal: dict["int"],
			expectErr:   false,
		},
		{
			desc:        "should get an int32",
			prop:        "i32",
			executor:    valExecutor[int32],
			expectedVal: dict["i32"],
			expectErr:   false,
		},
		{
			desc:        "should get an int64",
			prop:        "i64",
			executor:    valExecutor[int64],
			expectedVal: dict["i64"],
			expectErr:   false,
		},
		{
			desc:        "should get a string",
			prop:        "str",
			executor:    valExecutor[string],
			expectedVal: dict["str"],
			expectErr:   false,
		},
		{
			desc:        "should get a float32",
			prop:        "f32",
			executor:    valExecutor[float32],
			expectedVal: dict["f32"],
			expectErr:   false,
		},
		{
			desc:        "should get an int",
			prop:        "f64",
			executor:    valExecutor[float64],
			expectedVal: dict["f64"],
			expectErr:   false,
		},
		{
			desc:        "should get a bool (true)",
			prop:        "bool",
			executor:    valExecutor[bool],
			expectedVal: dict["bool"],
			expectErr:   false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.desc, func(t *testing.T) {
			val, err := tc.executor(dict, tc.prop)
			didErr := err != nil

			if tc.expectErr != didErr {
				t.Errorf(
					"expected an error: %v, but got one: %v(%v) val: %T(%v)",
					tc.expectErr,
					didErr,
					err,
					val,
					val,
				)
			} else if !tc.expectErr && !reflect.DeepEqual(val, tc.expectedVal) {
				t.Errorf(
					"expected val: %T(%v), but got: %T(%v)",
					tc.expectedVal,
					tc.expectedVal,
					val,
					val,
				)
			}
		})
	}
}
