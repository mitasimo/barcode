package barcode

import (
	"image"
	"image/color"
)

// New create new Barcode
func New(width, heigth int) *Barcode {
	return &Barcode{width: width, heigth: heigth}
}

// Barcode consist of data to draw image of 2D barcode
type Barcode struct {
	width, heigth int      // image size
	modules       []module // only bar module
}

// AddModule - add new bar module to Barcode
func (b *Barcode) AddModule(begin, end int) {
	b.modules = append(b.modules, module{begin, end})
}

// ColorModel is the implementation of the function interface image.Image
func (b *Barcode) ColorModel() color.Model {
	return color.RGBAModel
}

// Bounds is the implementation of the function interface image.Image
func (b *Barcode) Bounds() image.Rectangle {
	return image.Rect(0, 0, b.width, b.heigth)
}

// At is the implementation of the function interface image.Image
func (b *Barcode) At(x, y int) color.Color {
	for _, m := range b.modules {
		if m.begin <= x && x < m.end { // m.end out of range
			return colorBlack
		}
	}
	return colorWhite
}

var (
	colorWhite = color.RGBA{255, 255, 255, 0}
	colorBlack = color.RGBA{0, 0, 0, 255}
)

type module struct {
	begin, end int // range on x-axis
}
