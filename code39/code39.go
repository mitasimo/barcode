package code39

import (
	"errors"
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/mitasimo/barcode"
)

// Barcode is the specialization of the barcode.Barcode for code39
type Barcode struct {
	*barcode.Barcode
}

// New create new Barcode
func New(data string, width, heigth int) (*Barcode, error) {
	const minWidth = 1

	data = strings.Join([]string{"*", data, "*"}, "") // add start and stop symbols

	// count width of narrow and wide bars
	numRunes := utf8.RuneCountInString(data)
	totalNarrow := numRunes * (totalNarrowPerSymb + 1) // width of all symbols in narrow bars
	narrowWidth := width / totalNarrow                 // width in pixels of narrow bar
	wideWidth := narrowWidth * wideRatio               // width in pixels of wide bar
	barcodeWidth := totalNarrow * minWidth

	if narrowWidth < minWidth {
		return nil, fmt.Errorf("can't draw a %d characters barcode in a %d pixels wide image, need minimum %d pixels", numRunes, width, barcodeWidth)
	}

	bc := barcode.New(width, heigth)

	offset := (width - barcodeWidth) / 2
	for _, curRune := range data {
		isBar := true // starts with black bar
		curPattern, ok := patterns[curRune]
		if !ok {
			return nil, fmt.Errorf("unsupported character in data: %s", string(curRune))
		}
		for _, moduleType := range curPattern {
			var curWidth int
			if isNarrowModule(moduleType) {
				curWidth = narrowWidth
			} else {
				curWidth = wideWidth
			}
			if isBar {
				bc.AddModule(offset, offset+curWidth)
			}
			offset += curWidth
			isBar = !isBar
		}
		offset += narrowWidth // add offset between 2 symbols
	}
	return &Barcode{Barcode: bc}, nil
}

const (
	narrowPerSymb = 6 // number narrow bars per symbol
	widePerSymb   = 3 // number wide bars per symbol
)

var (
	patterns           map[rune]string
	wideRatio          int // narrow bars per wide bar
	totalNarrowPerSymb int // symbols width in narrow bars
)

func isNarrowModule(modType rune) bool {
	if modType == 'n' {
		return true
	}
	return false
}

// SetWideRatio set the default wide ratio
func SetWideRatio(r int) error {
	if r < 2 {
		return errors.New("wide ratio should be equal or grather then 2")
	}
	wideRatio = r
	totalNarrowPerSymb = narrowPerSymb + widePerSymb*wideRatio
	return nil
}

func init() {

	wideRatio = 3
	totalNarrowPerSymb = narrowPerSymb + widePerSymb*wideRatio

	patterns = map[rune]string{
		'1': "wnnwnnnnw", // w - wide bar, n - narrow bar
		'2': "nnwwnnnnw",
		'3': "wnwwnnnnn",
		'4': "nnnwwnnnw",
		'5': "wnnwwnnnn",
		'6': "nnwwwnnnn",
		'7': "nnnwnnwnw",
		'8': "wnnwnnwnn",
		'9': "nnwwnnwnn",
		'0': "nnnwwnwnn",
		'A': "wnnnnwnnw",
		'B': "nnwnnwnnw",
		'C': "wnwnnwnnn",
		'D': "nnnnwwnnw",
		'E': "wnnnwwnnn",
		'F': "nnwnwwnnn",
		'G': "nnnnnwwnw",
		'H': "wnnnnwwnn",
		'I': "nnwnnwwnn",
		'J': "nnnnwwwnn",
		'K': "wnnnnnnww",
		'L': "nnwnnnnww",
		'M': "wnwnnnnwn",
		'N': "nnnnwnnww",
		'O': "wnnnwnnwn",
		'P': "nnwnwnnwn",
		'Q': "nnnnnnwww",
		'R': "wnnnnnwwn",
		'S': "nnwnnnwwn",
		'T': "nnnnwnwwn",
		'U': "wwnnnnnnw",
		'V': "nwwnnnnnw",
		'W': "wwwnnnnnn",
		'X': "nwnnwnnnw",
		'Y': "wwnnwnnnn",
		'Z': "nwwnwnnnn",
		'-': "nwnnnnwnw",
		'.': "wwnnnnwnn",
		' ': "nwwnnnwnn",
		'*': "nwnnwnwnn",
		'$': "nwnwnwnnn",
		'/': "nwnwnnnwn",
		'+': "nwnnnwnwn",
		'%': "nnnwnwnwn",
	}

}
