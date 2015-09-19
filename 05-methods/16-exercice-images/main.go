// Remember the picture generator you wrote earlier? Let's write another one, but this time it will return an implementation of image.Image instead of a slice of data.

// Define your own Image type, implement the necessary methods, and call pic.ShowImage.

// Bounds should return a image.Rectangle, like image.Rect(0, 0, w, h).

// ColorModel should return color.RGBAModel.

// At should return a color; the value v in the last picture generator corresponds to color.RGBA{v, v, 255, 255} in this one.package main

package main

import (
	"image"
	"image/color"
	"math/rand"

	"golang.org/x/tour/pic"
)

type Image struct {
	Height, Width int
}

func (p Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (p Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, p.Width, p.Height)
}

func (p Image) At(x, y int) color.Color {
	c := rand.Int()
	return color.RGBA{uint8(c), uint8(c), 255, 255}
}

func main() {
	m := Image{256, 256}
	pic.ShowImage(m)
}
