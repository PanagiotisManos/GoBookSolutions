package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

var palette = []color.Color{
	color.RGBA{0x00, 0x00, 0x00, 0xff}, // HEX black color,
	color.RGBA{0x00, 0xff, 0x00, 0xff}, // HEX green color,
	color.RGBA{0x00, 0x00, 0xff, 0xff}, // HEX blue color,
	color.RGBA{0xff, 0x00, 0x00, 0xff}, // HEX red color,
}

const (
	blackIndex = 0 // first color in palette
	greenIndex = 1 // next color in palette
	blueIndex  = 2
	redIndex   = 3
)

func main() {
	lissajous(os.Stdout)
}
func lissajous(out io.Writer) {
	const (
		cycles    = 5
		res       = 0.001
		size      = 100
		nframes   = 64
		delay     = 8
		colorStep = 1 // constant that represents step size for changing the color
	)			      // index in each frame.

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	colorIndex := blackIndex // initial color index
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)

			/* Each time the loop runs, 'colorIndex' is incremented by 'colorStep' With
			the modulo operator '%' with 'len(palette)', the color index wraps around
			when it exceeds the number of colors in the palette. 'colorStep' essentially
			controls the rate of change in the color index. The larger its value, the more
			rapid color changes will be and the opposite. */
			colorIndex = (colorIndex + colorStep) % len(palette)

			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				uint8(colorIndex)) // we casted the variable 'colorIndex' to a uint8, since
		    }					   // the 'SetColorIndex' function expects a uint8 value as
								   // its third argument.
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
