// Package tempconv performs Celsius and Fahrenheit conversions
package main

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%gK", k) }

func CToK(c Celsius) Kelvin {
	return Kelvin(c - AbsoluteZeroC)
}

func KToC(k Kelvin) Celsius {
	return Celsius(k) + AbsoluteZeroC
}

func KToF(k Kelvin) Fahrenheit {
	return CToF(KToC(k))
}

func FToK(f Fahrenheit) Kelvin {
	return CToK(FToC(f))
}

// CToF converts a Celsius temparature to a Fahrenheit
func CToF(c Celsius) Fahrenheit {
	// 0 -> 32 9/5
	return Fahrenheit(9/5*c + 32)
}

// func FToC converts a Fahrenheit temparature to a Celsius
func FToC(f Fahrenheit) Celsius {
	//212
	return Celsius((f - 32) / 9 * 5)
}

func main() {
	fmt.Println(FToC(32))
	fmt.Println(FToK(32))
}
