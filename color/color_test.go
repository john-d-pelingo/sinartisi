package color

import (
	"testing"
)

type rgbColorNumberPair struct {
	input  *RGB
	result int
}

func TestRGBColorNumber(t *testing.T) {
	var tests = []rgbColorNumberPair{
		{&RGB{0x55, 0x12, 0xAF}, 0x5512AF},
		{&RGB{0x0, 0, 0x9}, 0x000009},
		{&RGB{0x1, 0x2, 0x3}, 0x010203},
		{&RGB{0, 0, 0}, 0},
		{&RGB{255, 255, 255}, 0xFFFFFF},
	}

	for _, pair := range tests {
		got := RGBColorNumber(pair.input)
		want := pair.result

		if want != got {
			t.Errorf(`
				Expected: %d
				Received: %d
			`, want, got)
		}
	}
}
