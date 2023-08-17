package main

import "fmt"

// The 'reverse' function takes an array pointer 's *[5]int' as an argument.
func reverse(s *[5]int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// In the 'main' function, we created an array 's' and then passed a pointer to
// this array to the 'reverse' function using '&s'.
func main() {
	s := [5]int{1, 2, 3, 4, 5}
	reverse(&s)
	fmt.Println(s)
}
