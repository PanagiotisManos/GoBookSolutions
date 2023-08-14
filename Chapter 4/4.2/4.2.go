package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)

// In the main function we start by defining a command-line flag named 'algorithm' with a default
// value of 'sha256'. This flag will be used to specify the hash algorithm to use. The 'flag.StringVar'
// function is used to bind the algorithm variable to the 'algorithm' flag.
func main() {
	var algorithm string
	flag.StringVar(&algorithm, "algorithm", "sha256", "Hash algorithm")

	// The 'flag.Parse()' function is called to parse the command-line flags provided by the user.
	flag.Parse()

	// Here we retrieve the input string from the command-line arguments using flag.Args().
	// It checks if any input string is provided, and if not, it prints an error message and exits the program.
	args := flag.Args()
	if len(args) == 0 {
		fmt.Println("Error: No input string provided.")
		return
	}
	input := args[0]

	// Next, we generate the hash output using the specified hash algorithm. It uses a switch
	// statement to handle different algorithms.
	switch algorithm {
	case "sha256":
		hash := sha256.Sum256([]byte(input))
		fmt.Printf("SHA-256 hash: %x\n", hash)
	case "sha384":
		hash := sha512.Sum384([]byte(input))
		fmt.Printf("SHA-384 hash: %x\n", hash)
	case "sha512":
		hash := sha512.Sum512([]byte(input))
		fmt.Printf("SHA-512 hash: %x\n", hash)
	default:
		fmt.Println("Error: Invalid hash algorithm specified.")
	}
}
