package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, mandelbrot(z))
		}
	}

	// Code we implemented for creating a png file output using the 'os' package.
	file, err := os.Create("mandelbrot.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	png.Encode(file, img)
}

/* Here, the mandelbrot function returns a 'color.RGBA' value, representing a
full-color pixel based on the number of iterations it takes for a point in the
complex plane to escape the Mandelbrot set. The RGB values are calculated based
on the iteration count to produce a gradient-like effect. */
func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			r := 255 - contrast*n
			g := 2 * contrast * n
			b := 2 * contrast * n
			return color.RGBA{r, g, b, 255}
		}
	}
	return color.Black
}
