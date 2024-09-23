package entity

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/rudolfkova/pek/vec"
)

type Object struct {
	X            float64
	Y            float64
	Width        uint
	Height       uint
	Img          *ebiten.Image
	XSpeed       float64
	YSpeed       float64
	SpdVec       Line.Line
	DynCollision *Crossroad
	XCenter      float64
	YCenter      float64
	AB           Line.Line
	BC           Line.Line
	DC           Line.Line
	AD           Line.Line
	Orient       int8
	Color        color.RGBA
}

type IObject interface {
	XCross()
	YCross()
	AnyCrossX()
	AnyCrossY()
	AllCross()
}

func (c Character) XCross(stat *Object) bool {
	return c.X+float64(c.Width) >= stat.X && c.X <= stat.X+float64(stat.Width)
}
func (c Character) YCross(stat *Object) bool {
	return c.Y+float64(c.Height) >= stat.Y && c.Y <= stat.Y+float64(stat.Height)
}

func (c Character) AnyCrossX(stat []*Object) bool {
	var b bool
	for _, s := range stat {
		b = c.XCross(s)
		if b {
			break
		}
	}
	return b
}

func (c Character) AnyCrossY(stat []*Object) bool {
	var b bool
	for _, s := range stat {
		b = c.YCross(s)
		if b {
			break
		}
	}
	return b
}

func (c Character) AllCross(stat []*Object) bool {
	var b bool
	for _, s := range stat {
		b = c.XCross(s) && c.YCross(s)
		if b {
			break
		}
	}
	return b
}

func (d Object) XCross(stat *Object) bool {
	return d.X+float64(d.Width) >= stat.X && d.X <= stat.X+float64(stat.Width)
}

func (d Object) YCross(stat *Object) bool {
	return d.Y+float64(d.Height) >= stat.Y && d.Y <= stat.Y+float64(stat.Height)
}

func (d Object) AnyCrossX(stat []*Object) bool {
	var b bool
	for _, s := range stat {
		b = d.XCross(s)
		if b {
			break
		}
	}
	return b
}

func (d Object) AnyCrossY(stat []*Object) bool {
	var b bool
	for _, s := range stat {
		b = d.YCross(s)
		if b {
			break
		}
	}
	return b
}
func (d Object) AllCross(stat []*Object) bool {
	return d.AnyCrossX(stat) && d.AnyCrossY(stat)
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
	o.DynCollision = new(Crossroad)
	o.Img.Fill(color)
	return o
}

func NewObjectSpd(o *Object, xSpd float64, ySpd float64) {
	o.XSpeed = xSpd
	o.YSpeed = ySpd
}

func (o *Object) Split() ([]*Object, error) {

	var splitedObj []*Object
	var l int
	if o.Width >= o.Height {
		if o.Width%o.Height != 0 {
			return nil, fmt.Errorf("Can't split")
		}
		l = int(o.Width) / int(o.Height)
		for i := 0; i < l; i++ {
			a := NewObject(o.X+float64(int(o.Height)*i), o.Y, o.Height, o.Height, o.Color)
			splitedObj = append(splitedObj, a)
		}
	} else {
		if o.Height%o.Width != 0 {
			return nil, fmt.Errorf("Can't split")
		}
		l = int(o.Height) / int(o.Width)
		for i := 0; i < l; i++ {
			a := NewObject(o.X, o.Y+float64(int(o.Width)*i), o.Width, o.Width, o.Color)
			splitedObj = append(splitedObj, a)
		}
	}
	return splitedObj, nil
}
