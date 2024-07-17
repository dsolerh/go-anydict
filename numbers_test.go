package anydict

import (
	"math"
	"testing"

	testutils "githug.com/dsolerh/go-anydict/test_utils"
)

func intExecutor[T integers](d Dict, prop string) (any, error) {
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

func intOrExecutor[T integers](defaultVal T) testutils.FuncExecutor {
	return func(m map[string]any, s string) (any, error) {
		val, err := IntegerOr(m, s, defaultVal)
		return any(val), err
	}
}

func Test_IntegerOr(t *testing.T) {
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

func floatExecutor[T floats](d Dict, prop string) (any, error) {
	val, err := Float[T](d, prop)
	return any(val), err
}

func Test_Float(t *testing.T) {
	dict := Dict{
		"f32":     float32(32.0),
		"f64":     float64(64.0),
		"int":     123,
		"f64_max": math.MaxFloat64,
	}
	testutils.RunTestCases(t, dict, []testutils.TestCase{
		{
			Desc:        "should return a float64 value",
			Prop:        "f32",
			Executor:    floatExecutor[float64],
			ExpectedVal: float64(32.0),
		},
		{
			Desc:        "should return a float32 value",
			Prop:        "f64",
			Executor:    floatExecutor[float32],
			ExpectedVal: float32(64.0),
		},
		{
			Desc:           "should return an error: prop is not present",
			Prop:           "no-prop",
			Executor:       floatExecutor[float32],
			ExpectErr:      true,
			CheckErrorWith: testutils.IsErrorOfType[*PropNotPresentError],
		},
		{
			Desc:           "should return an error: prop of invalid type",
			Prop:           "int",
			Executor:       floatExecutor[float32],
			ExpectErr:      true,
			CheckErrorWith: testutils.IsErrorOfType[*PropNotOfTypeError],
		},
		{
			Desc:           "should return an error: invalid downcast",
			Prop:           "f64_max",
			Executor:       floatExecutor[float32],
			ExpectErr:      true,
			CheckErrorWith: testutils.IsErrorOfType[*InvalidConversionError],
		},
	})
}

func floatOrExecutor[T floats](defaultVal T) testutils.FuncExecutor {
	return func(m map[string]any, s string) (any, error) {
		val, err := FloatOr(m, s, defaultVal)
		return any(val), err
	}
}

func Test_FloatOr(t *testing.T) {
	dict := Dict{
		"f32": float32(32.0),
	}
	testutils.RunTestCases(t, dict, []testutils.TestCase{
		{
			Desc:        "should return the value",
			Prop:        "f32",
			Executor:    floatOrExecutor(0.0),
			ExpectedVal: 32.0,
		},
		{
			Desc:        "should return the default value",
			Prop:        "no-prop",
			Executor:    floatOrExecutor(float64(64.0)),
			ExpectedVal: float64(64.0),
		},
	})
}
