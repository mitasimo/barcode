package barcode

import (
	"image"
	"testing"
)

func TestAt(t *testing.T) {
	b := New(100, 100)
	b.AddModule(0, 26)
	b.AddModule(51, 78)

	if b.At(0, 0) != colorBlack {
		t.Fatalf("The color at the point (0, 0) should be black")
	}
	if b.At(0, 5) != colorBlack {
		t.Fatalf("The color at the point (0, 5) should be black")
	}
	if b.At(60, 1000) != colorBlack {
		t.Fatalf("The color at the point (60, 1000) should be black")
	}
	if b.At(26, 5) != colorWhite {
		t.Fatalf("The color at point (26, 5) should be white")
	}
	if b.At(1000, 1000) != colorWhite {
		t.Fatalf("The color at point (1000, 1000) should be white")
	}
}

func TestAddModule(t *testing.T) {
	var err error

	b := New(100, 100)
	err = b.AddModule(10, 20)
	if err != nil {
		t.Fatalf("AddModule(10, 20) with Barcode(100,100) returns error but should return nil")
	}
	err = b.AddModule(99, 102)
	if err == nil {
		t.Fatalf("AddModule(99, 102) with Barcode(100,100) returns nil but should return error")
	}
	err = b.AddModule(66, 55)
	if err == nil {
		t.Fatalf("AddModule(66, 55) with Barcode(100,100) returns nil but should return error because begin(66) grather then end(55)")
	}
}

func TestBound(t *testing.T) {
	br := New(100, 100).Bounds()
	r := image.Rect(0, 0, 100, 100)
	if br != r {
		t.Fatalf("Barcode image bounds is %v, but must be %v", br, r)
	}
}
