package main

import "fmt"

/* The 'newRotate' function takes two parameters: 's' (a slice of integers) and
'k' (the number of positions to rotate to the left). The rotation is done in
a circular manner, so if 'k' is greater than the length of the slice, it wraps
around. The length of the slice is stored in the variable 'n'. The value of 'k' is
adjusted to be within the range of the slice's length using the modulo operation
(k %= n). If 'k' is negative, it is transformed to its positive equivalent that
represents the same rotation in the opposite direction. We created a temporary
slice called 'temp' to store the rotated elements. The main loop iterates through
each element of the original slice 's' (with length 'n'). Inside the loop, the
formula '(i+k)%n' calculates the new index for each element after rotation.
This ensures that the elements wrap around in a circular manner. The value
at the current index of the original slice 's' is placed in the corresponding
index of the 'temp' slice. After the loop is done, the 'copy' function is used to
copy the rotated elements from the 'temp' slice back to the original slice 's',
effectively performing the rotation. */
func newRotate(s []int, k int) {
	n := len(s)
	k %= n

	if k < 0 {
		k += n
	}

	temp := make([]int, n)

	for i := 0; i < n; i++ {
		temp[(i+k)%n] = s[i]
	}

	copy(s, temp)
}

func main() {
	s := []int{0, 1, 2, 3, 4, 5}

	// The variable k is set to -2, indicating that you want to rotate the
	// slice to the left by 2 positions in reverse direction.
	k := -2

	newRotate(s, k)
	fmt.Println(s)
}
