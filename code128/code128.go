package code128

import (
	"fmt"
	"unicode/utf8"

	"github.com/mitasimo/barcode"
)

const (
	// named pattern
	patternStartA = 103
	patternStartB = 104
	patternStartС = 105
	patternStop   = 106
	patternStop2  = 108
)

// Barcode specialize Barcode
type Barcode struct {
	*barcode.Barcode
}

// New create code28 image
func New(data string, width, heigth int) (*Barcode, error) {
	const modulesPerChar = 11

	runeNumber := utf8.RuneCountInString(data)
	totalModules := (runeNumber+3)*modulesPerChar + 2 // start symbol + data symbols + check digit + stop symbol + two bar modules at the end
	moduleWidth := width / totalModules
	if moduleWidth < 1 {
		return nil, fmt.Errorf("the width (%d pixels) of the image is not enough to draw the barcode, need %d pixels", width, totalModules*moduleWidth)
	}

	offset := 0
	checkDigit := 0

	bc := barcode.New(width, heigth)

	// draw start symbol
	offset = addBlock(bc, patterns[patternStartB], offset, moduleWidth)

	// draw data symbols
	for i, symb := range data {
		patternNumber, ok := symbolsB[symb]
		if !ok {
			return nil, fmt.Errorf("symbol %s", string(symb))
		}
		offset = addBlock(bc, patterns[patternNumber], offset, moduleWidth)

		checkDigit += i * patternNumber
	}

	// draw check symbol
	checkDigit = checkDigit % 103
	offset = addBlock(bc, patterns[checkDigit], offset, moduleWidth)

	// draw stop symbol
	addBlock(bc, patterns[patternStop2], offset, moduleWidth)

	return &Barcode{Barcode: bc}, nil
}

func isSpace(r rune) bool {
	return r == '0'
}

func isBar(r rune) bool {
	return !isSpace(r)
}

func addBlock(bcode *barcode.Barcode, pattern string, offset, moduleWidth int) int {
	for _, moduleType := range pattern {
		if isBar(moduleType) {
			bcode.AddModule(offset, offset+moduleWidth)
		}
		offset += moduleWidth
	}
	return offset
}

// smb struct describes
// symbol index and module pattern
type smb struct {
	index   int
	pattern string
}

var (
	// symbolsB maps rune to pattern index
	symbolsB = map[rune]int{
		'\u0020': 0, //space
		'!':      1,
		'\u0022': 2, // double quote
		'#':      3,
		'$':      4,
		'%':      5,
		'&':      6,
		'\u0027': 7, //single quote
		'(':      8,
		')':      9,
		'*':      10,
		'+':      11,
		',':      12,
		'-':      13,
		'.':      14,
		'/':      15,
		'0':      16,
		'1':      17,
		'2':      18,
		'3':      19,
		'4':      20,
		'5':      21,
		'6':      22,
		'7':      23,
		'8':      24,
		'9':      25,
		':':      26,
		';':      27,
		'<':      28,
		'=':      29,
		'>':      30,
		'?':      31,
		'@':      32,
		'A':      33,
		'B':      34,
		'C':      35,
		'D':      36,
		'E':      37,
		'F':      38,
		'G':      39,
		'H':      40,
		'I':      41,
		'J':      42,
		'K':      43,
		'L':      44,
		'M':      45,
		'N':      46,
		'O':      47,
		'P':      48,
		'Q':      49,
		'R':      50,
		'S':      51,
		'T':      52,
		'U':      53,
		'V':      54,
		'W':      55,
		'X':      56,
		'Y':      57,
		'Z':      58,
		'[':      59,
		'\u005C': 60, // back slash
		']':      61,
		'^':      62, // CIRCUMFLEX
		'_':      63,
		'`':      64, // GRAVE
		'a':      65,
		'b':      66,
		'c':      67,
		'd':      68,
		'e':      69,
		'f':      70,
		'g':      7,
		'h':      7,
		'i':      7,
		'j':      7,
		'k':      7,
		'l':      7,
		'm':      7,
		'n':      7,
		'o':      7,
		'p':      80,
		'q':      8,
		'r':      8,
		's':      8,
		't':      8,
		'u':      8,
		'v':      8,
		'w':      8,
		'x':      8,
		'y':      8,
		'z':      90,
		'{':      91,
		'|':      92,
		'}':      93,
		'~':      94,
	}

	patterns = []string{
		"11011001100",   //	0
		"11001101100",   //	1
		"11001100110",   //	2
		"10010001100",   //	3
		"10010011000",   //	4
		"10001001100",   //	5
		"10011001000",   //	6
		"10011000100",   //	7
		"10001100100",   //	8
		"11001001000",   //	9
		"11001000100",   //	10
		"11000100100",   //	11
		"10110011100",   //	12
		"10011011100",   //	13
		"10011001110",   //	14
		"10111001100",   //	15
		"10011101100",   //	16
		"10011100110",   //	17
		"11001110010",   //	18
		"11001011100",   //	19
		"11001001110",   //	20
		"11011100100",   //	21
		"11001110100",   //	22
		"11101101110",   //	23
		"11101001100",   //	24
		"11100101100",   //	25
		"11100100110",   //	26
		"11101100100",   //	27
		"11100110100",   //	28
		"11100110010",   //	29
		"11011011000",   //	30
		"11011000110",   //	31
		"11000110110",   //	32
		"10100011000",   //	33
		"10001011000",   //	34
		"10001000110",   //	35
		"10110001000",   //	36
		"10001101000",   //	37
		"10001100010",   //	38
		"11010001000",   //	39
		"11000101000",   //	40
		"11000100010",   //	41
		"10110111000",   //	42
		"10110001110",   //	43
		"10001101110",   //	44
		"10111011000",   //	45
		"10111000110",   //	46
		"10001110110",   //	47
		"11101110110",   //	48
		"11010001110",   //	49
		"11000101110",   //	50
		"11011101000",   //	51
		"11011100010",   //	52
		"11011101110",   //	53
		"11101011000",   //	54
		"11101000110",   //	55
		"11100010110",   //	56
		"11101101000",   //	57
		"11101100010",   //	58
		"11100011010",   //	59
		"11101111010",   //	60
		"11001000010",   //	61
		"11110001010",   //	62
		"10100110000",   //	63
		"10100001100",   //	64
		"10010110000",   //	65
		"10010000110",   //	66
		"10000101100",   //	67
		"10000100110",   //	68
		"10110010000",   //	69
		"10110000100",   //	70
		"10011010000",   //	71
		"10011000010",   //	72
		"10000110100",   //	73
		"10000110010",   //	74
		"11000010010",   //	75
		"11001010000",   //	76
		"11110111010",   //	77
		"11000010100",   //	78
		"10001111010",   //	79
		"10100111100",   //	80
		"10010111100",   //	81
		"10010011110",   //	82
		"10111100100",   //	83
		"10011110100",   //	84
		"10011110010",   //	85
		"11110100100",   //	86
		"11110010100",   //	87
		"11110010010",   //	88
		"11011011110",   //	89
		"11011110110",   //	90
		"11110110110",   //	91
		"10101111000",   //	92
		"10100011110",   //	93
		"10001011110",   //	94
		"10111101000",   //	95
		"10111100010",   //	96
		"11110101000",   //	97
		"11110100010",   //	98
		"10111011110",   //	99
		"10111101110",   //	100
		"11101011110",   //	101
		"11110101110",   //	102
		"11010000100",   //	103
		"11010010000",   //	104
		"11010011100",   //	105
		"11000111010",   //	106
		"11010111000",   //	107
		"1100011101011", //	108
	}
)
