package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct{
    W, H int
}

func (p Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (p Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, p.W, p.H)
}

func (p Image) At(x, y int) color.Color {
	v := uint8((x - y) / 2)
	return color.RGBA{R: v, G: v, B: 255, A: 255}
}

func cShowImage() {
	m := Image{W: 150, H: 150}
	pic.ShowImage(m)
}
