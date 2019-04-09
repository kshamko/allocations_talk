package example4

import (
	"testing"
)

type x struct {
	data string
}

func BenchmarkMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		c := new(x)
		m := make(map[string]*x, 0)
		m["foo"] = c
	}
}

func BenchmarkSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		c := new(x)
		s := make([]*x, 1)
		s[0] = c
	}
}
