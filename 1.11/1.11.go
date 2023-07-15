/* Using same code as in prior solution for obvious reasons. If
a website does not respond, an error message is returned by the 'http.Get'
function and sent to the 'ch' channel. The program handles failed HTTP requests
accordingly. */

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
	fileName := "output.txt"

	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("Failed to create/open file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	oldStdout := os.Stdout
	os.Stdout = file

	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}

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

/* If we were to improve performance of our code, we can introduce cancellation
functionality, using the 'context' package. The following code is using cancellation
to send a signal to the 'fetch' function to exit early when it encounters a failed
request. */

/*
package main

import (
	"context" // new package introducing cancellation functionality
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
	fileName := "cancellation.txt"

	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("Failed to create/open file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	oldStdout := os.Stdout
	os.Stdout = file

	// creating a 'context.WithCancel' function which returns 'ctx' and 'cancel'
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for _, url := range os.Args[1:] {
		go fetch(ctx, url, ch)
	}

	// we created a seperate go routine that sleeps for 2 seconds and then calls the cancel function.
	// The 2 second mark is given so as the fetch function can have some time to complete its purpose, that
	// is the fetching operation, before being cancelled by the cancellation singal.
	go func() {
		time.Sleep(2 * time.Second)
		cancel()
	}()

	for range os.Args[1:] {
		fmt.Println(<-ch)
	}

	os.Stdout = oldStdout

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
	fmt.Printf("Output saved to %s", fileName)
}

func fetch(ctx context.Context, url string, ch chan<- string) { // passing the cancellation to 'fetch' function
	start := time.Now()

	// The 'http.NewRequestWithContext' function creates an HTTP request 'req'
	// with the provided cancellation context 'ctx'.
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		ch <- fmt.Sprintf("Error creating request for %s: %v", url, err)
		return
	}

	// The 'http.DefaultClient.Do(req)' method sends the created request 'req' and
	// returns 'resp' or an error. The 'Do' method performs the actual HTTP request.
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		ch <- fmt.Sprintf("Error while fetching %s: %v", url, err)
		return
	}
	defer resp.Body.Close()

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("Error while reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
*/
