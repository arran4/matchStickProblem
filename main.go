package main

import (
	"flag"
	"fmt"
	"github.com/golang/freetype"
	"golang.org/x/image/font"
	"golang.org/x/image/font/inconsolata"
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"log"
	"math"
	"os"
	"sort"
	"time"
)

const (
	multiplier      = 10
	matchWidth      = 1 * multiplier
	matchHeadLength = 1 * multiplier
	matchLength     = 10 * multiplier
	digitHeight     = matchWidth*3 + matchLength*2
	digitWidth      = matchWidth*2 + matchLength*1
	marginHeight    = matchWidth
	marginWidth     = matchWidth
	spacing         = matchWidth
)

var (
	backgroundColour = color.Black
	matchColour      = color.RGBA{0xA5, 0x2A, 0x2A, math.MaxUint8}
	matchHeadColour  = color.RGBA{255, 0, 0, math.MaxUint8}
	outfn            = flag.String("out", fmt.Sprintf("out-%d.gif", time.Now().Unix()), "output filename")
)

func drawMatch(img draw.Image, x, y int, leftRight bool) error {
	xlim := matchWidth
	for i := 0; i < (matchWidth * matchHeadLength); i++ {
		img.Set(x+(i%xlim), y+(i/xlim), matchHeadColour)
	}
	mlim := matchLength - matchHeadLength
	xOff := matchHeadLength
	yOff := 0
	if !leftRight {
		mlim = matchWidth
		xOff, yOff = yOff, xOff
	}
	for i := 0; i < (matchWidth * (matchLength - matchHeadLength)); i++ {
		img.Set(x+(i%mlim)+xOff, y+(i/mlim)+yOff, matchColour)
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
		case 1, 2:
			y += matchWidth
		}
		err := drawMatch(img, x, y, left)
		if err != nil {
			return err
		}
	}
	return nil
}

func countthem(a []bool) (t int, f int) {
	for _, e := range a {
		if e {
			t++
		} else {
			f++
		}
	}
	return
}

func findthem(a []bool) (t []int, f []int) {
	for i, e := range a {
		if e {
			t = append(t, i)
		} else {
			f = append(f, i)
		}
	}
	return
}

func isADigit(a []bool) (int, int, bool) {
	switch {
	case a[0] && a[1] && a[2] && a[3] && a[4] && a[5] && a[6]:
		return 8, 1, true
	case a[0] && a[1] && !a[2] && a[3] && a[4] && a[5] && a[6]:
		return 6, 1, true
	case a[0] && a[1] && a[2] && !a[3] && a[4] && a[5] && a[6]:
		return 0, 1, true
	case a[0] && a[1] && a[2] && a[3] && !a[4] && a[5] && a[6]:
		return 9, 1, true
	case a[0] && a[1] && a[2] && a[3] && !a[4] && a[5] && !a[6]:
		return 9, 1, true
	case a[0] && !a[1] && a[2] && !a[3] && !a[4] && a[5] && !a[6]:
		return 7, 1, true
	case a[0] && a[1] && !a[2] && a[3] && !a[4] && a[5] && a[6]:
		return 5, 1, true
	case !a[0] && a[1] && a[2] && a[3] && !a[4] && a[5] && !a[6]:
		return 4, 1, true
	case a[0] && !a[1] && a[2] && a[3] && !a[4] && a[5] && a[6]:
		return 3, 1, true
	case a[0] && !a[1] && a[2] && a[3] && a[4] && !a[5] && a[6]:
		return 2, 1, true
	case !a[0] && a[1] && !a[2] && !a[3] && a[4] && !a[5] && !a[6]:
		return 1, 1, true
	case !a[0] && !a[1] && a[2] && !a[3] && !a[4] && a[5] && !a[6]:
		return 1, 1, true
	case !a[0] && a[1] && a[2] && !a[3] && a[4] && a[5] && !a[6]:
		return 11, 2, true
	case !a[0] && !a[1] && !a[2] && !a[3] && !a[4] && !a[5] && !a[6]:
		return 0, 0, true
	}
	return 0, 0, false
}

func isANumber(a []bool) (int, bool) {
	n := 0
	totalDigits := 0
	for i := 0; i < len(a); i += 7 {
		val, digits, ok := isADigit(a[i : i+7])
		if !ok {
			return 0, false
		}
		if digits > 0 {
			for k := 0; k < digits; k++ {
				n *= 10
			}
			n += val
			totalDigits += digits
		}
	}
	if totalDigits == 0 {
		return 0, false
	}
	return n, true
}

func main() {
	start := time.Now()
	flag.Parse()
	initial := []bool{
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

		false,
		false, false,
		false,
		false, false,
		false,
	}
	outf, err := os.Create(*outfn)
	if err != nil {
		log.Panicf("%v", err)
	}

	fontSize, _ := font.BoundString(inconsolata.Regular8x16, "01234\n56789")

	digitBase := digitHeight*1 + marginHeight*2
	r := image.Rect(0, 0, digitWidth*len(initial)/7+spacing*3+marginWidth*2, digitBase+fontSize.Max.Y.Ceil())
	p := color.Palette{
		backgroundColour,
		matchColour,
		matchHeadColour,
		color.White,
	}
	img := image.NewPaletted(r, p)
	for i := 0; i < img.Bounds().Dy()*img.Bounds().Dx(); i++ {
		img.Set(i%img.Bounds().Dx(), i/img.Bounds().Dx(), backgroundColour)
	}
	err = drawPic(initial, img)
	if err != nil {
		log.Panicf("%v", err)
	}
	notfree, free := countthem(initial)
	permutations := free * notfree * (free - 1) * (notfree - 1)
	log.Printf("Permutations: %d", permutations)

	delay := 10

	g := gif.GIF{
		Delay: []int{delay},
		Image: []*image.Paletted{img},
	}

	found := []int{}
	foundat := []int{}
	sortedList := []int{}
	last := 0
	top5 := ""

	if n, ok := isANumber(initial); ok {
		log.Printf("Got number %d (initial) ", n)
		found = append(found, n)
		sortedList = append(sortedList, n)
		foundat = append(foundat, -1)
		last = n
		top5 = fmt.Sprintf("%d", n)
		d := &font.Drawer{
			Face: inconsolata.Regular8x16,
			Dot:  freetype.Pt(0, digitBase),
			Src:  image.White,
			Dst:  img,
		}
		d.DrawString(fmt.Sprintf("Last: %d   Best 5: %s", last, top5))
	}

	nonfreePos, freePos := findthem(initial)

	for i := 0; i < permutations; i++ {
		mutate := make([]bool, len(initial))
		copy(mutate, initial)

		move1 := i % (free * notfree)
		move1To := move1 % free
		move1From := move1 / free
		move2 := (i / (free * notfree)) % ((free - 1) * (notfree - 1))
		move2To := move2 % (free - 1)
		move2From := move2 / (free - 1)

		if move2To >= move1To {
			move2To += 1
		}
		if move2From >= move1From {
			move2From += 1
		}

		mutate[nonfreePos[move1From]] = false
		mutate[freePos[move1To]] = true
		mutate[nonfreePos[move2From]] = false
		mutate[freePos[move2To]] = true

		if n, ok := isANumber(mutate); ok {
			last = n
			log.Printf("Got number %d at: %.0f%% %d/%d", n, float64(i)/float64(permutations)*100, i, permutations)
			found = append(found, n)
			foundat = append(foundat, i)

			if a := sort.SearchInts(sortedList, n); len(sortedList) <= a || sortedList[a] != n {
				sortedList = append(sortedList, n)
				sort.Ints(sortedList)
				top5 = ""
				for ii := 0; ii < int(math.Min(float64(5), float64(len(sortedList)))); ii++ {
					top5 = top5 + fmt.Sprintf("%d,", sortedList[len(sortedList)-1-ii])
				}
			}
		}

		img2 := image.NewPaletted(r, p)
		err = drawPic(mutate, img2)
		if err != nil {
			log.Panicf("%v", err)
		}
		d := &font.Drawer{
			Face: inconsolata.Regular8x16,
			Dot:  freetype.Pt(0, digitBase),
			Src:  image.White,
			Dst:  img2,
		}
		d.DrawString(fmt.Sprintf("Last: %d   Best 5: %s", last, top5))

		g.Image = append(g.Image, img2)
		g.Delay = append(g.Delay, delay)
	}

	log.Printf("Permutations done, generating gif")

	err = gif.EncodeAll(outf, &g)
	if err != nil {
		log.Panicf("%v", err)
	}

	log.Printf("Gif generated saving: %s", *outfn)

	err = outf.Close()
	if err != nil {
		log.Panicf("%v", err)
	}

	log.Printf("Done in %s", time.Now().Sub(start))
}
