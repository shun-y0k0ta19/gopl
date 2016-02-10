// Author: "Shun Yokota"
// Copyright © 2016 RICOH Co, Ltd. All rights reserved

// Package tempconv2 performs Celsius, Fahrenheit and Kelvin conversions
package tempconv2

//CToF is converting celcius to fahrenheit
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

//CToK is converting celcius to kelvin
func CToK(c Celsius) Kelvin { return Kelvin(c - AbsoluteZeroC) }

//FToC is converting fahrenheit to celcius
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

//FToK is converting fahrenheit to kelvin
func FToK(f Fahrenheit) Kelvin { return Kelvin(CToK(FToC(f))) }

//KToC is converting kelvin to celcius
func KToC(k Kelvin) Celsius { return Celsius(k + Kelvin(AbsoluteZeroC)) }

//KToF is converting kelvin to fahrenheit
func KToF(k Kelvin) Fahrenheit { return Fahrenheit(CToF(KToC(k))) }
