// OUTPUT spacings are not neat in order to keep code simpler.

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

// The 'getCategory' function is used to determine the Unicode category of the current rune.
func getCategory(r rune) string {
	if unicode.IsControl(r) {
		return "Control"
	}
	if unicode.IsDigit(r) {
		return "Digit"
	}
	if unicode.IsLetter(r) {
		return "Letter"
	}
	if unicode.IsMark(r) {
		return "Mark"
	}
	if unicode.IsPunct(r) {
		return "Punctuation"
	}
	if unicode.IsSymbol(r) {
		return "Symbol"
	}
	if unicode.IsSpace(r) {
		return "Space"
	}
	return "Other"
}

func main() {
	categoryCounts := make(map[string]int)

	// Using a slice instead of an array is a safer approach when dealing with dynamic sizing,
	// as it avoids issues related to constant array lengths. The 'utf8.UTFMax + 1' expression,
	// while it represents a constant value, might not be recognized as a constant by the compiler
	// in certain scenarios, leading to an error: "[utf8.UTFMax + 1]int (type) is not an expression".
	utflen := make([]int, utf8.UTFMax+1)
	invalid := 0

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}

		// The 'getCategory' function is called to determine the Unicode category of the current rune.
		// The counts are updated in the 'categoryCounts' map based on the category of the rune. If the
		// UTF-8 length is larger than or equal to the length of the 'utflen' slice, the slice's size is
		// extended accordingly. The UTF-8 length counts are updated in the 'utflen' slice.
		category := getCategory(r)
		categoryCounts[category]++
		if n >= len(utflen) {
			utflen = append(utflen, make([]int, n+1-len(utflen))...)
		}
		utflen[n]++
	}

	fmt.Printf("category\tcount\n")
	for c, n := range categoryCounts {
		fmt.Printf("%s\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if n > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
