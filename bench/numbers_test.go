package bench_test

import (
	"fmt"
	"testing"

	"githug.com/dsolerh/go-anydict"
)

func Benchmark_Integer_downcast_check(b *testing.B) {
	var dict = anydict.Dict{
		"prop": int64(12),
	}
	directCast := func(d anydict.Dict, prop string) (int8, error) {
		if someval, exist := d[prop]; !exist {
			return 0, fmt.Errorf("not present")
		} else {
			switch val := someval.(type) {
			case int64:
				return int8(val), nil
			default:

				return 0, fmt.Errorf("prop %s is not of type %T", prop, val)
			}
		}
	}

	var _val int8
	b.Run("using Integer[int8]", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_val, _ = anydict.Integer[int8](dict, "prop")
		}
		someint = int(_val)
	})
	b.Run("using direct cast", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_val, _ = directCast(dict, "prop")
		}
		someint = int(_val)
	})
}
