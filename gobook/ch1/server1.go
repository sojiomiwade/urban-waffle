// Server1 is a minimal "echo" server
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	hostIp := "localhost"
	port := os.Args[1]
	host := fmt.Sprintf("%s:%s", hostIp, port)
	fmt.Println("host: ", host)
	http.HandleFunc("/", handler) // each request calls handler
	log.Fatal(http.ListenAndServe(host, nil))
}

// handler echoes the Path component of the request URL r.
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
