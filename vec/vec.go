package vec

import (
	"fmt"
	"math"
)

type Vec struct {
	X1, Y1 float64
	X2, Y2 float64
}

func NewVec(x1, y1, x2, y2 float64) *Vec {
	return &Vec{
		X1: x1,
		Y1: y1,
		X2: x2,
		Y2: y2,
	}
}
func (v *Vec) Length() float64 {
	return math.Sqrt(math.Pow(v.X2-v.X1, 2) + math.Pow(v.Y2-v.Y1, 2))
}

func (v *Vec) XLength() float64 {
	return math.Abs(v.X2 - v.X1)
}
func (v *Vec) YLength() float64 {
	return math.Abs(v.Y2 - v.Y1)
}

// Метод, который определяет, пересекаются ли два отрезка
func (v *Vec) Intersect(v2 *Vec) (float64, float64, error) {
	var x float64
	var y float64
	K1 := (v.Y1 - v.Y2) / (v.X1 - v.X2)
	K2 := (v2.Y1 - v2.Y2) / (v2.X1 - v2.X2)
	if math.Abs(K1-K2) < 0.1 && math.Abs(K1-K2) > -0.1 {
		return 0, 0, fmt.Errorf("Отрезки не пересекаются")
	}
	if K1 == math.Inf(-1) && K2 == math.Inf(-1) {
		return 0, 0, fmt.Errorf("Отрезки не пересекаются")
	}
	if K1 == math.Inf(1) && K2 == math.Inf(1) {
		return 0, 0, fmt.Errorf("Отрезки не пересекаются")
	}
	if K2 == 0{
		x = -(v.X2 * v.Y1 - v.Y2 * v.X1 - v.X2 * v2.Y1 + v.X1 * v2.Y1) / (v.Y2 - v.Y1)
		y = v2.Y1
	}
	if K2 == math.Inf(-1) || K2 == math.Inf(1) {
		y = (v.X2 * v.Y1 - v.Y2 * v.X1 + v.Y2 * v2.X1 - v2.X1 * v.Y1) / (v.X2 - v.X1)
		x = v2.X1
	}
	// x = (((v.X2 * v.Y1 - v.Y2 * v.X1)/(v.X2 - v.X1)) - ((v2.X2 * v2.Y1 - v2.Y2 * v2.X1)/(v2.X2 - v2.X1))) / (((v.Y2-v.Y1)/(v.X2 - v.X1)) - ((v2.Y2 - v2.Y1)/(v2.X2 - v2.X1)))
	// if K1 < 0 || K2 < 0 {
	// 	x = -x
	// }
	// y = (v.Y2 * x + v.X2 * v.Y1 - v.Y2 * v.X1 - x*v.Y1) / (v.X2 - v.X1)
	//if K1 > 0 && K2 > 0 {
	//	x = (v.X2*v2.X1 + v.X2*v.Y1 - v.Y2*v.X1 - v.X2*v2.Y1 - v.X1*v2.X1 + v.X1*v2.Y1) / (v.X2 - v.Y2 - v.X1 + v.Y1)
	//	y = ((x-v.X1)*(v.Y2-v.Y1) + v.Y1*(v.X2-v.X1)) / (v.X2 - v.X1)
	//}
	//if K1 < 0 && K2 < 0 {
	//	x = (v.X2*v.Y1 - v.X2*v2.X1 + v.Y2*v.X1 - v.X2*v2.Y1 + v.X1*v2.X1 - 2*v.X1*v.Y1 + v.X1*v2.Y1) / (v.X2 - v.Y2 - v.X1 + v.Y1)
	//	y = ((v.X1-x)*(v.Y2-v.Y1) + v.Y1*(v.X2-v.X1)) / (v.X2 - v.X1)
	//}
	//if K1 < 0 && K2 > 0 {
	//	x = (v.X2*v2.X1 + v.X2*v.Y1 + v.Y2*v.X1 - v.X2*v2.Y1 - v.X1*v2.X1 - 2*v.X1*v.Y1 + v.X1*v2.Y1) / (v.X2 + v.Y2 - v.X1 - v.Y1)
	//	y = ((v.X1-x)*(v.Y2-v.Y1) + v.Y1*(v.X2-v.X1)) / (v.X2 - v.X1)
	//}
	//if K1 > 0 && K2 < 0 {
	//	x = (v.X2*v2.X1 - v.X2*v.Y1 + v.Y2*v.X1 + v.X2*v2.Y1 - v.X1*v2.X1 - v.X1*v2.Y1) / (v.X2 + v.Y2 - v.X1 - v.Y1)
	//	y = ((x-v.X1)*(v.Y2-v.Y1) + v.Y1*(v.X2-v.X1)) / (v.X2 - v.X1)
	//}
	return x, y, nil
}

func (v *Vec) Dist (x, y float64, err error) (float64, error) {
	if err!= nil {
		return 0, err
	}
	return math.Sqrt(math.Pow(v.X1 - x, 2) + math.Pow(v.Y1 - y, 2)), nil
}