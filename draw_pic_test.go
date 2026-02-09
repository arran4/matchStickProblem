package main

import (
	"image"
	"image/draw"
	"testing"
)

func TestDrawPic(t *testing.T) {
	// Table-driven tests for drawPic
	tests := []struct {
		name     string
		input    []bool
		checkPts []struct {
			x, y     int
			expected bool // true if pixel should be colored (non-background)
		}
	}{
		{
			name:  "Empty Input",
			input: []bool{},
			checkPts: []struct{ x, y int; expected bool }{
				{0, 0, false},
				{50, 50, false},
			},
		},
		{
			name:  "Segment 0 (Top)",
			input: []bool{true, false, false, false, false, false, false},
			checkPts: []struct{ x, y int; expected bool }{
				{70, 15, true},   // Middle of top segment
				{15, 70, false},  // Middle of top-left segment (inactive)
				{70, 125, false}, // Middle of middle segment (inactive)
			},
		},
		{
			name:  "Segment 1 (Top Left)",
			input: []bool{false, true, false, false, false, false, false},
			checkPts: []struct{ x, y int; expected bool }{
				{15, 70, true},   // Middle of top-left segment
				{70, 15, false},  // Middle of top segment (inactive)
			},
		},
		{
			name:  "Segment 6 (Bottom)",
			input: []bool{false, false, false, false, false, false, true},
			checkPts: []struct{ x, y int; expected bool }{
				{70, 235, true},  // Middle of bottom segment
				{70, 15, false},  // Top segment inactive
			},
		},
		{
			name:  "All Segments (8)",
			input: []bool{true, true, true, true, true, true, true},
			checkPts: []struct{ x, y int; expected bool }{
				{70, 15, true},   // Top
				{15, 70, true},   // Top Left
				{125, 70, true},  // Top Right
				{70, 125, true},  // Middle
				{15, 180, true},  // Bottom Left
				{125, 180, true}, // Bottom Right
				{70, 235, true},  // Bottom
			},
		},
		{
			name: "Multiple Digits",
			// First digit: Top segment (index 0)
			// Second digit: Top segment (index 7)
			input: []bool{
				true, false, false, false, false, false, false, // Digit 1
				true, false, false, false, false, false, false, // Digit 2
			},
			checkPts: []struct{ x, y int; expected bool }{
				{70, 15, true},    // Digit 1 Top
				{200, 15, true},   // Digit 2 Top (70 + 120 + 10)
				{70, 125, false},  // Digit 1 Middle (inactive)
				{200, 125, false}, // Digit 2 Middle (inactive)
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Dimensions:
			// digitWidth = 10*2 + 100*1 = 120
			// digitHeight = 10*3 + 100*2 = 230
			// We add margins: marginWidth=10, marginHeight=10
			// Wait, drawPic uses marginWidth as starting offset.
			// Is the image size constrained inside drawPic? No.
			// We should make the image large enough.
			numDigits := (len(tc.input) + 6) / 7
			if numDigits == 0 {
				numDigits = 1
			}
			w := numDigits*(digitWidth+spacing) + marginWidth*2
			h := digitHeight + marginHeight*2
			img := image.NewRGBA(image.Rect(0, 0, w, h))

			// Fill with background (Black)
			draw.Draw(img, img.Bounds(), &image.Uniform{backgroundColour}, image.Point{}, draw.Src)

			err := drawPic(tc.input, img)
			if err != nil {
				t.Fatalf("drawPic failed: %v", err)
			}

			bgR, bgG, bgB, _ := backgroundColour.RGBA()

			for _, pt := range tc.checkPts {
				c := img.At(pt.x, pt.y)
				r, g, b, _ := c.RGBA()

				isBackground := (r == bgR && g == bgG && b == bgB)

				if pt.expected && isBackground {
					t.Errorf("Expected pixel at (%d, %d) to be colored, but it was background", pt.x, pt.y)
				}
				if !pt.expected && !isBackground {
					t.Errorf("Expected pixel at (%d, %d) to be background, but it was colored (%d, %d, %d)", pt.x, pt.y, r>>8, g>>8, b>>8)
				}
			}
		})
	}
}
