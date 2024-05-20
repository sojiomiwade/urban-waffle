// To execute Go code, please declare a func main() in a package "main"

package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fish(*os.Stdout, os.Stdout)
}

func fish(f1 os.File, f2 *os.File) {
	fmt.Fprintf(&f1, "Hello %s\n", f1.Name())
	fmt.Fprintf(f2, "Hello %s\n", f2.Name())
	rice(f2)
	rice(&f1)

}

func rice(writer io.Writer) {
	writer.Write(make([]byte, 0))
}

/*
  writer.Write(make([]byte, 0))
translates to
Write(obj, ....)
typedef struct File
f=malloc(sizeof(File))

func (f *File) Write(b []byte) (n int, err error) {
  if err := f.checkValid("write"); err != nil {
    return 0, err
  }
  n, e := f.write(b)
  if n < 0 {
    n = 0
  }


# command-line-arguments
./solution.go:19:7: cannot use f1 (variable of type os.File) as io.Writer value in argument to rice: os.File does not implement io.Writer (method Write has pointer receiver)

*/
