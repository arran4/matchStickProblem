package main

import (
	"image"
	"image/color"
	"log"
	"testing"
)

func TestIsANumber(t *testing.T) {
	expected := []struct {
		b     int
		ok    bool
		input []bool
	}{
		{0, false, []bool{
			false,
			false, false,
			false,
			false, false,
			true,
		}},
		{0, false, []bool{
			false,
			false, false,
			false,
			false, false,
			false,
		}},
		{1, true, []bool{
			false,
			true, false,
			false,
			true, false,
			false,
		}},
		{1, true, []bool{
			false,
			false, true,
			false,
			false, true,
			false,
		}},
		{2, true, []bool{
			true,
			false, true,
			true,
			true, false,
			true,
		}},
		{3, true, []bool{
			true,
			false, true,
			true,
			false, true,
			true,
		}},
		{4, true, []bool{
			false,
			true, true,
			true,
			false, true,
			false,
		}},
		{5, true, []bool{
			true,
			true, false,
			true,
			false, true,
			true,
		}},
		{6, true, []bool{
			true,
			true, false,
			true,
			true, true,
			true,
		}},
		{7, true, []bool{
			true,
			false, true,
			false,
			false, true,
			false,
		}},
		{8, true, []bool{
			true,
			true, true,
			true,
			true, true,
			true,
		}},
		{9, true, []bool{
			true,
			true, true,
			true,
			false, true,
			true,
		}},
		{9, true, []bool{
			true,
			true, true,
			true,
			false, true,
			false,
		}},
		{0, true, []bool{
			true,
			true, true,
			false,
			true, true,
			true,
		}},
		{11, true, []bool{
			false,
			true, true,
			false,
			true, true,
			false,
		}},
		{1111, true, []bool{
			false,
			true, true,
			false,
			true, true,
			false,
			false,
			true, true,
			false,
			true, true,
			false,
		}},
	}
	for i, each := range expected {
		if b, ok := isANumber(each.input); b != each.b || ok != each.ok {
			log.Printf("Failed on #%d (expected %d) got (%d %v)", i, each.b, b, ok)
			t.Fail()
		}
	}
}

func TestIsADigit(t *testing.T) {
	expected := []struct {
		b     string
		ok    bool
		input []bool
	}{
		{"", false, []bool{
			false,
			false, false,
			false,
			false, false,
			true,
		}},
		{"", true, []bool{
			false,
			false, false,
			false,
			false, false,
			false,
		}},
		{"1", true, []bool{
			false,
			true, false,
			false,
			true, false,
			false,
		}},
		{"1", true, []bool{
			false,
			false, true,
			false,
			false, true,
			false,
		}},
		{"2", true, []bool{
			true,
			false, true,
			true,
			true, false,
			true,
		}},
		{"3", true, []bool{
			true,
			false, true,
			true,
			false, true,
			true,
		}},
		{"4", true, []bool{
			false,
			true, true,
			true,
			false, true,
			false,
		}},
		{"5", true, []bool{
			true,
			true, false,
			true,
			false, true,
			true,
		}},
		{"6", true, []bool{
			true,
			true, false,
			true,
			true, true,
			true,
		}},
		{"7", true, []bool{
			true,
			false, true,
			false,
			false, true,
			false,
		}},
		{"8", true, []bool{
			true,
			true, true,
			true,
			true, true,
			true,
		}},
		{"9", true, []bool{
			true,
			true, true,
			true,
			false, true,
			true,
		}},
		{"9", true, []bool{
			true,
			true, true,
			true,
			false, true,
			false,
		}},
		{"0", true, []bool{
			true,
			true, true,
			false,
			true, true,
			true,
		}},
		{"11", true, []bool{
			false,
			true, true,
			false,
			true, true,
			false,
		}},
	}
	for i, each := range expected {
		if b, ok := isADigit(each.input); string(b) != each.b || ok != each.ok {
			log.Printf("Failed on #%d (expected %s) got (%s %v)", i, each.b, b, ok)
			t.Fail()
		}
	}
}

func TestDrawMatch(t *testing.T) {
	// Setup
	r := image.Rect(0, 0, 200, 200)
	img := image.NewRGBA(r)

	// Helper to check pixel color
	checkPixel := func(x, y int, expected color.Color, name string) {
		got := img.At(x, y)
		r1, g1, b1, a1 := got.RGBA()
		r2, g2, b2, a2 := expected.RGBA()
		if r1 != r2 || g1 != g2 || b1 != b2 || a1 != a2 {
			t.Errorf("%s: pixel at (%d, %d) expected %v, got %v", name, x, y, expected, got)
		}
	}

	// 1. Horizontal Match (leftRight = true)
	// drawMatch(img, 10, 10, true)
	// Head: (10,10) to (19,19)
	// Stick: (20,10) to (109,19)
	err := drawMatch(img, 10, 10, true)
	if err != nil {
		t.Fatalf("drawMatch failed: %v", err)
	}

	// Check Head
	checkPixel(10, 10, matchHeadColour, "Horizontal Head TopLeft")
	checkPixel(19, 19, matchHeadColour, "Horizontal Head BottomRight")

	// Check Stick
	checkPixel(20, 10, matchColour, "Horizontal Stick Start")
	checkPixel(109, 19, matchColour, "Horizontal Stick End")

	// Check surroundings (cleanliness)
	checkPixel(9, 10, color.Transparent, "Horizontal Left Clean") // Left of head
	checkPixel(10, 9, color.Transparent, "Horizontal Top Clean") // Above head
	checkPixel(110, 10, color.Transparent, "Horizontal Right Clean") // Right of stick
	checkPixel(10, 20, color.Transparent, "Horizontal Bottom Clean") // Below head

	// 2. Vertical Match (leftRight = false)
	// drawMatch(img, 10, 50, false)
	// Head: (10,50) to (19,59)
	// Stick: (10,60) to (19,149)
	err = drawMatch(img, 10, 50, false)
	if err != nil {
		t.Fatalf("drawMatch failed: %v", err)
	}

	// Check Head
	checkPixel(10, 50, matchHeadColour, "Vertical Head TopLeft")
	checkPixel(19, 59, matchHeadColour, "Vertical Head BottomRight")

	// Check Stick
	checkPixel(10, 60, matchColour, "Vertical Stick Start")
	checkPixel(19, 149, matchColour, "Vertical Stick End")

	// Check surroundings
	checkPixel(9, 50, color.Transparent, "Vertical Left Clean")
	checkPixel(10, 49, color.Transparent, "Vertical Top Clean")
	checkPixel(10, 150, color.Transparent, "Vertical Bottom Clean")
	checkPixel(20, 50, color.Transparent, "Vertical Right Clean")
}
