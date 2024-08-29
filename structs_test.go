package anydict

import (
	"testing"

	testutils "githug.com/dsolerh/go-anydict/test_utils"
)

type intAndFloat struct {
	i int
	f float32
}

var _ StructBuilder[*intAndFloat] = (*intAndFloat)(nil)

func (iaf *intAndFloat) FromDict(d Dict) (*intAndFloat, error) {
	var err error
	iaf.i, err = Integer[int](d, "int")
	if err != nil {
		return nil, err
	}
	iaf.f, err = Float[float32](d, "float")
	if err != nil {
		return nil, err
	}

	return iaf, nil
}

func structExecutor(d Dict, prop string) (any, error) {
	val, err := Struct[*intAndFloat](d, prop)
	return any(val), err
}

func structFnExecutor(d Dict, prop string) (any, error) {
	builder := &intAndFloat{}
	val, err := StructFn(d, prop, builder.FromDict)
	return any(val), err
}

func Test_Struct(t *testing.T) {
	dict := Dict{
		"prop": Dict{
			"int":   12,
			"float": 12.8,
		},
	}
	testutils.RunTestCases(t, dict, []testutils.TestCase{
		{
			Desc:           "should get an error: prop no present",
			Prop:           "no-prop",
			Executor:       structExecutor,
			ExpectErr:      true,
			CheckErrorWith: testutils.IsErrorOfType[*PropNotPresentError],
		},
		{
			Desc:        "should get the right struct",
			Prop:        "prop",
			Executor:    structExecutor,
			ExpectedVal: &intAndFloat{i: 12, f: 12.8},
		},
	},
	)
}
