/*To get the program to print out the names of the files that
have duplicate lines, we changed the map definition to 'map[string][]string',
which maps each line to a slice of filenames. We also added an argument
called 'filename' to the 'countLines' function, indicative of the names of
files being scanned. We then modified the map in the function to append the
'filename' to the slice of filenames associated with each line. Lastly, we
modified the 'for' loop so that it also prints out the names of the files,
using 'range' to iterate over the keys and values of the 'counts' map.*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, "stdin")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, arg)
			f.Close()
		}
	}
	for line, files := range counts {
		if len(files) > 1 {
			fmt.Printf("%d\t%s\t%v\n", len(files), line, files)
		} 
	}
}

func countLines(f *os.File, counts map[string][]string, filename string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		counts[line] = append(counts[line], filename)
	}
}
