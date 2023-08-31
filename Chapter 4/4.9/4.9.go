// To test the program, create an input text file named input.txt in the same
// directory as the program. This is the file whose word frequencies you want to analyze.
// Then just run the go program and the results will be printed in the console.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	wordCounts := make(map[string]int)

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		word := strings.ToLower(scanner.Text())
		wordCounts[word]++
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Word\tFrequency\n")
	fmt.Println(strings.Repeat("-", 20))
	for word, count := range wordCounts {
		fmt.Printf("%s\t%d\n", word, count)
	}
}
