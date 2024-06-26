// print the contents found in the specified URLs
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	httpPrefix := "http://"
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, httpPrefix) {
			url = httpPrefix + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: Get failed on %s. Error: %q", url, err)
			os.Exit(1)
		}
		body := resp.Body
		// _, err = io.Copy(os.Stdout, body)
		_, err = io.Copy(io.Discard, body)
		body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "copy failed for url %s: %s", url, err)
			os.Exit(1)
		}
		// fmt.Printf("URL: %s, number of bytes read: %d, http status code: %d\n", url, written, resp.StatusCode)

	}
}
