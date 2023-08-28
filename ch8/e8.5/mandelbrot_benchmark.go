// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 61.
//!+

// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"fmt"
	"image"
	"image/color"
	"math/cmplx"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

const (
	xmin, ymin, xmax, ymax = -2, -2, +2, +2
	width, height          = 4096, 4096
)

func main() {
	// In most cases, GOMAXPROCS workers is the best choice.
	// In a container, we should use the following method to get the correct CPU quota.
	// import _ "github.com/uber-go/automaxprocs"
	fmt.Printf("GOMAXPROCS is %d.\n", runtime.GOMAXPROCS(0))
	for i := 0; i < height; i++ {
		start := time.Now()
		rw := benchmark(i + 1)
		// If the workload is small, the more workers you have, the more workers you waste.
		// Change width and height to 1024 and try it.
		fmt.Printf("%d workers rendered in: %v, %d workers were not working\n", i+1, time.Since(start), i+1-rw)
	}
}

func benchmark(workers int) int {
	wg := &sync.WaitGroup{}
	rows := make(chan int, height)
	for row := 0; row < height; row++ {
		rows <- row
	}
	close(rows)
	var realworkers int64
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			gotWork := false
			for py := range rows {
				if !gotWork {
					gotWork = true
					// realworkers++ is not goroutine safe, see The Go Memory Model.
					atomic.AddInt64(&realworkers, 1)
				}
				y := float64(py)/height*(ymax-ymin) + ymin
				for px := 0; px < width; px++ {
					x := float64(px)/width*(xmax-xmin) + xmin
					z := complex(x, y)
					// Image point (px, py) represents complex value z.
					img.Set(px, py, mandelbrot(z))
				}
			}
			wg.Done()
		}()
	}
	wg.Wait()
	// png.Encode(os.Stdout, img) // NOTE: ignoring errors
	return int(realworkers)
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.RGBA{255 - contrast*n, contrast * n, 128 - contrast*n, 255}
		}
	}
	return color.Black
}

//!-

// Some other interesting functions:

func acos(z complex128) color.Color {
	v := cmplx.Acos(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{192, blue, red}
}

func sqrt(z complex128) color.Color {
	v := cmplx.Sqrt(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{128, blue, red}
}

// f(x) = x^4 - 1
//
// z' = z - f(z)/f'(z)
//
//	= z - (z^4 - 1) / (4 * z^3)
//	= z - (z - 1/z^3) / 4
func newton(z complex128) color.Color {
	const iterations = 37
	const contrast = 7
	for i := uint8(0); i < iterations; i++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(z*z*z*z-1) < 1e-6 {
			return color.Gray{255 - contrast*i}
		}
	}
	return color.Black
}
