package main

import (
	"image"
	"image/gif"
	"testing"
)

func BenchmarkSliceAppend_WithoutPreAlloc(b *testing.B) {
	permutations := 83232
	delay := 10
	img := &image.Paletted{}

	for n := 0; n < b.N; n++ {
		// Simulation of the current code
		g := gif.GIF{
			Delay: []int{delay},
			Image: []*image.Paletted{img},
		}

		for i := 0; i < permutations; i++ {
			g.Image = append(g.Image, img)
			g.Delay = append(g.Delay, delay)
		}
	}
}

func BenchmarkSliceAppend_WithPreAlloc(b *testing.B) {
	permutations := 83232
	delay := 10
	img := &image.Paletted{}

	for n := 0; n < b.N; n++ {
		// Simulation of the optimized code
		g := gif.GIF{
			Delay: make([]int, 0, permutations+1),
			Image: make([]*image.Paletted, 0, permutations+1),
		}
		// Initial append
		g.Delay = append(g.Delay, delay)
		g.Image = append(g.Image, img)

		for i := 0; i < permutations; i++ {
			g.Image = append(g.Image, img)
			g.Delay = append(g.Delay, delay)
		}
	}
}
