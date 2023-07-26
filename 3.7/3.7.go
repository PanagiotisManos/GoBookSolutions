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

	/* The nested loops iterate over each pixel in the image. The 'x' and 'y' values are calculated
	based on the pixel coordinates and the given boundaries of the complex plane. A complex
	number 'z' is created using these 'x' and 'y' values, representing the starting point for
	Newton's method. The 'newtonFractal' function is called to determine the color of the pixel
	based on the root that 'z' approaches. */
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, newtonFractal(z))
		}
	}
	png.Encode(os.Stdout, img)
}

/* The 'newtonFractal' function takes a complex number 'z' as input and returns the color
that corresponds to the root it approaches. The function uses Newton's method to
iteratively find the root of the equation 'z^4 - 1 = 0' that 'z' approaches. The functions
'f' and 'df' represent the equation 'z^4 - 1' and its derivative '4z^3', respectively. The method
iterates a maximum of 'maxRootIterations' times or until 'z' is within the specified tolerance
of a root. Once a root is found, the color of the pixel is determined based on which root 'z'
approaches: red for root 1, green for root -1, blue for root i, and yellow for root -i. If the
method does not find a root within the maximum iterations, the pixel is colored black. */
func newtonFractal(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	f := func(z complex128) complex128 {
		return cmplx.Pow(z, 4) - 1
	}
	df := func(z complex128) complex128 {
		return 4 * cmplx.Pow(z, 3)
	}

	const tolerance = 1e-6
	const maxRootIterations = 50

	var root complex128
	for n := 0; n < maxRootIterations; n++ {
		if cmplx.Abs(f(z)) < tolerance {
			root = z
			break
		}
		z = z - f(z)/df(z)
	}

	switch {
	case cmplx.Abs(root-1) < tolerance:
		return color.RGBA{255, 0, 0, 255} // Red for root 1
	case cmplx.Abs(root+1) < tolerance:
		return color.RGBA{0, 255, 0, 255} // Green for root -1
	case cmplx.Abs(root+1i) < tolerance:
		return color.RGBA{0, 0, 255, 255} // Blue for root i
	case cmplx.Abs(root-1i) < tolerance:
		return color.RGBA{255, 255, 0, 255} // Yellow for root -i
	default:
		return color.Black
	}
}
