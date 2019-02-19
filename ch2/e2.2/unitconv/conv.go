// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 41.

//!+

package unitconv

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// MeterToInch converts Meter to Inch.
func MeterToInch(m Meter) Inch { return Inch(m / 0.0254) }

// InchToMeter converts Inch to Meter.
func InchToMeter(i Inch) Meter { return Meter(i * 0.0254) }

// KiloToPound converts Kilo to Pound.
func KiloToPound(k Kilo) Pound { return Pound(k / 0.4535924) }

// PoundToKilo converts Pound to Kilo.
func PoundToKilo(p Pound) Kilo { return Kilo(p * 0.4535924) }

//!-
