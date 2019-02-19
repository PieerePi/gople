// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 43.
//!+

// Cf converts its numeric argument to Celsius and Fahrenheit.
package main

import (
	"fmt"
	"os"
	"strconv"

	"gople/ch2/e2.2/unitconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := unitconv.Fahrenheit(t)
		c := unitconv.Celsius(t)
		m := unitconv.Meter(t)
		i := unitconv.Inch(t)
		k := unitconv.Kilo(t)
		p := unitconv.Pound(t)
		fmt.Printf("%s = %s, %s = %s\n",
			f, unitconv.FToC(f), c, unitconv.CToF(c))
		fmt.Printf("%s = %s, %s = %s\n",
			m, unitconv.MeterToInch(m), i, unitconv.InchToMeter(i))
		fmt.Printf("%s = %s, %s = %s\n",
			k, unitconv.KiloToPound(k), p, unitconv.PoundToKilo(p))
	}
}

//!-
