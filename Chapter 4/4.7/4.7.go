/* The code does not allocate new memory for reversing the UTF-8-encoded string.
It follows an in-place reversal approach, meaning it modifies the existing memory of the
input '[]byte' slice to achieve the reversal.
The primary goal of the code is to reverse the characters of a UTF-8-encoded string without
creating a new '[]byte' slice. It first reverses the entire slice using the 'reverseBytes'
function, which is an in-place reversal of byte sequences. Then, it iterates through the
reversed slice and reverses the individual UTF-8 character sequences in place while maintaining
the UTF-8 encoding. */

package main

import (
	"fmt"
	"unicode/utf8"
)

// This is a helper function that reverses the bytes of a given '[]byte' slice 'b'.
func reverseBytes(b []byte) {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
}

// This is the main function that reverses the UTF-8-encoded string.
// It takes a '[]byte' slice 's' as input.
func reverseUTF8(s []byte) {

	reverseBytes(s)

	for i := 0; i < len(s); {
		_, size := utf8.DecodeRune(s[i:])
		if size > 1 {
			reverseBytes(s[i : i+size])
		}
		i += size
	}
}

func main() {
	input := "Hello World!"
	fmt.Println("Original:", input)

	byteSlice := []byte(input)

	reverseUTF8(byteSlice)

	reversedString := string(byteSlice)
	fmt.Println("Reversed:", reversedString)
}
