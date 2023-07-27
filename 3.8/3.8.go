package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"math/big"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	// We created an array of fractals, each represented by a name and a function.
	// The function for each fractal is a Mandelbrot set computation method. Four Mandelbrot
	// set computation methods are defined, each using different data types for precision.
	fractals := []struct {
		name string
		fn   func(complex128) color.Color
	}{
		{"complex64_fractal.png", mandelbrot64},
		{"complex128_fractal.png", mandelbrot},
		{"bigFloat_fractal.png", mandelbrotBigFloat},
		{"bigRat_fractal.png", mandelbrotRat},
	}

	// For each pixel in the image, the code maps the pixel coordinates to the corresponding
	// complex value 'z'. It then calls the corresponding Mandelbrot set computation function 'fn'
	// for that pixel's complex value 'z'. The return value from the computation function is used
	// to set the color of the pixel in the image.
	for _, f := range fractals {
		img := image.NewRGBA(image.Rect(0, 0, width, height))
		for py := 0; py < height; py++ {
			y := float64(py)/height*(ymax-ymin) + ymin
			for px := 0; px < width; px++ {
				x := float64(px)/width*(xmax-xmin) + xmin
				z := complex(x, y)
				img.Set(px, py, f.fn(z))
			}
		}

		// By using the 'os' package, we can save each function output in a seperate file.
		fl, err := os.Create(f.name)
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}
		png.Encode(fl, img)
		fmt.Println("Saved:", f.name)
		fl.Close()
	}
}

// Function using cmplx64.
func mandelbrot64(z complex128) color.Color {
	const iterations = 200
	// const contrast = 15
	var v complex64
	for n := uint8(0); n < iterations; n++ {
		v = v*v + complex64(z)
		if cmplx.Abs(complex128(v)) > 2 {
			switch {
			case n > 50:
				return color.RGBA{100, 0, 0, 255}
			default:
				// Logarithmic blue gradient to show small differences on the
				// periphery of the fractal.
				logScale := math.Log(float64(n)) / math.Log(float64(iterations))
				return color.RGBA{0, 0, 255 - uint8(logScale*255), 255}
			}
		}
	}
	return color.Black
}

// Function using cmplx128.
func mandelbrot(z complex128) color.Color {
	const iterations = 200
	// const contrast = 15
	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			switch {
			case n > 50:
				return color.RGBA{100, 0, 0, 255}
			default:

				logScale := math.Log(float64(n)) / math.Log(float64(iterations))
				return color.RGBA{0, 0, 255 - uint8(logScale*255), 255}
			}
		}
	}
	return color.Black
}

// Function using BigFloat.
func mandelbrotBigFloat(z complex128) color.Color {
	const iterations = 200
	// const contrast = 15
	zR := (&big.Float{}).SetFloat64(real(z))
	zI := (&big.Float{}).SetFloat64(imag(z))
	var vR, vI = &big.Float{}, &big.Float{}
	for i := uint8(0); i < iterations; i++ {
		// (r+i)^2 = r^2 + 2ri + i^2
		vR2, vI2 := &big.Float{}, &big.Float{}
		vR2.Mul(vR, vR).Sub(vR2, (&big.Float{}).Mul(vI, vI)).Add(vR2, zR)
		vI2.Mul(vR, vI).Mul(vI2, big.NewFloat(2)).Add(vI2, zI)
		vR, vI = vR2, vI2
		squareSum := &big.Float{}
		squareSum.Mul(vR, vR).Add(squareSum, (&big.Float{}).Mul(vI, vI))
		if squareSum.Cmp(big.NewFloat(4)) == 1 {
			switch {
			case i > 50:
				return color.RGBA{100, 0, 0, 255}
			default:

				logScale := math.Log(float64(i)) / math.Log(float64(iterations))
				return color.RGBA{0, 0, 255 - uint8(logScale*255), 255}
			}
		}
	}
	return color.Black
}

// Function using BigRat.
func mandelbrotRat(z complex128) color.Color {
	// High-resolution images take an extremely long time to render.
	// Multiplying arbitrary precision numbers has algorithmic complexity
	// of at least O(n*log(n)*log(log(n))). Here we used 10 iterations to generate an image
	// that is not complete, but gives the idea of being extremely precise.
	// (https://en.wikipedia.org/wiki/Arbitrary-precision_arithmetic#Implementation_issues).
	const iterations = 200
	// const contrast = 15
	zR := (&big.Rat{}).SetFloat64(real(z))
	zI := (&big.Rat{}).SetFloat64(imag(z))
	var vR, vI = &big.Rat{}, &big.Rat{}
	for i := uint8(0); i < iterations; i++ {
		// (r+i)^2 = r^2 + 2ri + i^2
		vR2, vI2 := &big.Rat{}, &big.Rat{}
		vR2.Mul(vR, vR).Sub(vR2, (&big.Rat{}).Mul(vI, vI)).Add(vR2, zR)
		vI2.Mul(vR, vI).Mul(vI2, big.NewRat(2, 1)).Add(vI2, zI)
		vR, vI = vR2, vI2
		squareSum := &big.Rat{}
		squareSum.Mul(vR, vR).Add(squareSum, (&big.Rat{}).Mul(vI, vI))
		if squareSum.Cmp(big.NewRat(4, 1)) == 1 {
			switch {
			case i > 50:
				return color.RGBA{100, 0, 0, 255}
			default:

				logScale := math.Log(float64(i)) / math.Log(float64(iterations))
				return color.RGBA{0, 0, 255 - uint8(logScale*255), 255}
			}
		}
	}
	return color.Black
}
