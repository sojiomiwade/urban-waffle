package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	if len(os.Args)<3{
		fmt.Println("usage: fetchall output_filename <at least 1 url>")
		os.Exit(1)
	}
  start := time.Now()
  ch := make(chan string)
  urls := os.Args[2:]
  output_filename := os.Args[1]
  ofile, err := os.Create(output_filename)
  if err != nil {
    fmt.Fprintf(os.Stderr, "could not open %s for writing. error: %v\n", output_filename, err)
    os.Exit(1)
  }

  for idx, url := range urls {
    go fetch(url, idx, ch)
  }
  for range urls {
    ofile.WriteString(<-ch)
  }
  ofile.WriteString(fmt.Sprintf("%.2fs elapsed\n", time.Since(start).Seconds()))
  err = ofile.Close()
  if err != nil {
    fmt.Fprintf(os.Stderr, "could not close output file %v; error: %v\n", ofile, err)
  }
}

func fetch(url string, idx int, ch chan<- string) {
  start := time.Now()
  if idx == 3 {
    time.Sleep(10 * time.Second)
  }
	httpPrefix:="http://"
	if !strings.HasPrefix(url,httpPrefix) {
		url=httpPrefix+url
	}
  resp, err := http.Get(url)
  if err != nil {
    ch <- fmt.Sprintf("get failed for url %q, err: %v\n",url,err)
    return
  }
  nbytes, err := io.Copy(io.Discard, resp.Body)

  resp.Body.Close()
  if err != nil {
    ch <- fmt.Sprintf("while reading %s: %v\n", url, err)
    return
  }
  secs := time.Since(start).Seconds()
  ch <- fmt.Sprintf("%.2fs %7dK %s\n", secs, nbytes/1e3, url)
}
