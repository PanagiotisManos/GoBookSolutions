// generates and serves various fractal images over HTTP. It uses the 'image' and 'image/color'
// packages to create images, and the 'image/png' package to encode images in PNG format. The program
// can generate four types of fractals: Mandelbrot, Julia, Tricorn, and Newton.

package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"math/cmplx"
	"net/http"
	"strconv"
)

const (
	xmin, ymin, xmax, ymax = -2, -2, +2, +2
	width, height          = 1024, 1024
)

// Basic Handler code.
func main() {
	http.HandleFunc("/fractal", fractalHandler)
	http.ListenAndServe(":8000", nil)
}

// The 'fractalHandler' function handles HTTP requests to the '/fractal' route. It extracts the query parameters 'x',
// 'y', 'zoom' and 'type' from the request URL to determine the location and type of fractal to be generated. Depending
// on the 'type' parameter, it calls different fractal generation functions (mandelbrot, julia, tricorn, or newton) to
// generate the corresponding fractal image.
func fractalHandler(w http.ResponseWriter, r *http.Request) {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	x, y, zoom := getParameters(r)
	fractalType := r.URL.Query().Get("type")

	var img *image.RGBA
	switch fractalType {
	case "mandelbrot":
		img = generateFractal(x, y, zoom, width, height, mandelbrot)
	case "julia":
		cx, cy := -0.7, 0.27015 // You can change these values to explore different Julia fractals.
		img = generateFractal(x, y, zoom, width, height, func(z complex128) color.Color {
			return julia(z, complex(cx, cy))
		})
	case "tricorn":
		img = generateFractal(x, y, zoom, width, height, tricorn)
	case "newton":
		img = generateFractal(x, y, zoom, width, height, newton)
	default:
		img = generateFractal(x, y, zoom, width, height, mandelbrot)
	}

	w.Header().Set("Content-Type", "image/png")
	png.Encode(w, img)
}

// This function generates the specified fractal image by iterating over each pixel in the image and mapping it to
// the corresponding complex value in the fractal's region. It then calls the provided 'fractalFunc' (e.g., mandelbrot,
// julia, etc.) to determine the color of each pixel based on the complex value. The resulting image is returned as an
// RGBA image. When `zoom` is greater than 1, the fractal is zoomed out (made smaller) and when `zoom` is less than 1,
// the fractal is zoomed in (made larger). This is because the calculation for `xCoord` and `yCoord` will be scaled
// down and therefore cover a smaller area for a higher `zoom` and a larger area for a lower `zoom`.
func generateFractal(x, y, zoom float64, width, height int, fractalFunc func(complex128) color.Color) *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		yCoord := (float64(py)/float64(height)*(ymax-ymin) + ymin + y) * zoom
		for px := 0; px < width; px++ {
			xCoord := (float64(px)/float64(width)*(xmax-xmin) + xmin + x) * zoom
			z := complex(xCoord, yCoord)
			img.Set(px, py, fractalFunc(z))
		}
	}
	return img
}

// 'getParameters' parses the 'x', 'y', and 'zoom' parameters from the HTTP request and returns their 'float64' values.
// If any of the parameters are missing or are not valid floats, default values are used.
func getParameters(r *http.Request) (float64, float64, float64) {
	xParam := r.URL.Query().Get("x")
	yParam := r.URL.Query().Get("y")
	zoomParam := r.URL.Query().Get("zoom")

	x, err := strconv.ParseFloat(xParam, 64)
	if err != nil {
		x = 0.0
	}

	y, err := strconv.ParseFloat(yParam, 64)
	if err != nil {
		y = 0.0
	}

	zoom, err := strconv.ParseFloat(zoomParam, 64)
	if err != nil {
		zoom = 1.0
	}

	return x, y, zoom
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15
	var v complex128
	var magSquared float64
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		magSquared = real(v)*real(v) + imag(v)*imag(v)
		if magSquared > 4 {
			// Smooth coloring based on the number of iterations and the magnitude of the complex value.
			logZn := math.Log(magSquared) / 2
			nu := math.Log(logZn/math.Log(2)) / math.Log(2)
			n = uint8(float64(n) + 1 - contrast + contrast*math.Log1p(nu)/math.Log(2))
			r := uint8(255 * (float64(n) / float64(iterations)))
			g := uint8(255 * (1 - float64(n)/float64(iterations)))
			b := uint8(255 * (float64(n) / float64(iterations)))
			return color.RGBA{r, g, b, 255}
		}
	}
	return color.Black
}

func julia(z complex128, c complex128) color.Color {
	const iterations = 200
	const contrast = 15
	for i := uint8(0); i < iterations; i++ {
		z = z*z + c
		if cmplx.Abs(z) > 2 {
			// Color based on the number of iterations (Julia Fractal coloring).
			r := uint8((1 + math.Cos(float64(i)*0.08)) * 128)
			g := uint8((1 + math.Sin(float64(i)*0.1)) * 128)
			b := uint8((1 - math.Sin(float64(i)*0.05)) * 128)
			return color.RGBA{r, g, b, 255}
		}
	}
	return color.Black
}

func tricorn(z complex128) color.Color {
	const iterations = 200
	const contrast = 15
	var v complex128
	var magSquared float64
	for n := uint8(0); n < iterations; n++ {
		v = cmplx.Conj(v)*cmplx.Conj(v) + z
		magSquared = real(v)*real(v) + imag(v)*imag(v)
		if magSquared > 4 {
			// Smooth coloring based on the number of iterations and the magnitude of the complex value.
			logZn := math.Log(magSquared) / 2
			nu := math.Log(logZn/math.Log(2)) / math.Log(2)
			n = uint8(float64(n) + 1 - contrast + contrast*math.Log1p(nu)/math.Log(2))
			r := uint8(0)
			g := uint8(0)
			b := uint8(255 * (float64(n) / float64(iterations)))
			return color.RGBA{r, g, b, 255}
		}
	}
	return color.Black
}

func newton(z complex128) color.Color {
	const iterations = 37
	const contrast = 7
	for i := uint8(0); i < iterations; i++ {
		z = z - (z*z*z*z-1)/(4*z*z*z)
		if cmplx.Abs(z*z*z*z-1) < 1e-6 {
			// Color based on the roots of the equation (Newton Fractal coloring).
			theta := math.Atan2(imag(z), real(z))
			red := uint8((1 + math.Cos(theta)) * 128)
			green := uint8((1 + math.Sin(theta)) * 128)
			blue := uint8((1 - math.Sin(theta)) * 128)
			return color.RGBA{red, green, blue, 255}
		}
	}
	return color.Black
}
