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
	XCenter   float64
	YCenter   float64
	AB        vec.Vec
	BC        vec.Vec
	DC        vec.Vec
	AD        vec.Vec
	Orient    int8
	Color     color.RGBA
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
	o.Color = color
	o.XCenter = x + float64(o.Width)/2
	o.YCenter = y + float64(o.Height)/2
	o.Collision = new(Crossroad)
	o.Img.Fill(color)
	return o
}

func NewObjectSpd(o *Object, xSpd float64, ySpd float64) {
	o.XSpeed = xSpd
	o.YSpeed = ySpd
}

func (o *Object) Split() []*Object{
	var splitedObj []*Object
	var l int
	var n int
	if o.Width >= o.Height {
		l = int(o.Width) / int(o.Height)
		n = int(o.Width) / l
		for i := 0; i < n; i++ {
			a := NewObject(o.X+float64(l*i), o.Y, uint(l), uint(l), o.Color)
			splitedObj = append(splitedObj, a)
		}
	} else {
		l = int(o.Height) / int(o.Width)
		n = int(o.Height) / l
		for i := 0; i < n; i++ {
			a := NewObject(o.X, o.Y+float64(l*i), uint(l), uint(l), o.Color)
			splitedObj = append(splitedObj, a)
		}
	}
	return splitedObj
}
