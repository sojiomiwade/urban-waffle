package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	http.HandleFunc("/", handler) // each request calls handler
	log.Fatal(http.ListenAndServe("localhost:8001", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	cycles := 5
	for k, v := range r.Form {
		fmt.Printf("Form[%q] = %q\n", k, v)
	}
	if scycles := r.Form["cycles"]; scycles != nil {
		sscycles := strings.Join(scycles, "-")
		cycles, _ = strconv.Atoi(sscycles)
	}
	nframes := 64
	if snframes := r.Form["nframes"]; snframes != nil {
		ssnframes := strings.Join(snframes, "-")
		nframes, _ = strconv.Atoi(ssnframes)
	}
	fmt.Printf("cycles: %d, nframes: %d\n", cycles, nframes)
	lissajous(w, cycles, nframes)
	fmt.Println("\n\n")
}

// RGBA colors courtesy of https://www.color-hex.com/color-palette/1017391
// RGBA colors below are: red orange yellow green blue
var palette = []color.Color{
	color.White,
	color.RGBA{204, 0, 0, 0xff},
	color.RGBA{230, 145, 56, 0xff},
	color.RGBA{241, 194, 50, 0xff},
	color.RGBA{106, 168, 79, 0xff},
	color.RGBA{61, 133, 198, 0xff},
}

func lissajous(out io.Writer, cycles int, nframes int) {
	const (
		// cycles  = 5    //number of complete x oscillator revolutions
		res  = .001 //   angular resolution
		size = 100  //  image canvas covers [-size..+size]
		// nframes = 64   // number of animation frames
		delay = 8 //delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 //relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0                                 //phase difference
	width, sliceCount := nframes, len(palette)-1 //finish point
	coloridx, mark := -1, 0
	sliceWidth := width / sliceCount
	for i := 0; i < nframes; i++ {
		if i >= mark {
			mark += sliceWidth
			coloridx++
		}
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		/**
		  0      5      10      15      20      25<--- 3 colors needed for four slices
		  26 frames
		  3 colors
		  26/3 = 8
		  0-7|8-15|16-23|24-25
		  sliceCount can be 10
		  coloridx = 0
		  initialize mark=sliceWidth=width/sliceCount

		  loop on t, using coloridx
		  once you go past the mark, (that is t goes past mark), add sliceWidth to mark; inc coloridx

		*/
		for t := 0.0; t < 2*math.Pi*float64(cycles); t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(
				size+int(x*size+.5),
				size+int(y*size+.5),
				uint8(coloridx%sliceCount+1),
			)
		}
		phase += .1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) //NOTE: ignoring encoding errors
}
