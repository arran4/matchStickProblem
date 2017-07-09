package main

import (
	"image"
	"image/gif"
	"os"
	"fmt"
	"time"
	"log"
	"image/color"
	"image/draw"
	"math"
)

const (
	multiplier      = 10
	matchWidth      = 1 * multiplier
	matchHeadLength = 1 * multiplier
	matchLength     = 10 * multiplier
	digitHeight     = matchWidth * 3 + matchLength * 2
	digitWidth      = matchWidth * 2 + matchLength * 1
	marginHeight    = matchWidth
	marginWidth     = matchWidth
	spacing         = matchWidth
)

var (
	backgroundColour = color.Black
	matchColour = color.RGBA{0xA5,0x2A,0x2A, math.MaxUint8,}
	matchHeadColour = color.RGBA{255,0,0, math.MaxUint8,}
)

func drawMatch(img draw.Image, x,y int, leftRight bool) error {
	xlim := matchWidth
	for i := 0; i < (matchWidth * matchHeadLength); i++ {
		img.Set(x + (i % xlim), y + (i / xlim), matchHeadColour)
	}
	mlim := (matchLength - matchHeadLength)
	xOff := matchHeadLength
	yOff := 0
	if !leftRight {
		mlim = matchWidth
		xOff, yOff = yOff, xOff
	}
	for i := 0; i < (matchWidth * (matchLength - matchHeadLength)); i++ {
		img.Set(x + (i % mlim) + xOff, y + (i / mlim) + yOff, matchColour)
	}
	return nil
}

func drawPic(input []bool, img draw.Image) error {
	for i, each := range input {
		if !each {
			continue
		}
		pos := i % 7
		x := marginWidth
		x += (i / 7) * (digitWidth + spacing)
		switch pos {
		case 1, 4:
		case 2, 5:
			x += matchLength
			fallthrough
		case 0, 3, 6:
			x += matchWidth
		}
		y := marginHeight
		left := pos == 0 || pos == 3 || pos == 6
		switch pos {
		case 6:
			y += matchLength
			fallthrough
		case 4, 5:
			y += matchWidth
			fallthrough
		case 3:
			y += matchLength
			fallthrough
		case 1,2:
			y += matchWidth
		}
		err := drawMatch(img, x,y, left)
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	initial := []bool {
		false,
		false, false,
		false,
		false, false,
		false,

		true,
		true, false,
		true,
		false, true,
		true,

		true,
		true, true,
		false,
		true, true,
		true,

		true,
		true, true,
		true,
		true, true,
		true,

	}
	outf, err := os.Create(fmt.Sprintf("out-%d.gif", time.Now().Unix()))
	if err != nil {
		log.Panicf("%v", err)
	}
	r := image.Rect(0,0,digitWidth * 4 + spacing * 3 + marginWidth * 2,digitHeight * 1 + marginHeight * 2)
	p := color.Palette{
		backgroundColour,
		matchColour,
		matchHeadColour,
		//color.White,
	}
	img := image.NewPaletted(r, p)
	for i := 0; i < img.Bounds().Dy() * img.Bounds().Dx(); i++ {
		img.Set((i % img.Bounds().Dx()), (i / img.Bounds().Dx()), backgroundColour)
	}

	err = drawPic(initial, img)
	if err != nil {
		log.Panicf("%v", err)
	}
	o := &gif.Options{
		NumColors: len(p),
		Drawer: draw.FloydSteinberg,
	}

	err = gif.Encode(outf, img, o)
	if err != nil {
		log.Panicf("%v", err)
	}


	err = outf.Close()
	if err != nil {
		log.Panicf("%v", err)
	}
}