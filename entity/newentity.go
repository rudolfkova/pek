package entity

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/rudolfkova/pek/vec"
)

type Object struct {
	X         float64
	Y         float64
	Width     uint
	Height    uint
	Img       *ebiten.Image
	XSpeed    float64
	YSpeed    float64
	Collision *Crossroad
	XCenter float64
	YCenter float64
	AB vec.Vec
	BC vec.Vec
	CD vec.Vec
	DA vec.Vec
}



type Crossroad int

const (
	Vertical Crossroad = iota
	Horizontal
	None
)

func NewObject(x float64, y float64, width uint, height uint, color color.RGBA) *Object {
	o := &Object{
		X:      x,
		Y:      y,
		Width:  width,
		Height: height,
		Img:    ebiten.NewImage(int(width), int(height)),
	}
	o.XCenter = x + float64(o.Width) / 2
	o.YCenter = y + float64(o.Height) / 2
	o.Collision = new(Crossroad)
	o.Img.Fill(color)
	return o
}

func NewObjectSpd(o *Object, xSpd float64, ySpd float64) {
	o.XSpeed = xSpd
	o.YSpeed = ySpd
}
