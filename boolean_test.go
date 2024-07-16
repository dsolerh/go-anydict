package anydict

import (
	"testing"

	testutils "githug.com/dsolerh/go-anydict/test_utils"
)

func boolExecutor(d Dict, prop string) (any, error) {
	return Boolean(d, prop)
}

func Test_Boolean(t *testing.T) {
	dict := Dict{
		"bool":    true,
		"invalid": 123,
	}
	testutils.RunTestCases(t, dict, []testutils.TestCase{
		{
			Desc:        "should return the value",
			Prop:        "bool",
			Executor:    boolExecutor,
			ExpectedVal: dict["bool"],
		},
		{
			Desc:           "should return an error: not present",
			Prop:           "no-prop",
			Executor:       boolExecutor,
			ExpectErr:      true,
			CheckErrorWith: testutils.IsErrorOfType[*PropNotPresentError],
		},
		{
			Desc:           "should return an error: invalid type",
			Prop:           "invalid",
			Executor:       boolExecutor,
			ExpectErr:      true,
			CheckErrorWith: testutils.IsErrorOfType[*PropNotOfTypeError],
		},
	})
}

func boolOrExecutor(defaultVal bool) testutils.FuncExecutor {
	return func(m map[string]any, s string) (any, error) {
		val, err := BooleanOr(m, s, defaultVal)
		return any(val), err
	}
}

func Test_BooleanOr(t *testing.T) {
	dict := Dict{
		"bool": true,
	}
	testutils.RunTestCases(t, dict, []testutils.TestCase{
		{
			Desc:        "should return the value",
			Prop:        "bool",
			Executor:    boolOrExecutor(false),
			ExpectedVal: dict["bool"],
		},
		{
			Desc:        "should return the default value",
			Prop:        "no-prop",
			Executor:    boolOrExecutor(false),
			ExpectedVal: false,
		},
	})
}
