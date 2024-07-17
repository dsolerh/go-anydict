package anydict

import (
	"testing"

	testutils "githug.com/dsolerh/go-anydict/test_utils"
)

func arrExecutor[T any](d Dict, prop string) (any, error) {
	val, err := Array[T](d, prop)
	return any(val), err
}

func Test_Array(t *testing.T) {
	dict := Dict{
		"arr1":    []int{1, 2},
		"arr2":    []any{'a', 'b'},
		"arr3":    []any{"1", 2},
		"not-arr": 12,
	}
	testutils.RunTestCases(t, dict, []testutils.TestCase{
		{
			Desc:           "should get an error: prop no present",
			Prop:           "no-prop",
			Executor:       arrExecutor[int],
			ExpectErr:      true,
			CheckErrorWith: testutils.IsErrorOfType[*PropNotPresentError],
		},
		{
			Desc:        "should get a []int (from []int type)",
			Prop:        "arr1",
			Executor:    arrExecutor[int],
			ExpectedVal: dict["arr1"],
		},
		{
			Desc:        "should get a []rune (from []any type)",
			Prop:        "arr2",
			Executor:    arrExecutor[rune],
			ExpectedVal: []rune{'a', 'b'},
		},
		{
			Desc:           "should get an error: invalid type at index 0",
			Prop:           "arr3",
			Executor:       arrExecutor[int],
			ExpectErr:      true,
			CheckErrorWith: testutils.IsErrorOfType[*InvalidTypeAtIndexError],
		},
		{
			Desc:           "should get an error: prop not of type",
			Prop:           "arr3",
			Executor:       arrExecutor[int],
			ExpectErr:      true,
			CheckErrorWith: testutils.IsErrorOfType[*PropNotOfTypeError],
		},
	},
	)
}

func arrOrExecutor[T any](defaultVal []T) testutils.FuncExecutor {
	return func(m map[string]any, s string) (any, error) {
		val, err := ArrayOr(m, s, defaultVal)
		return any(val), err
	}
}

func Test_ArrayOr(t *testing.T) {
	dict := Dict{
		"arr1": []any{true, false},
	}
	testutils.RunTestCases(t, dict, []testutils.TestCase{
		{
			Desc:        "should return the default value",
			Prop:        "no-prop",
			Executor:    arrOrExecutor([]int{1, 2}),
			ExpectedVal: []int{1, 2},
		},
		{
			Desc:        "should return an []bool",
			Prop:        "arr1",
			Executor:    arrOrExecutor([]bool{false}),
			ExpectedVal: []bool{true, false},
		},
	})
}
