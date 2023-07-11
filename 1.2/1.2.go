package main

import (
	"fmt"
	"os"
)

func main() {
	for i, arg := range os.Args {
		fmt.Println(i, arg)
	}
}

/*To print the index and value of each argument, we used
a 'for' loop that uses the 'range' keyword, which iterates
over 'os.Args', giving us both the index 'i' and the value
'arg' of each statement. Lastly, by using the 'Println'
function from the 'fmt' package, we get each index-value
in separate lines. We could also do that using '\n' (newline):

fmt.Printf("%d %s\n", i, arg), whereas %d is the verb for the
index in base 10, and %s is the verb for the uninterpreted bytes
of the string.

NOTE: THE ABOVE VERSION OF CODE DOES NOT PRINT INDEX '0'. TO OBTAIN IT,
WE CAN MODIFY THE CODE AS IN FOLLOWS:

----------------------------------------------------------------------

package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	for i, arg := range args {
		fmt.Println(i, arg)
	}
}

---------------------------------------------------------------------

Here, we used a 'slice' called 'args', to which we assigned
'os.Args[1:]'. Then again, by using the same logic as before,
we created a 'for' loop to iterate over the 'slice', giving us
both the index 'i' and value 'arg' of any given arguments through
the 'Println' function. Changing 'os.Args[1:]' to '0' would return
the first index. Should we wanted to prevent the index from
getting printed, we could use the blank identifier '_' in its place,
so that it would hide and later get discarded without any errors.
If we were to remove it completely, leaving 'args' only, we would
only get the indexes of the arguments and not the values themselves,
because the first variable initialized with the loop are the indexes*/
