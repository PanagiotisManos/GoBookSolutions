package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

/* The 'squashSpaces' function iterates through each byte of the input 'data' slice.
The 'utf8.DecodeRune(data[j:])' call decodes the first rune (Unicode character) at
the current position 'j', returning the decoded rune and its byte size 'size'. If the
current byte is a Unicode space: The loop advances 'j' until non-space bytes are
encountered. A single ASCII space is written to the position 'i' in the 'data' slice,
and 'i' is incremented. If the current byte is not a space: The rune is copied from
position 'j' to position 'i' in the 'data' slice. Both 'i' and 'j' are incremented by
the rune's size. The function returns the modified 'data' slice, containing the squashed
spaces. */
func squashSpaces(data []byte) []byte {
	i := 0
	for j := 0; j < len(data); {
		_, size := utf8.DecodeRune(data[j:])
		if unicode.IsSpace(rune(data[j])) {
			for j < len(data) && unicode.IsSpace(rune(data[j])) {
				j += size
			}
			data[i] = ' '
			i++
		} else {
			copy(data[i:i+size], data[j:j+size])
			i += size
			j += size
		}
	}
	return data[:i]
}

func main() {
	input := []byte("I		need	to  be				squashed!")
	output := squashSpaces(input)
	fmt.Printf("Squashed: %s\n", output)
}
