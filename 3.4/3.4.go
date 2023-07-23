// Working on the logic of the 'Lissajous Web Server' (Section 1.7), we created the following code.

package main

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
)

const (
	width, height = 600, 320 /* This is the default max canvas size. That means that we can't change the 2 values
							    any higher than 600x320. To do that, we would have to implement some HTTP/Javascript
							    code that would set the canvas size to whatever the native resolution of the user's
							    monitor would be. */
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/svg+xml") // You can check if the Content-Type header is correct
														// with a tool like curl in a terminal. Type the following:
														// curl -I http://localhost:8000

		// We used the 'r.FormValue' function to be able to extract the three query parameters
		// (width, height, color) from the HTTP request.
		widthStr	:= r.FormValue("width")
		heightStr	:= r.FormValue("height")
		color		:= r.FormValue("color")

		if color == "" {
			color = "white" // Here we set the default color of the surface, if no specific color is provided.
		}

		widthInt, err := strconv.Atoi(widthStr)
		if err != nil || widthInt <= 0 {
			widthInt = width
		}
		heightInt, err := strconv.Atoi(heightStr)
		if err != nil || heightInt <= 0 {
			heightInt = height
		}

		/* Here we use a format string '%s' to include the value of the color variable in place of '%s'.
		The value of color will be the fill color of the SVG elements, and it can be specified
		by the user as a query parameter in the HTTP request. */
		fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
			"style='stroke: grey; fill: %s; stroke-width: 0.7' "+
			"width='%d' height='%d'>", color, widthInt, heightInt)

		for i := 0; i < cells; i++ {
			for j := 0; j < cells; j++ {
				ax, ay := corner(i+1, j)
				bx, by := corner(i, j)
				cx, cy := corner(i, j+1)
				dx, dy := corner(i+1, j+1)
				fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
					ax, ay, bx, by, cx, cy, dx, dy)
			}
		}
		fmt.Fprintln(w, "</svg>")
	})

	http.ListenAndServe(":8000", nil)
}

func corner(i, j int) (float64, float64) {

	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x, y)

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}
