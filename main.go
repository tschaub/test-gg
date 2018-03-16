package main

import (
	"image"
	"image/png"
	"os"

	"github.com/fogleman/gg"
)

func main() {
	file, _ := os.Open("tile.png")
	im, _ := png.Decode(file)

	rgba := im.(*image.RGBA)

	cropped := rgba.SubImage(image.Rect(30, 40, 246, 246))

	file, _ = os.Create("cropped.png")
	png.Encode(file, cropped)

	file, _ = os.Create("gg.png")
	png.Encode(file, gg.NewContextForImage(cropped).Image())
}
