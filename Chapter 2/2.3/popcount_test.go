// This code is used to benchmark the two popcount functions. To run it, simply
// type 'go test -bench=.' and see the results.

package popcount

import "testing"

func BenchmarkPopCountExpression(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = PopCount(0xDEADBEEF)
	}
}

func BenchmarkPopCountLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = PopCountLoop(0xDEADBEEF)
	}
}

/*

1st run

BenchmarkPopCountExpression-4           1000000000               0.2695 ns/op
BenchmarkPopCountLoop-4                 154399776                7.731 ns/op

2nd run

BenchmarkPopCountExpression-4           1000000000               0.2862 ns/op
BenchmarkPopCountLoop-4                 145019065                8.140 ns/op

3rd run

BenchmarkPopCountExpression-4           1000000000               0.2847 ns/op
BenchmarkPopCountLoop-4                 136176702                8.038 ns/op

The tests indicate that the PopCountExpression function is performing better in the ns/op
scale (nanoseconds per operation). This version accesses the pc population count without
the need for a loop, summing it in a single expression, which could be optimized by the Go
compiler. The PopCountLoop function performs eight iterations to calculate the population
count and sums them up using a loop, involving more operations and memory accesses.

*/
