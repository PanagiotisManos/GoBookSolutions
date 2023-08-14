package main

import (
	"bytes"
	"fmt"
	"strings"
)

func comma(s string) string {
	var buf bytes.Buffer

	/* The code here checks whether the input number string 's' has a sign (positive or negative) by examining
	its prefix. If the string starts with either '+' or '-', it means there is a sign. The sign is extracted
	from the input string and stored in the sign variable. The sign is then removed from the input string 's'
	for further processing. */
	sign := ""
	if strings.HasPrefix(s, "+") || strings.HasPrefix(s, "-") {
		sign = s[:1]
		s = s[1:]
	}

	/* Here, the function splits the input number string 's' into its integral and fractional parts (if the number
	is a floating-point number). It uses the 'strings.Split' function to split the string at the decimal point '.'.
	The integral part is stored in the 'integralPart' variable, and the fractional part is stored in the
	'fractionalPart' variable. If there is no fractional part, 'fractionalPart' will remain an empty string. */
	parts := strings.Split(s, ".")
	integralPart := parts[0]
	fractionalPart := ""
	if len(parts) > 1 {
		fractionalPart = "." + parts[1]
	}

	/* Here we process the 'integralPart', which contains the digits before the decimal point. It adds commas
	to the 'buf' buffer to format the number. */
	n := len(integralPart)
	firstCommaPos := n % 3
	if firstCommaPos == 0 {
		firstCommaPos = 3
	}

	buf.WriteString(integralPart[:firstCommaPos])

	for i := firstCommaPos; i < n; i += 3 {
		buf.WriteByte(',')
		buf.WriteString(integralPart[i : i+3])
	}

	/* The function combines the sign (if present), the modified integral part (with commas), and the fractional
	part (if applicable) to create the final modified string named 'result'. This string is then returned as the output
	of the 'comma' function. */
	result := sign + buf.String() + fractionalPart
	return result
}

func main() {
	number1 := "1234567890"
	result := comma(number1)
	fmt.Println(result)

	number2 := "-9876543.210"
	result2 := comma(number2)
	fmt.Println(result2)
}
