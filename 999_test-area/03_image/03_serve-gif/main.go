package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		if c := q.Get("cycles"); c != "" {
			cy, err := strconv.Atoi(c)
			if err != nil {
				log.Fatal(err)
				return
			}
			lissajous(w, cy)
			return
		}
		lissajous(w, 5)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func lissajous(out io.Writer, c int) {
	pallete := []color.Color{color.White, color.Black}
	// Set the configs to gif
	cycles := float64(c)
	const (
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
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), 1)
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
