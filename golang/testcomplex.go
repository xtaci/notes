package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math/cmplx"
	"os"
)

func mandelbrot(a complex128) (z complex128) {
	for i := 0; i < 50; i++ {
		z = z*z + a
	}
	return
}

func main() {
	g := image.NewGray(image.Rectangle{image.Point{0, 0}, image.Point{1000, 1000}})

	for y := 1.0; y >= -1.0; y -= 0.001 {
		for x := 1.0; x >= -1.0; x -= 0.001 {
			h := cmplx.Abs(mandelbrot(complex(x, y)))
			if h < 2 {
				g.Set(int(x*1000)+500, int(y*1000), color.Gray{255})
			}
		}
	}

	file, err := os.OpenFile("frac.png", os.O_TRUNC|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Println(err)
	}
	png.Encode(file, g)
}
