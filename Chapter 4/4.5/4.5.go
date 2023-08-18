package main

import "fmt"

func fruitSlicer(strings []string) []string {

	// Simple test to check if there are any adjacent duplicates.
	if len(strings) <= 1 {
		return strings
	}

	currentIndex := 1

	/* The function's brain: This loop iterates through the input 'strings' slice
	starting from index '1' (since the first element was already considered unique).
	For each element at index 'i', it compares the element with the previous element
	(located at 'currentIndex-1'). If the current element is different from the
	previous one, it means that a new unique element is found. In this case,
	the current element is placed at the position of 'currentIndex', and then
	'currentIndex' is incremented. This process effectively eliminates adjacent
	duplicates by overwriting them with unique elements. */
	for i := 1; i < len(strings); i++ {
		if strings[i] != strings[currentIndex-1] {
			strings[currentIndex] = strings[i]
			currentIndex++
		}
	}

	// After the loop completes, the function returns a sub-slice of the original
	// 'strings' slice. This sub-slice contains only the unique elements, and its length
	// is equal to 'currentIndex', which keeps track of the last position where a unique
	// element was placed.
	return strings[:currentIndex]
}

func main() {
	input := []string{"apple", "apple", "banana", "banana", "banana", "orange", "apple", "apple", "apple", "grape", "grape"}

	result := fruitSlicer(input)
	fmt.Println("Output:", result)
}

// NOTE: This code works by returning a new sub-slice without modifying the original slice.
// If you want to modify the original slice in place, you can copy the elements from the result
// sub-slice back to the original slice.
