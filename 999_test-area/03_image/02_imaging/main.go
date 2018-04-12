package main

import (
	"image/color"
	"log"

	"github.com/disintegration/imaging"
)

func main() {
	// ROTATE AN IMAGE WITH GO
	img, err := imaging.Open("luf.jpg")
	if err != nil {
		log.Fatalln(err)
		return
	}
	// img, degrees, set a color to the background
	img = imaging.Rotate(img, 90, color.RGBA{0, 0, 0, 1})
	err = imaging.Save(img, "newLuf.jpg")
	if err != nil {
		log.Fatalln(err)
	}
}
