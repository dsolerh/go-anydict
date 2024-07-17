package bench_test

import (
	"fmt"
	"testing"
)

type simpleStruct struct {
	a any
	b []byte
	s string
}

var someArr []*simpleStruct

func with_capacity(n int) []*simpleStruct {
	arr := make([]*simpleStruct, 0, n)
	for i := 0; i < n; i++ {
		arr = append(arr, &simpleStruct{a: 1321, b: []byte{}, s: "simple"})
	}
	return arr
}

func default_capacity(n int) []*simpleStruct {
	arr := make([]*simpleStruct, 0)
	for i := 0; i < n; i++ {
		arr = append(arr, &simpleStruct{a: 1321, b: []byte{}, s: "simple"})
	}

	return arr
}

func with_size(n int) []*simpleStruct {
	arr := make([]*simpleStruct, n)
	for i := 0; i < n; i++ {
		arr[i] = &simpleStruct{a: 1321, b: []byte{}, s: "simple"}
	}

	return arr
}

func Benchmark_array_append(b *testing.B) {
	for n := 10; n < 1_000_001; n *= 10 {
		b.Run(fmt.Sprintf("with_cap_%d", n), func(b *testing.B) {
			arr := with_capacity(n)
			someArr = arr
		})
		b.Run(fmt.Sprintf("default_cap_%d", n), func(b *testing.B) {
			arr := default_capacity(n)
			someArr = arr
		})
		b.Run(fmt.Sprintf("with_size_%d", n), func(b *testing.B) {
			arr := with_size(n)
			someArr = arr
		})
	}
}
