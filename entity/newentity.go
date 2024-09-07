package entity

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Object struct {
	X      float64
	Y      float64
	Width  uint
	Height uint
	Img    *ebiten.Image
	XSpeed float64
	YSpeed float64
}

func NewObject(x float64, y float64, width uint, height uint, color color.RGBA) *Object {
	o := &Object{
		X:      x,
		Y:      y,
		Width:  width,
		Height: height,
		Img:    ebiten.NewImage(int(width), int(height)),
	}
	o.Img.Fill(color)
	return o
}

func NewObjectSpd(o *Object, xSpd float64, ySpd float64) {
	o.XSpeed = xSpd
	o.YSpeed = ySpd
}
