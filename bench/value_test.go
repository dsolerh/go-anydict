package bench_test

import (
	"fmt"
	"testing"

	"githug.com/dsolerh/go-anydict"
)

var str string

func Benchmark_Value_string(b *testing.B) {
	var dict = anydict.Dict{
		"prop": "some very interesting string",
	}
	directCast := func(d anydict.Dict, prop string) (string, error) {
		str, ok := d[prop].(string)
		if ok {
			return str, nil
		}
		return "", fmt.Errorf("invalid string")
	}

	var _str string
	b.Run("using Value[string]", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_str, _ = anydict.Value[string](dict, "prop")
		}
		str = _str
	})
	b.Run("using direct cast", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_str, _ = directCast(dict, "prop")
		}
		str = _str
	})
}

var someint int

func Benchmark_Value_int(b *testing.B) {
	var dict = anydict.Dict{
		"prop": 32,
	}
	directCast := func(d anydict.Dict, prop string) (int, error) {
		str, ok := d[prop].(int)
		if ok {
			return str, nil
		}
		return 0, fmt.Errorf("invalid string")
	}

	var _int int
	b.Run("using Value[string]", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_int, _ = anydict.Value[int](dict, "prop")
		}
		someint = _int
	})
	b.Run("using direct cast", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_int, _ = directCast(dict, "prop")
		}
		someint = _int
	})
}
