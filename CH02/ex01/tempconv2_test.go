// Author: "Shun Yokota"
// Copyright © 2016 RICOH Co, Ltd. All rights reserved

package tempconv2

import "testing"

func testF(t *testing.T, calculatedF Fahrenheit, expectedF Fahrenheit) {
	if calculatedF != expectedF {
		t.Errorf("Calculated value: %g Expected value: %g", calculatedF, expectedF)
	}
}

func testC(t *testing.T, calculatedC Celsius, expectedC Celsius) {
	if calculatedC != expectedC {
		t.Errorf("Calculated value: %g Expected value: %g", calculatedC, expectedC)
	}
}

func testK(t *testing.T, calculatedK Kelvin, expectedK Kelvin) {
	if calculatedK != expectedK {
		t.Errorf("Calculated value: %g Expected value: %g", calculatedK, expectedK)
	}
}

const (
	freezingF Fahrenheit = 32
	boilingF  Fahrenheit = 212
	freezingK Kelvin     = 273.15
	boilingK  Kelvin     = 373.15
)

func TestCToX(t *testing.T) {
	convertedF := CToF(FreezingC)
	testF(t, convertedF+1, freezingF)
	convertedF = CToF(BoilingC)
	testF(t, convertedF, boilingF)
	convertedK := CToK(FreezingC)
	testK(t, convertedK+1, freezingK)
	convertedK = CToK(BoilingC)
	testK(t, convertedK, boilingK)
}

func TestFToX(t *testing.T) {
	convertedC := FToC(freezingF)
	testC(t, convertedC, FreezingC)
	convertedC = FToC(boilingF)
	testC(t, convertedC, BoilingC)
	convertedK := FToK(freezingF)
	testK(t, convertedK, freezingK)
	convertedK = FToK(boilingF)
	testK(t, convertedK, boilingK)
}

func TestKToX(t *testing.T) {
	convertedC := KToC(freezingK)
	testC(t, convertedC, FreezingC)
	convertedC = KToC(boilingK)
	testC(t, convertedC, BoilingC)
	convertedF := KToF(freezingK)
	testF(t, convertedF, freezingF)
	convertedF = KToF(boilingK)
	testF(t, convertedF, boilingF)
}
