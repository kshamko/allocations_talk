package example4

import (
	"testing"
)

func BenchmarkAppend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sliceAppend(1000)
	}
}

func BenchmarkIndex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sliceIndex(1000)
	}
}
