package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	num, _ := strconv.Atoi(os.Args[1])
	switch {
	case num > 0:
		fmt.Println("+ve")
	default:
		fmt.Println("0")
	case num < 0:
		fmt.Println("-ve")
	}
}
