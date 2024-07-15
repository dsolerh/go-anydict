package anydict

import (
	"testing"

	testutils "githug.com/dsolerh/go-anydict/test_utils"
)

func valExecutor[T any](d Dict, prop string) (any, error) {
	val, err := Value[T](d, prop)
	return any(val), err
}

func Test_Value(t *testing.T) {
	dict := Dict{
		"int": int(23),
	}
	testutils.RunTestCases(t, dict, []testutils.TestCase{
		{
			Desc:           "should get an error: invalid prop type (string -> int)",
			Prop:           "int",
			Executor:       valExecutor[string],
			ExpectErr:      true,
			CheckErrorWith: testutils.IsErrorOfType[*PropNotOfTypeError],
		},
		{
			Desc:           "should get an error: prop not present",
			Prop:           "no-prop",
			Executor:       valExecutor[int],
			ExpectErr:      true,
			CheckErrorWith: testutils.IsErrorOfType[*PropNotPresentError],
		},
		{
			Desc:        "should get the prop as any",
			Prop:        "int",
			Executor:    Value[any],
			ExpectedVal: dict["int"],
		},
		{
			Desc:        "should get an int",
			Prop:        "int",
			Executor:    valExecutor[int],
			ExpectedVal: dict["int"],
		},
	},
	)
}

func valOrExecutor[T any](defaultVal T) testutils.FuncExecutor {
	return func(m map[string]any, s string) (any, error) {
		val, err := ValueOr(m, s, defaultVal)
		return any(val), err
	}
}

func Test_ValueOr(t *testing.T) {
	dict := Dict{
		"str": "some interesting quote",
	}
	testutils.RunTestCases(t, dict, []testutils.TestCase{
		{
			Desc:        "should return the value from the dict",
			Prop:        "str",
			Executor:    valOrExecutor("something"),
			ExpectedVal: dict["str"],
		},
		{
			Desc:        "should return the default value",
			Prop:        "not-key",
			Executor:    valOrExecutor("something"),
			ExpectedVal: "something",
		},
		{
			Desc:           "should return an error: invalid type",
			Prop:           "str",
			Executor:       valOrExecutor(19),
			ExpectErr:      true,
			CheckErrorWith: testutils.IsErrorOfType[*PropNotOfTypeError],
		},
	})
}
