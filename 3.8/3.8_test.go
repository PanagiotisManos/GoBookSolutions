// With the benchmark test file we can examine details regarding the operation of each
// function. To run the test, type 'go test -bench . -benchmem' in your terminal. This
// test also generates a memory output file for further examination.

package main

import (
	"fmt"
	"image/color"
	"os"
	"runtime"
	"runtime/pprof"
	"testing"
)

func BenchmarkMandelbrot64(b *testing.B) {
	memProfileFile := startMemProfile()
	defer stopMemProfile(memProfileFile)

	benchmarkMandelbrot(b, mandelbrot64)
}

func BenchmarkMandelbrot(b *testing.B) {
	memProfileFile := startMemProfile()
	defer stopMemProfile(memProfileFile)

	benchmarkMandelbrot(b, mandelbrot)
}

func BenchmarkMandelbrotBigFloat(b *testing.B) {
	memProfileFile := startMemProfile()
	defer stopMemProfile(memProfileFile)

	benchmarkMandelbrot(b, mandelbrotBigFloat)
}

func BenchmarkMandelbrotRat(b *testing.B) {
	memProfileFile := startMemProfile()
	defer stopMemProfile(memProfileFile)

	benchmarkMandelbrot(b, mandelbrotRat)
}

func benchmarkMandelbrot(b *testing.B, f func(complex128) color.Color) {
	for i := 0; i < b.N; i++ {
		f(complex(float64(i), float64(i)))
	}
}

func startMemProfile() *pprof.Profile {
	runtime.GC()
	return pprof.Lookup("heap")
}

func stopMemProfile(profile *pprof.Profile) {
	if profile != nil {
		memProfileFile, err := os.Create("mem_profile.out")
		if err != nil {
			fmt.Println("Error creating memory profile:", err)
			return
		}
		defer memProfileFile.Close()

		if err := profile.WriteTo(memProfileFile, 0); err != nil {
			fmt.Println("Error writing memory profile:", err)
		}
	}
}
