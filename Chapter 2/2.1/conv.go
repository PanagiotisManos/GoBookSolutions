package tempconv

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// Added functons that convert Kelvin to Celsius and Fahrenheit and vice-versa
func KToC(k Kelvin) Celsius { return Celsius(k + Kelvin(AbsoluteZeroC)) }

func CToK(c Celsius) Kelvin { return Kelvin(c - AbsoluteZeroC) }

func KToF(k Kelvin) Fahrenheit { return Fahrenheit((k+Kelvin(AbsoluteZeroC))*9/5 + 32) }

func FToK(f Fahrenheit) Kelvin { return Kelvin((f + 459.67) * 5 / 9) }
