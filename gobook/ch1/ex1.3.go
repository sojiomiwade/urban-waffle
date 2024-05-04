// To execute Go code, please declare a func main() in a package "main"

package main

import (
  "fmt"
  "strings"
  "time"
)

func main() {
  // bigNumber:=2
  bigNumber:=10000
  initialTime:=time.Now()
  operand:="Hello World"
  s1,s2:="",""
  var buf []string
  for i := 0; i < bigNumber; i++ {
    s1+=operand
  }
	time1:=time.Now().Sub(initialTime)

	initialTime=time.Now()
	for i := 0; i < bigNumber; i++ {
    buf = append(buf, operand)
  }
	// s2:=""
	s2=strings.Join(buf,"")
	time2:=time.Now().Sub(initialTime)

	fmt.Println(len(s1),len(s2))
  fmt.Printf("t1: %v, t2: %v\n",time1,time2)
}
