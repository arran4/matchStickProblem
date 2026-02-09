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

import "testing"

func BenchmarkIsANumber(b *testing.B) {
	input := []bool{
		// 1
		false,
		false, true,
		false,
		false, true,
		false,
		// 2
		true,
		false, true,
		true,
		true, false,
		true,
		// 3
		true,
		false, true,
		true,
		false, true,
		true,
		// 11
		false,
		true, true,
		false,
		true, true,
		false,
		// 8
		true,
		true, true,
		true,
		true, true,
		true,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		isANumber(input)
	}
}
