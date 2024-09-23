package Line

import (
	"fmt"
	"image/color"
	"math"
)

type Line struct {
	X1, Y1 float64
	X2, Y2 float64
	Color  color.RGBA
}

func NewVec(x1, y1, x2, y2 float64) *Line {
	return &Line{
		X1: x1,
		Y1: y1,
		X2: x2,
		Y2: y2,
	}
}
func (v *Line) Length() float64 {
	return math.Sqrt(math.Pow(v.X2-v.X1, 2) + math.Pow(v.Y2-v.Y1, 2))
}

func (v *Line) XLength() float64 {
	return math.Abs(v.X2 - v.X1)
}
func (v *Line) YLength() float64 {
	return math.Abs(v.Y2 - v.Y1)
}

// Метод, который определяет, пересекаются ли два отрезка
func (v *Line) Intersect(v2 *Line) (float64, float64, error) {
	var x float64
	var y float64
	K1 := (v.Y1 - v.Y2) / (v.X1 - v.X2)
	K2 := (v2.Y1 - v2.Y2) / (v2.X1 - v2.X2)
	if K1 == math.Inf(-1) && K2 == math.Inf(-1) {
		return 0, 0, fmt.Errorf("Отрезки не пересекаются")
	}
	if K1 == math.Inf(1) && K2 == math.Inf(1) {
		return 0, 0, fmt.Errorf("Отрезки не пересекаются")
	}
	if K2 == 0 {
		x = -(v.X2*v.Y1 - v.Y2*v.X1 - v.X2*v2.Y1 + v.X1*v2.Y1) / (v.Y2 - v.Y1)
		y = v2.Y1
	}
	if K2 == math.Inf(-1) || K2 == math.Inf(1) {
		y = (v.X2*v.Y1 - v.Y2*v.X1 + v.Y2*v2.X1 - v2.X1*v.Y1) / (v.X2 - v.X1)
		x = v2.X1
	}
	return x, y, nil
}

func (v *Line) Dist(x, y float64, err error) (float64, error) {
	if err != nil {
		return 0, err
	}
	return math.Sqrt(math.Pow(v.X1-x, 2) + math.Pow(v.Y1-y, 2)), nil
}

func (v *Line) Signs(xc, yc float64) bool {
	var positive bool = ((xc-v.X1 > 0) && (yc-v.Y1 > 0)) && ((v.X2-v.X1 > 0) && (v.Y2-v.Y1 > 0))
	var negative bool = ((xc-v.X1 < 0) && (yc-v.Y1 < 0)) && ((v.X2-v.X1 < 0) && (v.Y2-v.Y1 < 0))
	var positiveNegative bool = ((xc-v.X1 > 0) && (yc-v.Y1 < 0)) && ((v.X2-v.X1 > 0) && (v.Y2-v.Y1 < 0))
	var negativePositive bool = ((xc-v.X1 < 0) && (yc-v.Y1 > 0)) && ((v.X2-v.X1 < 0) && (v.Y2-v.Y1 > 0))
	var right bool = ((xc-v.X1 > 0) && (yc-v.Y1 == 0)) && ((v.X2-v.X1 > 0) && (v.Y2-v.Y1 == 0))
	var left bool = ((xc-v.X1 < 0) && (yc-v.Y1 == 0)) && ((v.X2-v.X1 < 0) && (v.Y2-v.Y1 == 0))
	var down bool = ((xc-v.X1 == 0) && (yc-v.Y1 > 0)) && ((v.X2-v.X1 == 0) && (v.Y2-v.Y1 > 0))
	var up bool = ((xc-v.X1 == 0) && (yc-v.Y1 < 0)) && ((v.X2-v.X1 == 0) && (v.Y2-v.Y1 < 0))
	return positive || negative || positiveNegative || negativePositive || right || left || down || up
}

type Sign int

const (
	Positive Sign = iota
	Negative
	PositiveNegative
	NegativePositive
	Zero
	Right
	Left
	Down
	Up
)

func (v *Line) Sign() Sign {
	//Право-низ
	if (v.X2-v.X1 > 0) && (v.Y2-v.Y1 > 0) {
		return Positive
	}
	//Лево-верх
	if (v.X2-v.X1 < 0) && (v.Y2-v.Y1 < 0) {
		return Negative
	}
	//Право-верх
	if (v.X2-v.X1 > 0) && (v.Y2-v.Y1 < 0) {
		return PositiveNegative
	}
	//Лево-низ
	if (v.X2-v.X1 < 0) && (v.Y2-v.Y1 > 0) {
		return NegativePositive
	}
	//Право
	if (v.X2-v.X1 > 0) && (v.Y2-v.Y1 == 0) {
		return Right
	}
	//Лево
	if (v.X2-v.X1 < 0) && (v.Y2-v.Y1 == 0) {
		return Left
	}
	//Низ
	if (v.X2-v.X1 == 0) && (v.Y2-v.Y1 > 0) {
		return Down
	}
	//Верх
	if (v.X2-v.X1 == 0) && (v.Y2-v.Y1 < 0) {
		return Up
	}
	return Zero
}
