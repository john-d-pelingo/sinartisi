package color

import (
	"math/rand"
	"time"
)

// RGB represents the Reg Green Blue spectrum
type RGB struct {
	Red, Green, Blue uint8
}

// RGBColorNumber transforms the RGB struct into its corresponding number
func RGBColorNumber(rgb *RGB) int {
	return int(rgb.Red)*0x10000 + int(rgb.Green)*0x100 + int(rgb.Blue)
}

// RandomRGBColorNumber generate an int between 0x000 and 0xFFF which represents an RGB color.
func RandomRGBColorNumber() int {
	const White = 0xFFFFFF

	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	return random.Intn(White + 1)
}
