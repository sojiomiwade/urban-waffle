// print the contents found in the specified URLs
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: Get failed on %s. Error: %q", url, err)
			os.Exit(1)
		}
		body := resp.Body
		bs, err := ioutil.ReadAll(body)
		body.Close()
		if err != nil {
			fmt.Fprintln(os.Stderr, "ReadAll for url %s: %s", url, err)
			os.Exit(1)
		}
		fmt.Printf("URL: %s. Body: \n %s\n", bs)

	}
}
