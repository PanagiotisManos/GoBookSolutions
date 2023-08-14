package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; stroke-width: 0.7' "+ // We removed 'fill: white;' because now the color of each polygon will be determined by its height.
		"width='%d' height='%d' >", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {

			// The 'corner' function is now called with three 'return' values.
			ax, ay, aValid := corner(i+1, j)
			bx, by, bValid := corner(i, j)
			cx, cy, cValid := corner(i, j+1)
			dx, dy, dValid := corner(i+1, j+1)

			// For each polygon, it checks if all four corner points are valid
			// (aValid, bValid, cValid, and dValid) before drawing it.
			if aValid && bValid && cValid && dValid {

				// If all corners are valid, it calculates the color for the polygon
				// using the 'colorForPolygon' function and fills the polygon with the determined color.
				color := colorForPolygon(ax, ay, bx, by, cx, cy, dx, dy)
				fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='%s'/>\n",
					ax, ay, bx, by, cx, cy, dx, dy, color)
			}
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64, bool) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	/* We added an 'if' statement to check whether 'z' is a valid number.
	If the calculated 'z' value is either 'NaN' or '±Inf', it means there is an
	issue with the computation, and the function returns '(0, 0, false)'.
	The false boolean value indicates that the corner is not valid, and it
	should not be used to draw the polygon. */
	z := f(x, y)
	if math.IsInf(z, 0) || math.IsNaN(z) {
		return 0, 0, false
	}

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, true
}

/* The function 'colorForPolygon' calculates the color for each polygon based on the
heights of its four corners. For each corner, it calculates the 'z' value (height)
by converting the (x, y) coordinates back to their original scale using '(ax-width/2)/xyscale',
'(ay-height/2)/xyscale', etc., and then calling the 'f' function to calculate the height.
The average height of the polygon is then calculated as '(za + zb + zc + zd) / 4'. If the
average height is greater than zero, it means the polygon is a peak, and it is filled with
red color '#ff0000', whereas if the average height is less than or equal to zero, it means the
polygon is a valley, and it is filled with blue color '#0000ff'. */
func colorForPolygon(ax, ay, bx, by, cx, cy, dx, dy float64) string {
	za := f((ax-width/2)/xyscale, (ay-height/2)/xyscale)
	zb := f((bx-width/2)/xyscale, (by-height/2)/xyscale)
	zc := f((cx-width/2)/xyscale, (cy-height/2)/xyscale)
	zd := f((dx-width/2)/xyscale, (dy-height/2)/xyscale)

	averageHeight := (za + zb + zc + zd) / 4

	/* Color interpolation

	The color interpolation is done by adjusting the red 'r' and blue 'b' components of
	the RGB color representation based on the average height. The 'g' component is set to '0'
	to maintain a pure blue or pure red color. The color gradient ranges from blue (#0000ff)
	for lower heights (averageHeight close to -1) to white (#ffffff) for heights around 0,
	and finally to red (#ff0000) for higher heights (averageHeight close to 1). The fmt.Sprintf
	function is then used to convert the 'r', 'g', and 'b' integer values into a hexadecimal RGB
	color code of the form '#rrggbb', where 'rr', 'gg', and 'bb' are two-digit hexadecimal
	representations of the red, green, and blue components, respectively. */
	r := int(255 * (averageHeight + 1) / 2)
	g := 0
	b := int(255 * (1 - averageHeight) / 2)

	return fmt.Sprintf("#%02x%02x%02x", r, g, b)
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}
