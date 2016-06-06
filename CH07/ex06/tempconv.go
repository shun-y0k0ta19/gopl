// Copyright © 2016 "Shun Yokota" All rights reserved

// Package tempconv performs Celsius and Fahrenheit temperature computations.
package main

import (
	"flag"
	"fmt"
)

//Celsius is a type of celsius
type Celsius float64

//Fahrenheit is a type of fahrenheit
type Fahrenheit float64

//Kelvin is is a type of kelvin
type Kelvin float64

//CToF is converting celcius to fahrenheit
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9.0/5.0 + 32.0) }

//CToK is converting celcius to kelvin
func CToK(c Celsius) Kelvin { return Kelvin(c - AbsoluteZeroC) }

//FToC is converting fahrenheit to celcius
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32.0) * 5.0 / 9.0) }

//FToK is converting fahrenheit to kelvin
func FToK(f Fahrenheit) Kelvin { return Kelvin(CToK(FToC(f))) }

//KToC is converting kelvin to celcius
func KToC(k Kelvin) Celsius { return Celsius(k + Kelvin(AbsoluteZeroC)) }

//KToF is converting kelvin to fahrenheit
func KToF(k Kelvin) Fahrenheit { return Fahrenheit(CToF(KToC(k))) }

//AbsoluteZeroC is absolute zero in celsius
const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }

// *celsiusFlag satisfies the flag.Value interface.
type celsiusFlag struct{ Celsius }

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit) // no error check needed
	switch unit {
	case "C", "°C":
		f.Celsius = Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = FToC(Fahrenheit(value))
		return nil
	case "K", "":
		f.Celsius = KToC(Kelvin(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

// CelsiusFlag defines a Celsius flag with the specified name,
// default value, and usage, and returns the address of the flag variable.
// The flag argument must have a quantity and a unit, e.g., "100C".
func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}
