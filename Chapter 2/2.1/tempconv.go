package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64 // Added the Kelvin custom type

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
	AbsoluteZeroK Kelvin  = 0 // Added the absolute zero temperature in Kelvin (0K)
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%gK", k) } // Added stringer method for Kelvin
