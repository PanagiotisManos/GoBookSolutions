// This code is used to benchmark the two popcount functions. To run it, simply
// type 'go test -bench=.' and see the results.

package popcount

import "testing"

func BenchmarkPopCountExpression(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = PopCount(0xDEADBEEF)
	}
}

func BenchmarkPopCountShift(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = PopCountShift(0xDEADBEEF)
	}
}

/*

1st run

BenchmarkPopCountExpression-4           1000000000               0.2709 ns/op
BenchmarkPopCountShift-4                39009420                 35.28 ns/op

2nd run

BenchmarkPopCountExpression-4           1000000000               0.2890 ns/op
BenchmarkPopCountShift-4                31600401                 37.04 ns/op

3rd run

BenchmarkPopCountExpression-4           1000000000               0.2693 ns/op
BenchmarkPopCountShift-4                29447852                 35.41 ns/op

The tests indicate that the PopCountExpression function is still performing better in the
ns/op scale (nanoseconds per operation) for the same reasons as before.

*/
