// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

//!+

// Package tempconv performs Celsius and Fahrenheit conversions.
package unitconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Meter float64
type Inch float64
type Kilo float64
type Pound float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (m Meter) String() string      { return fmt.Sprintf("%g(m)", m) }
func (i Inch) String() string       { return fmt.Sprintf("%g(in)", i) }
func (k Kilo) String() string       { return fmt.Sprintf("%g(kg)", k) }
func (p Pound) String() string      { return fmt.Sprintf("%g(lb)", p) }

//!-
