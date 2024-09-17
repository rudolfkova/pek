package entity

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/rudolfkova/pek/vec"
)

type Character struct {
	Name string
	X float64
	Y float64
	XSpeed float64
	XSpeedConst float64
	YSpeed float64
	YSpeedConst float64
	Img *ebiten.Image
	Width uint
	Height uint
	HP int
	Collision *Crossroad
	XCenter float64
	YCenter float64
	SpdVec vec.Vec
}

func NewCharacter(name string, x float64, y float64, width uint, height uint, color color.RGBA) *Character {
	c:= &Character{
		Name: name,
		X: x,
		Y: y,
		Width: width,
		Height: height,
		HP: 100,
		Img: ebiten.NewImage(int(width), int(height)),
		XSpeedConst: 2,
		YSpeedConst: 2,
	}
	c.XCenter = x + float64(width) / 2
	c.YCenter = y + float64(height) / 2
	c.Collision = new(Crossroad)
	c.Img.Fill(color)
	return c
}




