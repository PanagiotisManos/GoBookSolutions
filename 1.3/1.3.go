/*To test the performance of all the versions of the program we
used the 'time' package measuring the execution duration. In all
cases the result was '0.00s seconds' of elapsed time, which proves
that all versions are capable of delivering fast execution times,
demonstrating the efficiency with which Golang can handle different
versions of the same logic.*/

//echo1

package main

import (
	"fmt"
	"os"
	"time"
)


	func main() {
		start := time.Now()
		var s, sep string
		for i := 1; i < len(os.Args); i++ {
			s += sep + os.Args[i]
			sep = " "
		}
		fmt.Println(s)
		fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
	}


/*echo2

package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
	s += sep + arg
	sep = " "
	}
	fmt.Println(s)
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
	}


echo3

package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
	}
*/
