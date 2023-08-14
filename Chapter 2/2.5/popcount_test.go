// This code is used to benchmark the two popcount functions. To run it, simply
// type 'go test -bench=.' and see the results.

package popcount

import "testing"

func BenchmarkPopCountExpression(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = PopCount(0xDEADBEEF)
	}
}

func BenchmarkPopCountClearRightmost(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = PopCountClearRightmost(0xDEADBEEF)
	}
}

/*

1st run

BenchmarkPopCountExpression-4           1000000000               0.2782 ns/op
BenchmarkPopCountClearRightmost-4       62613486                 20.20 ns/op

2nd run

BenchmarkPopCountExpression-4           1000000000               0.2722 ns/op
BenchmarkPopCountClearRightmost-4       59606004                 20.75 ns/op

3rd run

BenchmarkPopCountExpression-4           1000000000               0.2774 ns/op
BenchmarkPopCountClearRightmost-4       56598166                 21.08 ns/op

The tests indicate that the PopCountExpression function is performing faster in the ns/op
scale (nanoseconds per operation) for the same reasons as before.

*/
