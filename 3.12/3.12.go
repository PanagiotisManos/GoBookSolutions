package main

import (
	"sort"
	"strings"
)

func isAnagram(str1, str2 string) bool {

	/* The first step inside the 'isAnagram' function is to convert both input strings 'str1' and 'str2' to lowercase
	using 'strings.ToLower'. This step is taken to make the comparison case-insensitive. For example, "Listen" and
	"silent" should be considered anagrams, and this conversion ensures that both strings are in the same case for
	comparison. */
	str1 = strings.ToLower(str1)
	str2 = strings.ToLower(str2)

	/*To compare the two strings' characters efficiently, the code converts each string to a slice of runes (Unicode
	code points) using '[]rune(str)'. This step allows the comparison to work with Unicode characters as well. */
	runes1 := []rune(str1)
	runes2 := []rune(str2)

	/* Next, the code uses the 'sort.Slice' function to sort the slices of runes 'runes1' and 'runes2' in lexicographical
	order. Sorting the slices ensures that the characters in each string are arranged in ascending order. Sorting is
	necessary for anagrams since rearranging the letters results in the same characters but in different orders. */
	sort.Slice(runes1, func(i, j int) bool { return runes1[i] < runes1[j] })
	sort.Slice(runes2, func(i, j int) bool { return runes2[i] < runes2[j] })

	/* Finally, the sorted rune slices 'runes1' and 'runes2' are converted back to strings using 'string(runes1)' and
	'string(runes2)', respectively. The function then compares the two sorted strings. If the sorted strings are equal,
	it means that the two input strings are anagrams of each other, and the function returns 'true'. Otherwise, it
	returns 'false'. */
	return string(runes1) == string(runes2)
}

func main() {
	str1 := "Listen"
	str2 := "Silent"

	if isAnagram(str1, str2) {
		println("The strings are anagrams.")
	} else {
		println("The strings are not anagrams.")
	}
}
