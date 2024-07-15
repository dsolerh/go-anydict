package anydict

import (
	"math"
	"testing"

	testutils "githug.com/dsolerh/go-anydict/test_utils"
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
	testutils.RunTestCases(t, dict, []testutils.TestCase{
		{
			Desc:        "should return an int",
			Prop:        "int",
			Executor:    intExecutor[int],
			ExpectedVal: dict["int"],
		},
		{
			Desc:           "should return an error: not present",
			Prop:           "no-prop",
			Executor:       intExecutor[int],
			ExpectErr:      true,
			CheckErrorWith: testutils.IsErrorOfType[*PropNotPresentError],
		},
		{
			Desc:           "should return an error: invalid type",
			Prop:           "invalid",
			Executor:       intExecutor[int],
			ExpectErr:      true,
			CheckErrorWith: testutils.IsErrorOfType[*PropNotOfTypeError],
		},
		{
			Desc:           "should return an error: cannot downcast",
			Prop:           "big",
			Executor:       intExecutor[int8],
			ExpectErr:      true,
			CheckErrorWith: testutils.IsErrorOfType[*InvalidConversionError],
		},
		{
			Desc:        "should return a custom type",
			Prop:        "int",
			Executor:    intExecutor[myCustomInt],
			ExpectedVal: myCustomInt(23),
		},
	})
}

func intOrExecutor[T IntegerLike](defaultVal T) testutils.FuncExecutor {
	return func(m map[string]any, s string) (any, error) {
		val, err := IntegerOr(m, s, defaultVal)
		return any(val), err
	}
}

func Test_IntegerOr(t *testing.T) {
	type myCustomInt int
	dict := Dict{
		"int": int(23),
	}
	testutils.RunTestCases(t, dict, []testutils.TestCase{
		{
			Desc:        "should return an int",
			Prop:        "int",
			Executor:    intOrExecutor(123),
			ExpectedVal: dict["int"],
		},
		{
			Desc:        "should return the default value",
			Prop:        "no-prop",
			Executor:    intOrExecutor(9999),
			ExpectedVal: 9999,
		},
	})
}
