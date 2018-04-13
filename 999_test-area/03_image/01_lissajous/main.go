package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"os"
	"time"
)

var pallete = []color.Color{color.White, color.Black, color.RGBA{54, 236, 12, 1}, color.RGBA{12, 39, 136, 1}, color.RGBA{136, 12, 12, 1}, color.RGBA{136, 12, 132, 1}}

const (
	white  = 0
	black  = 1
	green  = 2
	blue   = 3
	red    = 4
	purple = 5
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	lissajous(os.Stdout) // ./main.go >out01.gif
}

func lissajous(out io.Writer) {
	// Set the configs to gif
	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		react := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(react, pallete)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			cl := uint8(rand.Intn(2) + 2)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), cl)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	err := gif.EncodeAll(out, &anim)
	if err != nil {
		log.Println("ERRRRRRRRRRRRRRRRRRRRRRRRRRRR: ", err)
		return
	}
}
