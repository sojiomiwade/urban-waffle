// minimal echo and counter server
package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

/*
[x] set up two handlers
[x] create the two handler functions
[x] use a mutex to avoid the counter-read, and counter-update race condition
*/

var count int
var mu sync.Mutex

func main() {
	http.HandleFunc("/", updateCount)
	http.HandleFunc("/counter", readCount)
	log.Fatal(http.ListenAndServe("localhost:8002", nil))
}

// updateCount echoes the url path, and increments request count
func updateCount(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "url.path: %v\n", r.URL.Path)
}

// readCount returns the number of request counts so far
func readCount(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "count: %d\n", count)
	mu.Unlock()
}
