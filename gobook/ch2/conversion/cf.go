// Cf converts its numeric argument to Celsius and Fahrenheit
/**
usage cf 32 0
32F = 0C, 32C = ..F
0F = ..C, 0C = 32F
-40F = -40c, -40C = -40F
*/
package main

import (
	"fmt"
	"os"
	"strconv"
	"/Users/soji/urban-waffle/gobook/ch2/tempconv"
)

func main() {
	for temp, _ := range os.Args[1:] {
		var tempF Fahrenheit = temp
		var tempC Celsius = temp
		fmt.Printf("%s = %s, ", tempF, tempconv.FToC(tempF))
		fmt.Printf("%s = %s\n", tempC, tempconv.CToF(tempC))
	}
}
