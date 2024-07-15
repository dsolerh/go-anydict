package testutils

import (
	"errors"
	"reflect"
	"testing"
)

type FuncExecutor func(map[string]any, string) (any, error)

func IsErrorOfType[E error](err error) bool {
	var e E
	return errors.As(err, &e)
}

type TestCase struct {
	Desc           string
	Prop           string
	Executor       FuncExecutor
	ExpectedVal    any
	ExpectErr      bool
	CheckErrorWith func(err error) bool
}

func RunTestCases(t *testing.T, dict map[string]any, testCases []TestCase) {
	for _, tc := range testCases {
		t.Run(tc.Desc, func(t *testing.T) {
			val, err := tc.Executor(dict, tc.Prop)

			if tc.ExpectErr {
				// some error is expected
				if err == nil {
					t.Error("expected an error but got none")
				} else if tc.CheckErrorWith != nil && !tc.CheckErrorWith(err) {
					// the error should be of the expected type
					t.Errorf(
						"expected an error, but got: %v | val: %T(%v)",
						err,
						val,
						val,
					)
				}
			} else {
				// no error is expected
				if err != nil {
					t.Errorf("expected no error but got: %v", err)
				}
				if !reflect.DeepEqual(val, tc.ExpectedVal) {
					// there should be a value if there's no expected error
					t.Errorf(
						"expected val: %T(%v), but got: %T(%v)",
						tc.ExpectedVal,
						tc.ExpectedVal,
						val,
						val,
					)
				}
			}
		})
	}
}
