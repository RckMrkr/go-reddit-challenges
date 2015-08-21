package easy228

import (
	"testing"
)

func BenchmarkInOrder(b *testing.B) {
	for n := 0; n < b.N; n++ {
		checkLineInFile("test_input")
	}
}
