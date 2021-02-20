// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 58.
//!+

// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"io"
	"math"
	"net/http"
	"os"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

type zFunc func(x, y float64) float64

func svg(w io.Writer, f zFunc, r *http.Request) {
	// http://localhost:8000/?shape=eggbox
	// http://localhost:8000/?shape=saddle
	if r != nil {
		f = eggbox
		if err := r.ParseForm(); err == nil {
			for k, v := range r.Form {
				fmt.Fprintf(os.Stdout, "Form[%q] = %q\n", k, v)
				if k == "shape" {
					switch v[0] {
					case "eggbox":
						f = eggbox
					case "saddle":
						f = saddle
					}
				}
			}
		}
	}
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, az := corner(i+1, j, f)
			bx, by, bz := corner(i, j, f)
			cx, cy, cz := corner(i, j+1, f)
			dx, dy, dz := corner(i+1, j+1, f)
			if az < 0 && bz < 0 && cz < 0 && dz < 0 {
				fmt.Fprintf(w, "<polygon style='fill: blue' points='%g,%g %g,%g %g,%g %g,%g'/>\n",
					ax, ay, bx, by, cx, cy, dx, dy)
			} else {
				fmt.Fprintf(w, "<polygon style='fill: red' points='%g,%g %g,%g %g,%g %g,%g'/>\n",
					ax, ay, bx, by, cx, cy, dx, dy)
			}
		}
	}
	fmt.Fprintln(w, "</svg>")
}

func corner(i, j int, f zFunc) (float64, float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)
	if math.IsNaN(z) || math.IsInf(z, 1) || math.IsInf(z, -1) {
		z = 0
	}

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z
}

func eggbox(x, y float64) float64 {
	return -0.1 * (math.Cos(x) + math.Cos(y))
}

func saddle(x, y float64) float64 {
	a := 30.0
	b := 15.0
	a2 := a * a
	b2 := b * b
	return (y*y/a2 - x*x/b2)
}

func main() {
	var f zFunc
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: surface eggbox|saddle|web")
		os.Exit(1)
	}
	switch os.Args[1] {
	case "eggbox":
		f = eggbox
	case "saddle":
		f = saddle
	case "web":
		handler := func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "image/svg+xml")
			svg(w, nil, r)
		}
		http.HandleFunc("/", handler)
		if err := http.ListenAndServe("localhost:8000", nil); err != nil {
			fmt.Fprintf(os.Stderr, "surface: http.ListenAndServe returns error, %v\n", err)
			os.Exit(1)
		}
		os.Exit(0)
	default:
		fmt.Fprintln(os.Stderr, "usage: surface eggbox|saddle|web")
		os.Exit(1)
	}
	svg(os.Stdout, f, nil)
}

//!-
