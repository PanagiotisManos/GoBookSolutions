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
		xmin, ymin, xmax, ymax = -1, -1, +3, +3 // Changed the min and max values to fit the Mandelbrot set in the canvas.
		width, height          = 1024, 1024
		ssWidth, ssHeight      = 2, 2 // Supersampling width and height. Since we set the value for both to 2, supersampling
									  // is set to 2x2 (4 subpixels per pixel).
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		for px := 0; px < width; px++ {

			// The 'avgColor' variable is used to store the accumulated color values of the subpixels. For each subpixel, it
			// calls the 'mandelbrot' function to get the color and adds it to the 'avgColor' using the 'addColors' helper function.
			var avgColor color.RGBA
			for subX := 0; subX < ssWidth; subX++ {
				for subY := 0; subY < ssHeight; subY++ {
					x := float64(px)/width*(xmax-xmin) + xmin
					y := float64(py)/height*(ymax-ymin) + ymin
					x += float64(subX)/float64(ssWidth*width)*(xmax-xmin) + xmin
					y += float64(subY)/float64(ssHeight*height)*(ymax-ymin) + ymin
					z := complex(x, y)
					avgColor = addColors(avgColor, mandelbrot(z))
				}
			}

			// Here we average the color values.
			avgColor.R = uint8(uint32(avgColor.R) / (ssWidth * ssHeight))
			avgColor.G = uint8(uint32(avgColor.G) / (ssWidth * ssHeight))
			avgColor.B = uint8(uint32(avgColor.B) / (ssWidth * ssHeight))
			avgColor.A = 255
			img.Set(px, py, avgColor)
		}
	}
	png.Encode(os.Stdout, img)

	// Code we implemented for creating a png file output using the 'os' package.
	file, err := os.Create("mandelbrot.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	png.Encode(file, img)
}

func mandelbrot(z complex128) color.RGBA {
	const (
		iterations = 200
		contrast   = 15
	)

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.RGBA{255 - contrast*n, 0, 0, 255}
		}
	}
	return color.RGBA{0, 0, 0, 255}
}

// The 'addColors' helper function takes two 'color.RGBA' values and adds their RGB components together. It is used to accumulate the
// color values of the subpixels.
func addColors(c1, c2 color.RGBA) color.RGBA {
	return color.RGBA{
		R: c1.R + c2.R,
		G: c1.G + c2.G,
		B: c1.B + c2.B,
		A: 0,
	}
}
