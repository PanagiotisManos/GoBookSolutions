// example web site is https://reddit.com

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	fileName := "output.txt" // output file name

	/* we use 'os.OpenFile' in append mode using the append flag 'os.O_APPEND'.
	That is to make sure that the output of the program execution is stored in the
	same file whenever we run it. The file is opened with write-only permission
	'os.O_WRONLY' and its mode is set to '0644' to ensure read and write permission
	for the owner and read-only for others. If the file to open and append does not
	exist, we create it with 'os.O_CREATE'. */
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("Failed to create file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	// redirect standard output to the file.
	oldStdout := os.Stdout
	os.Stdout = file

	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}

	/* restore standard output. The reason we do this is because if we didn't,
	the rest of the code (the two 'fmt.Printf' functions) would be diplayed in
	the file we created and not in the terminal. By restoring the standard output
	we make sure than any subsequent output is printed in the terminal. */
	os.Stdout = oldStdout

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
	fmt.Printf("Output saved to %s", fileName)
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}

/*
First run:  1.39s (secs) 1137157 (bytes)
Second run: 1.81s (secs) 1137596 (bytes)

We can see a difference in times and number of bytes called every time we run the program.
That can happen due to several reasons explained below:

1. Caching. When we make a request to a website, the response can be cached at different
levels (browser, OS, intermediate proxy servers). If a response is cached, subsequent
requests for the same URL may receive the cached response instead of fetching it again
from the server. Caching can significantly affect the number of bytes received and the
elapsed time since the response is not fetched anew.

2. Network conditions. Congestion, bandwidth limitations, or fluctuations in network
performance can impact the time taken to establish a connection, send the request,
and receive the response.

3. Server load. The more requests a server gets worldwide the more time it will need
to effectively access all of them.

4. Dynamic content. If the content on a website in dynamically generated, the response
body may vary slightly between different requests, leading to differences in the number
of bytes received.
*/
