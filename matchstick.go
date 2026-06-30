package matchStickProblem

import (
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
	"strconv"
	"time"
)

type Orientation int

const (
	Horizontal Orientation = iota
	Vertical
	DiagonalForward
	DiagonalBackward
)

const (
	multiplier      = 10
	matchWidth      = 1 * multiplier
	matchHeadLength = 1 * multiplier
	matchLength     = 10 * multiplier
	digitHeight     = matchWidth*3 + matchLength*2
	digitWidth      = matchWidth*3 + matchLength*2
	marginHeight    = matchWidth
	marginWidth     = matchWidth
	spacing         = matchWidth

	segA1 = 0
	segA2 = 1
	segB  = 2
	segC  = 3
	segD1 = 4
	segD2 = 5
	segE  = 6
	segF  = 7
	segG1 = 8
	segG2 = 9
	segH  = 10
	segI  = 11
	segJ  = 12
	segM  = 13
	segL  = 14
	segK  = 15

	segCount = 16
)

var (
	backgroundColour = color.Black
	matchColour      = color.RGBA{0xA5, 0x2A, 0x2A, math.MaxUint8}
	matchHeadColour  = color.RGBA{255, 0, 0, math.MaxUint8}
	digitLookup      = map[int]string{
		255:  "0",
		12:   "1",
		887:  "2",
		831:  "3",
		908:  "4",
		955:  "5",
		1019: "6",
		15:   "7",
		1023: "8",
		959:  "9",
		204:  "11", // E, F, B, C (two vertical lines)
		0:    "",
	}
)

func drawMatch(img draw.Image, x, y int, o Orientation) error {
	switch o {
	case Horizontal, Vertical:
		leftRight := o == Horizontal
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
	case DiagonalForward, DiagonalBackward:
		for i := 0; i < matchLength; i++ {
			for j := 0; j < matchWidth; j++ {
				c := matchColour
				if i < matchHeadLength {
					c = matchHeadColour
				}
				px, py := 0, 0
				if o == DiagonalForward {
					// \ direction
					px = x + i + j
					py = y + i + (matchWidth - j)
				} else {
					// / direction
					px = x + i + j
					py = y - i + (matchWidth - j)
				}
				img.Set(px, py, c)
				// fill holes
				img.Set(px+1, py, c)
			}
		}
		return nil
	default:
		return fmt.Errorf("unsupported orientation: %v", o)
	}
}

func drawPic(input []bool, img draw.Image) error {
	for i, each := range input {
		if !each {
			continue
		}
		pos := i % segCount
		x := marginWidth
		x += (i / segCount) * (digitWidth + spacing)

		y := marginHeight
		var o Orientation

		switch pos {
		case segA1: // 0
			x += matchWidth
			o = Horizontal
		case segA2: // 1
			x += matchWidth + matchLength
			o = Horizontal
		case segB: // 2
			x += matchWidth + matchLength*2
			y += matchWidth
			o = Vertical
		case segC: // 3
			x += matchWidth + matchLength*2
			y += matchWidth*2 + matchLength
			o = Vertical
		case segD1: // 4
			x += matchWidth
			y += matchWidth*2 + matchLength*2
			o = Horizontal
		case segD2: // 5
			x += matchWidth + matchLength
			y += matchWidth*2 + matchLength*2
			o = Horizontal
		case segE: // 6
			y += matchWidth*2 + matchLength
			o = Vertical
		case segF: // 7
			y += matchWidth
			o = Vertical
		case segG1: // 8
			x += matchWidth
			y += matchWidth + matchLength
			o = Horizontal
		case segG2: // 9
			x += matchWidth + matchLength
			y += matchWidth + matchLength
			o = Horizontal
		case segH: // 10
			x += matchWidth
			y += matchWidth
			o = DiagonalForward
		case segI: // 11
			x += matchWidth + matchLength
			y += matchWidth
			o = Vertical
		case segJ: // 12
			x += matchWidth + matchLength
			y += matchWidth + matchLength
			o = DiagonalBackward
		case segM: // 13
			x += matchWidth
			y += matchWidth*2 + matchLength*2
			o = DiagonalBackward
		case segL: // 14
			x += matchWidth + matchLength
			y += matchWidth*2 + matchLength
			o = Vertical
		case segK: // 15
			x += matchWidth + matchLength
			y += matchWidth*2 + matchLength
			o = DiagonalForward
		}

		err := drawMatch(img, x, y, o)
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
	return t, f
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

func isADigit(a []bool) ([]byte, bool) {
	mask := 0
	for i, v := range a {
		if v {
			mask |= 1 << i
		}
	}
	if val, ok := digitLookup[mask]; ok {
		return []byte(val), true
	}
	return []byte{}, false
}

func isANumber(a []bool) (int, bool) {
	str := []byte{}
	for i := 0; i < len(a); i += segCount {
		if b, ok := isADigit(a[i : i+segCount]); !ok {
			return 0, false
		} else {
			str = append(str, b...)
		}
	}
	if i, err := strconv.ParseInt(string(str), 10, 64); err != nil {
		return 0, false
	} else {
		return int(i), true
	}
}

// Run is a subcommand `matchStickProblem run`
//
// Flags:
//
//	outfn: --out -out (default: "") output filename
func Run(outfn string) {
	start := time.Now()

	initial := []bool{
		false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false,
		true, true, false, true, true, true, false, true, true, true, false, false, false, false, false, false,
		true, true, true, true, true, true, true, true, false, false, false, false, false, false, false, false,
		true, true, true, true, true, true, true, true, true, true, false, false, false, false, false, false,
		false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false,
	}

	var outf *os.File
	var err error
	if outfn == "" {
		// Note: os.CreateTemp creates a persistent file that is not automatically deleted.
		outf, err = os.CreateTemp(".", "out-*.gif")
		if err != nil {
			log.Panicf("%v", err)
		}
		outfn = outf.Name()
	} else {
		outf, err = os.Create(outfn)
		if err != nil {
			log.Panicf("%v", err)
		}
	}

	fontSize, _ := font.BoundString(inconsolata.Regular8x16, "01234\n56789")

	digitBase := digitHeight*1 + marginHeight*2
	r := image.Rect(0, 0, digitWidth*len(initial)/segCount+spacing*3+marginWidth*2, digitBase+fontSize.Max.Y.Ceil())
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

	log.Printf("Gif generated saving: %s", outfn)

	err = outf.Close()
	if err != nil {
		log.Panicf("%v", err)
	}

	log.Printf("Done in %s", time.Now().Sub(start))
}
