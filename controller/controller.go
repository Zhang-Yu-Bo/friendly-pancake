package controller

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if len(name) == 0 {
		name = "stranger"
	}

	fmt.Fprintf(w, "Hello, %s", name)
}

func RawImage(w http.ResponseWriter, r *http.Request) {

	fileName := r.URL.Query().Get("file_name")

	if len(fileName) == 0 {
		fileName = "image.png"
	}

	width := 200
	height := 100

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	// Colors are defined by Red, Green, Blue, Alpha uint8 values.
	cyan := color.RGBA{100, 200, 200, 0xff}

	// Set color for each pixel.
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			switch {
			case x < width/2 && y < height/2: // upper left quadrant
				img.Set(x, y, cyan)
			case x >= width/2 && y >= height/2: // lower right quadrant
				img.Set(x, y, color.White)
			default:
				// Use zero value.
			}
		}
	}
	png.Encode(w, img)
	// fmt.Fprintf(w, "%s", fileName)

}
