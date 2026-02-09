package main

import (
	"image"
	"image/color"
	"testing"
)

func BenchmarkDrawMatch(b *testing.B) {
	// Setup
	r := image.Rect(0, 0, 1000, 1000)
	p := color.Palette{
		backgroundColour,
		matchColour,
		matchHeadColour,
		color.White,
	}
	img := image.NewPaletted(r, p)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		drawMatch(img, 10, 10, true)
		drawMatch(img, 100, 100, false)
	}
}
