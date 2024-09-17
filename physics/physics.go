package physics

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/rudolfkova/pek/entity"
	"github.com/rudolfkova/pek/vec"
)

/*Физика отражений. Динамические объекты сталкиваются со статическими и отражаются.
*Угол падения равен углу отражения. Отражение происходит с сохранением модуля вектора скорости.
*
*
 */

// Проверка пересечения динамического объекста со статическим по оси X
func XCross(dyn entity.Object, stat entity.Object) bool {
	return dyn.X+float64(dyn.Width) > stat.X && dyn.X < stat.X+float64(stat.Width)
}

// Проверка пересечения динамического объекта со статическим по оси Y
func YCross(dyn entity.Object, stat entity.Object) bool {
	return dyn.Y+float64(dyn.Height) > stat.Y && dyn.Y < stat.Y+float64(stat.Height)
}

func XCrossCharacter(dyn entity.Character, stat entity.Object) bool {
	return dyn.X+float64(dyn.Width) >= stat.X && dyn.X <= stat.X+float64(stat.Width)
}

// Проверка пересечения динамического объекта со статическим по оси Y
func YCrossCharacter(dyn entity.Character, stat entity.Object) bool {
	return dyn.Y+float64(dyn.Height) >= stat.Y && dyn.Y <= stat.Y+float64(stat.Height)
}

// Слайсы для хранения динамических и статических объектов
var dynObj []*entity.Object
var statObj []*entity.Object

// Слайсы для хранения векторов границ статических объектов
var statVec []*vec.Vec
// Для рисовки векторов границ статических объектов
var StatVectors []*vec.Vec = statVec

// Инициализация динамических объектов
func InitDyn(dyn ...*entity.Object) {
	dynObj = append(dynObj, dyn...)
}

// Инициализация статических объектов
func InitStat(stat ...*entity.Object) {
	statObj = append(statObj, stat...)
}

// Инициализация векторов границ статических объектов
func NewStatVec() {
	for _, s := range statObj {
		//AB
		X1_AB := s.X
		Y1_AB := s.Y
		X2_AB := s.X + float64(s.Width)
		Y2_AB := s.Y
		//BC
		X1_BC := s.X + float64(s.Width)
		Y1_BC := s.Y
		X2_BC := s.X + float64(s.Width)
		Y2_BC := s.Y + float64(s.Height)
		//DC
		X1_DC := s.X
		Y1_DC := s.Y + float64(s.Height)
		X2_DC := s.X + float64(s.Width)
		Y2_DC := s.Y + float64(s.Height)
		//AD
		X1_AD := s.X
		Y1_AD := s.Y
		X2_AD := s.X
		Y2_AD := s.Y + float64(s.Height)
		s.AB = *vec.NewVec(X1_AB, Y1_AB, X2_AB, Y2_AB)
		s.BC = *vec.NewVec(X1_BC, Y1_BC, X2_BC, Y2_BC)
		s.DC = *vec.NewVec(X1_DC, Y1_DC, X2_DC, Y2_DC)
		s.AD = *vec.NewVec(X1_AD, Y1_AD, X2_AD, Y2_AD)
		statVec = append(statVec, &s.AB)
		statVec = append(statVec, &s.BC)
		statVec = append(statVec, &s.DC)
		statVec = append(statVec, &s.AD)
	}
}

// Инициализация столкновений динамических объектов со статическими
func Collision() {
	for _, d := range dynObj {
		for _, s := range statObj {
			if *d.Collision == entity.Vertical && XCross(*d, *s) && YCross(*d, *s) {
				*d.Collision = entity.None
				d.XSpeed = -d.XSpeed
			}
			if *d.Collision == entity.Horizontal && XCross(*d, *s) && YCross(*d, *s) {
				*d.Collision = entity.None
				d.YSpeed = -d.YSpeed
			}
			switch true {
			case YCross(*d, *s):
				*d.Collision = entity.Vertical
			case XCross(*d, *s):
				*d.Collision = entity.Horizontal
			}
		}
	}
}

// Столкновение динамического объекта с краями экрана
func ScreenCollision(screenWidth int, screenHeight int) {
	for _, d := range dynObj {
		if d.X+float64(d.Width) > float64(screenWidth) || d.X < 0 {
			d.XSpeed = -d.XSpeed
		}
		if d.Y+float64(d.Height) > float64(screenHeight) || d.Y < 0 {
			d.YSpeed = -d.YSpeed
		}
	}
}

// Функция для движения динамических объектов
func Move() {
	for _, d := range dynObj {
		d.X += d.XSpeed
		d.Y += d.YSpeed
	}
}

func AnyCross(c *entity.Character) bool {
	var b bool
	for _, s := range statObj {
		b = XCrossCharacter(*c, *s) && YCrossCharacter(*c, *s)
		if b {
			break
		}
	}
	return b
}
func AnyCrossX(c *entity.Character) bool {
	var b bool
	for _, s := range statObj {
		b = XCrossCharacter(*c, *s)
		if b {
			break
		}
	}
	return b
}

func AnyCrossY(c *entity.Character) bool {
	var b bool
	for _, s := range statObj {
		b = YCrossCharacter(*c, *s)
		if b {
			break
		}
	}
	return b
}

func AnyCrossAB(c *entity.Character) bool {
	var b bool
	for _, s := range statObj {
		b = s.Orient == 1
		if b {
			break
		}
	}
	return b
}

func AnyCrossBC(c *entity.Character) bool {
	var b bool
	for _, s := range statObj {
		b = s.Orient == 3
		if b {
			break
		}
	}
	return b

}

func AnyCrossDC(c *entity.Character) bool {
	var b bool
	for _, s := range statObj {
		b = s.Orient == 2
		if b {
			break
		}
	}
	return b

}

func AnyCrossAD(c *entity.Character) bool {
	var b bool
	for _, s := range statObj {
		b = s.Orient == 4
		if b {
			break
		}
	}
	return b
}

func CharacterMove(c *entity.Character) {
	c.XCenter = c.X + float64(c.Width)/2
	c.YCenter = c.Y + float64(c.Height)/2
	c.SpdVec = *vec.NewVec(c.XCenter, c.YCenter, c.XCenter+c.XSpeed, c.YCenter+c.YSpeed)
	VecCollision(c)
	if ebiten.IsKeyPressed(ebiten.KeyW) && disKeyWS != disKeyW {
		if c.YSpeed >= 0 {
			c.YSpeed = -c.YSpeedConst
		}
		c.XCenter = c.X + float64(c.Width)/2
		c.YCenter = c.Y + float64(c.Height)/2
		c.SpdVec = *vec.NewVec(c.XCenter, c.YCenter, c.XCenter+c.XSpeed, c.YCenter+c.YSpeed)
		VecCollision(c)
			c.Y += c.YSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) && disKeyWS != disKeyS {
		if c.YSpeed <= 0 {
			c.YSpeed = c.YSpeedConst
		}
		c.XCenter = c.X + float64(c.Width)/2
		c.YCenter = c.Y + float64(c.Height)/2
		c.SpdVec = *vec.NewVec(c.XCenter, c.YCenter, c.XCenter+c.XSpeed, c.YCenter+c.YSpeed)
		VecCollision(c)
		if (!AnyCrossAB(c) || !AnyCross(c)) && (c.SpdVec.Sign() != vec.Down || c.SpdVec.Sign() != vec.Positive || c.SpdVec.Sign() != vec.NegativePositive) {
			c.Y += c.YSpeed
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) && disKeyAD != disKeyA {
		if c.XSpeed >= 0 {
			c.XSpeed = -c.XSpeedConst
		}
		c.XCenter = c.X + float64(c.Width)/2
		c.YCenter = c.Y + float64(c.Height)/2
		c.SpdVec = *vec.NewVec(c.XCenter, c.YCenter, c.XCenter+c.XSpeed, c.YCenter+c.YSpeed)
		VecCollision(c)
			c.X += c.XSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) && disKeyAD != disKeyD {
		if c.XSpeed <= 0 {
			c.XSpeed = c.XSpeedConst
		}
		c.XCenter = c.X + float64(c.Width)/2
		c.YCenter = c.Y + float64(c.Height)/2
		c.SpdVec = *vec.NewVec(c.XCenter, c.YCenter, c.XCenter+c.XSpeed, c.YCenter+c.YSpeed)
		VecCollision(c)
			c.X += c.XSpeed
	}
	if !ebiten.IsKeyPressed(ebiten.KeyW) && !ebiten.IsKeyPressed(ebiten.KeyS) {
		c.YSpeed = 0
	}
	if !ebiten.IsKeyPressed(ebiten.KeyA) && !ebiten.IsKeyPressed(ebiten.KeyD) {
		c.XSpeed = 0
	}
}

func ABSign(c *entity.Character) bool {
	return c.SpdVec.Sign() == vec.Down || c.SpdVec.Sign() == vec.Positive || c.SpdVec.Sign() == vec.NegativePositive
}

func BCSign(c *entity.Character) bool {
	return c.SpdVec.Sign() == vec.Left || c.SpdVec.Sign() == vec.Negative || c.SpdVec.Sign() == vec.NegativePositive
}

func DCSign(c *entity.Character) bool {
	return c.SpdVec.Sign() == vec.Up || c.SpdVec.Sign() == vec.PositiveNegative || c.SpdVec.Sign() == vec.Negative
}

func ADSign(c *entity.Character) bool {
	return c.SpdVec.Sign() == vec.Right || c.SpdVec.Sign() == vec.PositiveNegative || c.SpdVec.Sign() == vec.Positive
}

type disKey int8

const (
	disKeyNone disKey = iota
	disKeyW
	disKeyA
	disKeyS
	disKeyD
)

var disKeyWS disKey
var disKeyAD disKey

func VecCollision(c *entity.Character) {
	for _, stat := range statObj {
		colLine := vec.NewVec(c.XCenter, c.YCenter, stat.XCenter, stat.YCenter)
		xc1, _, err1 := colLine.Intersect(&stat.AB)
		xc2, _, err2 := colLine.Intersect(&stat.DC)
		_, yc3, err3 := colLine.Intersect(&stat.BC)
		_, yc4, err4 := colLine.Intersect(&stat.AD)
		// var len1 float64
		// var len2 float64
		// var len3 float64
		// var len4 float64
		// var minLen []float64
		if (c.YCenter + float64(c.Height/2) == stat.AB.Y1) && (xc1 > stat.AB.X1 && xc1 < stat.AB.X2) && ABSign(c) && AnyCross(c) && err1 == nil {
			disKeyWS = disKeyS
			// len1 = math.Sqrt(math.Pow(xc1-c.SpdVec.X1, 2) + math.Pow(yc1-c.SpdVec.Y1, 2))
			// minLen = append(minLen, len1)
		}

		if (c.YCenter - float64(c.Height/2) == stat.DC.Y1) && (xc2 > stat.DC.X1 && xc2 < stat.DC.X2) && DCSign(c) && AnyCross(c) && err2 == nil {
			disKeyWS = disKeyW
			// len2 = math.Sqrt(math.Pow(xc2-c.SpdVec.X1, 2) + math.Pow(yc2-c.SpdVec.Y1, 2))
			// minLen = append(minLen, len2)
		}

		if (c.XCenter - float64(c.Width/2) == stat.BC.X1) && (yc3 > stat.BC.Y1 && yc3 < stat.BC.Y2) && BCSign(c) && AnyCross(c) && err3 == nil {
			disKeyAD = disKeyA
			// len3 = math.Sqrt(math.Pow(xc3-c.SpdVec.X1, 2) + math.Pow(yc3-c.SpdVec.Y1, 2))
			// minLen = append(minLen, len3)
		}

		if (c.XCenter + float64(c.Width/2) == stat.AD.X1) &&(yc4 > stat.AD.Y1 && yc4 < stat.AD.Y2) && ADSign(c) && AnyCross(c) && err4 == nil {
			disKeyAD = disKeyD
			// len4 = math.Sqrt(math.Pow(xc4-c.SpdVec.X1, 2) + math.Pow(yc4-c.SpdVec.Y1, 2))
			// minLen = append(minLen, len4)
		}
		if !AnyCross(c) {
			disKeyWS = disKeyNone
			disKeyAD = disKeyNone
		}

		// if len(minLen) > 0 {
		// 	minLenValue := minLen[0]
		// 	for _, value := range minLen {
		// 		minLenValue = math.Min(minLenValue, value)
		// 	}
		// 	if minLenValue == len1 {
		// 		if c.SpdVec.Signs(xc1, yc1) {
		// 			stat.Orient = 1
		// 		}
		// 	}
		// 	if minLenValue == len2 {
		// 		if c.SpdVec.Signs(xc2, yc2) {
		// 			stat.Orient = 2
		// 		}
		// 	}
		// 	if minLenValue == len3 {
		// 		if c.SpdVec.Signs(xc3, yc3) {
		// 			stat.Orient = 3
		// 		}
		// 	}
		// 	if minLenValue == len4 {
		// 		if c.SpdVec.Signs(xc4, yc4) {
		// 			stat.Orient = 4
		// 		}
		// 	}
		// }
		// if !AnyCross(c) {
		// 	stat.Orient = 0
		// }
	}
}
