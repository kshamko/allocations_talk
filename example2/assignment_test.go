package example2

import "testing"

type X struct {
	p *int
}

func BenchmarkAssignmentIndirect(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var i int
		x := &X{}
		x.p = &i // BAD: Cause of i escape
	}
}

func BenchmarkAssignmentDirect(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var i int
		x := &X{
			p: &i, // GOOD: i does not escape
		}
		_ = x
	}
}
