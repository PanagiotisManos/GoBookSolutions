// Code that supports conversions for temperature, length and weight.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"gopl.io/ch2/tempconv"
)

func main() {
	if len(os.Args) > 1 {
		convertArgs(os.Args[1:])
	} else {
		convertStdin()
	}
}

func convertArgs(args []string) {
	for _, arg := range args {
		convertAndPrint(arg)
	}
}

func convertStdin() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		arg := scanner.Text()
		convertAndPrint(arg)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "failed to read input: %v\n", err)
		os.Exit(1)
	}
}

func convertAndPrint(arg string) {
	value, err := strconv.ParseFloat(arg, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "invalid input: %s\n", arg)
		return
	}

	// Temperature conversion
	c := tempconv.Celsius(value)
	f := tempconv.Fahrenheit(value)
	fmt.Printf("%.2f°C = %.2f°F\n", c, tempconv.CToF(c))
	fmt.Printf("%.2f°F = %.2f°C\n", f, tempconv.FToC(f))

	// Length conversion
	m := Meter(value)
	ft := Foot(value)
	fmt.Printf("%.2fm = %.2fft\n", m, MToFt(m))
	fmt.Printf("%.2fft = %.2fm\n", ft, FtToM(ft))

	// Weight conversion
	kg := Kilogram(value)
	lb := Pound(value)
	fmt.Printf("%.2fkg = %.2flb\n", kg, KGToLb(kg))
	fmt.Printf("%.2flb = %.2fkg\n", lb, LbToKG(lb))

	fmt.Println() // Terminal prints an empty line between conversions
}

// Length conversion functions
type Meter float64
type Foot float64

func MToFt(m Meter) Foot {
	return Foot(m * 3.28084)
}

func FtToM(ft Foot) Meter {
	return Meter(ft / 3.28084)
}

// Weight conversion functions
type Kilogram float64
type Pound float64

func KGToLb(kg Kilogram) Pound {
	return Pound(kg * 2.20462)
}

func LbToKG(lb Pound) Kilogram {
	return Kilogram(lb / 2.20462)
}
