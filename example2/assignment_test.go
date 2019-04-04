package example2

import "testing"

func BenchmarkAssignmentIndirect(b *testing.B) {
	type X struct {
		p *int
	}
	for i := 0; i < b.N; i++ {
		var i int
		x := &X{}
		x.p = &i // BAD: Cause of i escape
	}
}

func BenchmarkAssignmentDirect(b *testing.B) {
	type X struct {
		p *int
	}
	for i := 0; i < b.N; i++ {
		var i int
		x := &X{
			p: &i, // GOOD: i does not escape
		}

		_ = x
	}
}
