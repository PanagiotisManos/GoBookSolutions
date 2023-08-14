/* In this non-recursive version, we use a 'bytes.Buffer' function to efficiently build the modified string with
commas inserted. The buffer is more efficient than string concatenation because it minimizes the memory allocation
and copying overhead.
The function first determines the position of the first comma based on the length of the input string 's'. It then
writes the leading characters (before the first comma) to the buffer. After that, it iterates through the remaining
characters in the input string, inserting a comma after every three characters and writing them to the buffer.
Finally, the function returns the content of the buffer as a string using buf.String(). */

package main

import (
	"bytes"
	"fmt"
)

func comma(s string) string {
	var buf bytes.Buffer
	n := len(s)

	firstCommaPos := n % 3
	if firstCommaPos == 0 {
		firstCommaPos = 3
	}

	buf.WriteString(s[:firstCommaPos])

	for i := firstCommaPos; i < n; i += 3 {
		buf.WriteByte(',')
		buf.WriteString(s[i : i+3])
	}

	return buf.String()

}

func main() {
	number := "1234567890"
	result := comma(number)
	fmt.Println(result)
}
