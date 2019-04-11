package barcode

import (
	"fmt"
	"image"
	"image/color"
)

// New create new Barcode
func New(width, heigth int) *Barcode {
	return &Barcode{width: width, heigth: heigth}
}

// Barcode contains data to draw image of 2D barcode
type Barcode struct {
	width, heigth int      // image size
	modules       []module // only bar module
}

// AddModule - add new bar module to the Barcode
func (b *Barcode) AddModule(begin, end int) error {
	if begin >= b.width {
		return fmt.Errorf("Module begin (%d) grather then barcode width (%d)", begin, b.width)
	}
	if end >= b.width {
		return fmt.Errorf("Module end (%d) grather then barcode width (%d)", end, b.width)
	}
	if begin >= end {
		return fmt.Errorf("Module begin (%d) grather then module end (%d)", begin, end)
	}
	b.modules = append(b.modules, module{begin, end})

	return nil
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
