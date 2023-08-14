package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings" // added the 'strings' package
)

func main() {
	for _, url := range os.Args[1:] {

		/* We added another 'if' block inside the 'for' loop to check whether the url
		that is passed as an argument when executing the program has "http://" or
		"https://" as a prefix. If it does not, we make sure the program prepends "http://"
		automatically, so it can continue running. */
		if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}
