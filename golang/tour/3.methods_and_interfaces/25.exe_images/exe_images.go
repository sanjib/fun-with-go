package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

// Image is a struct
type Image struct {
	image *image.RGBA
}

// ColorModel returns the Image's color model.
func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

// Bounds returns the domain for which At can return non-zero color.
func (i Image) Bounds() image.Rectangle {
	return i.image.Rect
}

// At returns the color of the pixel at (x, y).
func (i Image) At(x, y int) color.Color {
	return color.RGBA{0, 0, 255, 255}
}

func main() {
	m := Image{image.NewRGBA(image.Rect(0, 0, 100, 100))}
	pic.ShowImage(m)
}
