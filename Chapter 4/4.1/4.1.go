package main

import (
	"crypto/sha256"
	"fmt"
)

// The variable 'pc' is declared as an array of 256 bytes. This array will be used to calculate
// the population count (number of set bits) for each 8-bit value. It will be used later in the
// 'PopCount' function.
var pc [256]byte

// The 'init' function is defined to initialize the 'pc' array. The loop iterates over each element in
// the array (values from 0 to 255) and calculates the population count using bit manipulation. It
// uses the formula 'pc[i] = pc[i/2] + byte(i&1)' to calculate the population count for each value.
func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// The 'PopCount' function takes a 64-bit unsigned integer 'x' as input and returns the number of set bits
// '(1s)' in it. It does this by using the pc array to calculate the population count for each 8-bit chunk
// of 'x'. The function extracts each 8-bit chunk using bitwise right shift and masking with 'byte(x>>(i*8)) & 0xff'.
// It then looks up the precomputed population count from the pc array for each 8-bit chunk and adds them
// together to get the total number of set bits in 'x'.
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
	pc[byte(x>>(1*8))] +
	pc[byte(x>>(2*8))] +
	pc[byte(x>>(3*8))] +
	pc[byte(x>>(4*8))] +
	pc[byte(x>>(5*8))] +
	pc[byte(x>>(6*8))] +
	pc[byte(x>>(7*8))])
}

// The 'CountDifferentBits' function takes two SHA256 hashes as input and calculates the number of different
// bits between them. It iterates over each byte of the hashes and converts them to uint64 values to use in
// the 'PopCount' function. It XORs the two uint64 values to get the differences between the hashes and passes
// the result to 'PopCount' to count the number of differing bits. The function accumulates the count and returns
// the total.
func CountDifferentBits(hash1, hash2 [32]byte) int {
	count := 0
	for i := 0; i < len(hash1); i++ {
		x := uint64(hash1[i])
		y := uint64(hash2[i])

		diff := x ^ y
		count += PopCount(diff)
	}
	return count
}

// In the main function, we calculate SHA256 hashes for the strings "x" and "X" using the 'sha256.Sum256' function
// and stores them in 'c1' and 'c2', respectively. It then prints the hexadecimal representations of the two hashes
// using 'fmt.Printf'. Next, it calls the 'CountDifferentBits' function with the two hashes 'c1' and 'c2' as arguments
// to calculate the number of differing bits. Finally, it prints the number of different bits between the two
// hashes using 'fmt.Printf'.
func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))

	fmt.Printf("Hash 1: %x\n", c1)
	fmt.Printf("Hash 2: %x\n", c2)

	differences := CountDifferentBits(c1, c2)

	fmt.Printf("Number of different bits: %d\n", differences)
}
