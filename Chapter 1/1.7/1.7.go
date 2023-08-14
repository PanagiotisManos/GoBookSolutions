package main

import (
	"fmt"
	"io" // replaced 'io/ioutil' with 'io' package
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := io.Copy(os.Stdout, resp.Body) // replaced 'ioutil.ReadAll' function with 'io.Copy'
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: copying %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%d", b) // replaced '%s' with 'd'
	}
}

/* In this code, we replaced the 'io/ioutil' package with 'io', since it
provides the necessary 'io.Copy' function. Using 'os.Stdout', which is the
standard output, as the destination, helps us avoid loading the entire
response body in the memory. If an error was to occur during the copying
process, the 'err' block would be executed, printing an error message to
the standard error stream 'os.Stderr', after which the program would exit
with a non-zero status code. If we run the program, at the end where the
last 'fmt.Printf' function is executed, we would also get the number of
bytes copied. This is because of the function signature of 'io.Copy':

func Copy(dst Writer, src Reader) (written int64, err error),

where 'written' is the number of bytes successfully copied from 'src' to
'dst'. You can notice in the function signature that 'written' has to be
of an 'int64' type. This is why we changed the last fmt.Printf function to
return 'b' as an integer using the '%d' verb, instead of the former '%s'.

If we wanted to discard the byte count to save memory, we would have to
replace 'b' with an underscore '_'. */
