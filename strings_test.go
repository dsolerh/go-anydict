package anydict

import (
	"testing"

	testutils "githug.com/dsolerh/go-anydict/test_utils"
)

func strExecutor[T strings](d Dict, prop string) (any, error) {
	val, err := String[T](d, prop)
	return any(val), err
}

func Test_String(t *testing.T) {
	type customBytes []byte
	dict := Dict{
		"str":      "the boy and the heron",
		"non_utf8": []byte{0xC3, 0x28},
		"invalid":  123,
	}
	testutils.RunTestCases(t, dict, []testutils.TestCase{
		{
			Desc:        "should return the value as a string",
			Prop:        "str",
			Executor:    strExecutor[string],
			ExpectedVal: dict["str"],
		},
		{
			Desc:        "should return the value as a []byte",
			Prop:        "str",
			Executor:    strExecutor[[]byte],
			ExpectedVal: []byte(dict["str"].(string)),
		},
		{
			Desc:        "should return the value as a custom []byte",
			Prop:        "str",
			Executor:    strExecutor[customBytes],
			ExpectedVal: customBytes(dict["str"].(string)),
		},
		{
			Desc:        "should return the value as a string (non utf8 valid)",
			Prop:        "non_utf8",
			Executor:    strExecutor[string],
			ExpectedVal: string(dict["non_utf8"].([]byte)),
		},
		{
			Desc:           "should return an error: not present",
			Prop:           "no-prop",
			Executor:       strExecutor[string],
			ExpectErr:      true,
			CheckErrorWith: testutils.IsErrorOfType[*PropNotPresentError],
		},
		{
			Desc:           "should return an error: invalid type",
			Prop:           "invalid",
			Executor:       strExecutor[string],
			ExpectErr:      true,
			CheckErrorWith: testutils.IsErrorOfType[*PropNotOfTypeError],
		},
	})
}

func strOrExecutor[T strings](defaultVal T) testutils.FuncExecutor {
	return func(m map[string]any, s string) (any, error) {
		val, err := StringOr(m, s, defaultVal)
		return any(val), err
	}
}

func Test_StringOr(t *testing.T) {
	type customBytes []byte
	dict := Dict{
		"str":      "the boy and the heron",
		"non_utf8": []byte{0xC3, 0x28},
		"invalid":  123,
	}
	testutils.RunTestCases(t, dict, []testutils.TestCase{
		{
			Desc:        "should return the value",
			Prop:        "str",
			Executor:    strOrExecutor("just in case"),
			ExpectedVal: dict["str"],
		},
		{
			Desc:        "should return the default value",
			Prop:        "no-prop",
			Executor:    strOrExecutor("just in case"),
			ExpectedVal: "just in case",
		},
	})
}
