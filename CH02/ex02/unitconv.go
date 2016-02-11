// Author: "Shun Yokota"
// Copyright © 2016 RICOH Co, Ltd. All rights reserved

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"gopl.io/ch2/tempconv"
)

func feetToMeter(ft feet) meter {
	return meter(ft * 0.3048)
}

func poundToKg(pd pound) kgram {
	return kgram(pd * 453.59237 / 1000)
}

func unitconv(param float64) (c tempconv.Celsius, mt meter, kg kgram) {
	c = tempconv.FToC(tempconv.Fahrenheit(param))
	mt = feetToMeter(feet(param))
	kg = poundToKg(pound(param))
	return c, mt, kg
}

func main() {
	param := os.Args[1:]
	if len(param) < 1 {
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			line := input.Text()
			t, err := strconv.ParseFloat(line, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "unitconv: %v\n", err)
			}
			c, mt, kg := unitconv(t)
			fmt.Printf("%s %s %s\n", c, mt, kg)
		}
	} else {
		t, err := strconv.ParseFloat(param[0], 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "unitconv: %v\n", err)
			os.Exit(1)
		}
		c, mt, kg := unitconv(t)
		fmt.Printf("%s %s %s\n", c, mt, kg)
	}

}
