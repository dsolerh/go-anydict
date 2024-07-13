package anydict

import (
	"math"
	"reflect"
	"testing"
)

func intExecutor[T IntegerLike](d Dict, prop string) (any, error) {
	val, err := Integer[T](d, prop)
	return any(val), err
}

func Test_Integer(t *testing.T) {
	type myCustomInt int
	dict := Dict{
		"int":     int(23),
		"big":     math.MinInt64,
		"invalid": "invalid",
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
			desc:        "should return an int",
			prop:        "int",
			executor:    intExecutor[int],
			expectedVal: dict["int"],
		},
		{
			desc:      "should return an error: not present",
			prop:      "no-prop",
			executor:  intExecutor[int],
			expectErr: true,
		},
		{
			desc:      "should return an error: invalid type",
			prop:      "invalid",
			executor:  intExecutor[int],
			expectErr: true,
		},
		{
			desc:      "should return an error: cannot downcast",
			prop:      "big",
			executor:  intExecutor[int8],
			expectErr: true,
		},
		{
			desc:        "should return a custom type",
			prop:        "int",
			executor:    intExecutor[myCustomInt],
			expectedVal: myCustomInt(23),
		},
	}
	for _, tc := range testcases {
		t.Run(tc.desc, func(t *testing.T) {
			val, err := tc.executor(dict, tc.prop)
			didErr := err != nil

			if tc.expectErr != didErr {
				t.Errorf(
					"expected an error: %t, but got one: %t(%v) val: %T(%v)",
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
