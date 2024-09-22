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

// Инициализация столкновений динамических объектов со статическими
func DynCollision() {
	for _, c := range dynObj {
		for _, stat := range statObj {
			c.XCenter = c.X + float64(c.Width)/2
			c.YCenter = c.Y + float64(c.Height)/2
			colLine := vec.NewVec(c.XCenter, c.YCenter, stat.XCenter, stat.YCenter)
			xc1, _, err1 := colLine.Intersect(&stat.AB)
			xc2, _, err2 := colLine.Intersect(&stat.DC)
			_, yc3, err3 := colLine.Intersect(&stat.BC)
			_, yc4, err4 := colLine.Intersect(&stat.AD)
			c.SpdVec = *vec.NewVec(c.XCenter, c.YCenter, c.XCenter+c.XSpeed, c.YCenter+c.YSpeed)
			if (c.YCenter+float64(c.Height/2) == stat.AB.Y1) && (xc1 > stat.AB.X1 && xc1 < stat.AB.X2) && ABSign(c.SpdVec) && c.AllCross(statObj) && err1 == nil {
				c.YSpeed = -c.YSpeed
			}

			if (c.YCenter-float64(c.Height/2) == stat.DC.Y1) && (xc2 > stat.DC.X1 && xc2 < stat.DC.X2) && DCSign(c.SpdVec) && c.AllCross(statObj) && err2 == nil {
				c.YSpeed = -c.YSpeed
			}

			if (c.XCenter-float64(c.Width/2) == stat.BC.X1) && (yc3 > stat.BC.Y1 && yc3 < stat.BC.Y2) && BCSign(c.SpdVec) && c.AllCross(statObj) && err3 == nil {
				c.XSpeed = -c.XSpeed
			}

			if (c.XCenter+float64(c.Width/2) == stat.AD.X1) && (yc4 > stat.AD.Y1 && yc4 < stat.AD.Y2) && ADSign(c.SpdVec) && c.AllCross(statObj) && err4 == nil {
				c.XSpeed = -c.XSpeed

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
func DynMove() {
	for _, d := range dynObj {
		d.X += d.XSpeed
		d.Y += d.YSpeed
	}
}

func CharacterMove(c *entity.Character) {
	c.XCenter = c.X + float64(c.Width)/2
	c.YCenter = c.Y + float64(c.Height)/2
	c.SpdVec = *vec.NewVec(c.XCenter, c.YCenter, c.XCenter+c.XSpeed, c.YCenter+c.YSpeed)
	CharacterCollision(c)
	if ebiten.IsKeyPressed(ebiten.KeyW) && disKeyWS != disKeyW {
		if c.YSpeed >= 0 {
			c.YSpeed = -c.YSpeedConst
		}
		c.XCenter = c.X + float64(c.Width)/2
		c.YCenter = c.Y + float64(c.Height)/2
		c.SpdVec = *vec.NewVec(c.XCenter, c.YCenter, c.XCenter+c.XSpeed, c.YCenter+c.YSpeed)
		CharacterCollision(c)
		c.Y += c.YSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) && disKeyWS != disKeyS {
		if c.YSpeed <= 0 {
			c.YSpeed = c.YSpeedConst
		}
		c.XCenter = c.X + float64(c.Width)/2
		c.YCenter = c.Y + float64(c.Height)/2
		c.SpdVec = *vec.NewVec(c.XCenter, c.YCenter, c.XCenter+c.XSpeed, c.YCenter+c.YSpeed)
		CharacterCollision(c)
		c.Y += c.YSpeed

	}
	if ebiten.IsKeyPressed(ebiten.KeyA) && disKeyAD != disKeyA {
		if c.XSpeed >= 0 {
			c.XSpeed = -c.XSpeedConst
		}
		c.XCenter = c.X + float64(c.Width)/2
		c.YCenter = c.Y + float64(c.Height)/2
		c.SpdVec = *vec.NewVec(c.XCenter, c.YCenter, c.XCenter+c.XSpeed, c.YCenter+c.YSpeed)
		CharacterCollision(c)
		c.X += c.XSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) && disKeyAD != disKeyD {
		if c.XSpeed <= 0 {
			c.XSpeed = c.XSpeedConst
		}
		c.XCenter = c.X + float64(c.Width)/2
		c.YCenter = c.Y + float64(c.Height)/2
		c.SpdVec = *vec.NewVec(c.XCenter, c.YCenter, c.XCenter+c.XSpeed, c.YCenter+c.YSpeed)
		CharacterCollision(c)
		c.X += c.XSpeed
	}
	if !ebiten.IsKeyPressed(ebiten.KeyW) && !ebiten.IsKeyPressed(ebiten.KeyS) {
		c.YSpeed = 0
	}
	if !ebiten.IsKeyPressed(ebiten.KeyA) && !ebiten.IsKeyPressed(ebiten.KeyD) {
		c.XSpeed = 0
	}
}

func ABSign(v vec.Vec) bool {
	b := v.Sign()
	return b == vec.Down || b == vec.Positive || b == vec.NegativePositive
}

func BCSign(v vec.Vec) bool {
	b := v.Sign()
	return b == vec.Left || b == vec.Negative || b == vec.NegativePositive
}

func DCSign(v vec.Vec) bool {
	b := v.Sign()
	return b == vec.Up || b == vec.PositiveNegative || b == vec.Negative
}

func ADSign(v vec.Vec) bool {
	b := v.Sign()
	return b == vec.Right || b == vec.PositiveNegative || b == vec.Positive
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

func CharacterCollision(c *entity.Character) {
	for _, stat := range statObj {
		colLine := vec.NewVec(c.XCenter, c.YCenter, stat.XCenter, stat.YCenter)
		xc1, _, err1 := colLine.Intersect(&stat.AB)
		xc2, _, err2 := colLine.Intersect(&stat.DC)
		_, yc3, err3 := colLine.Intersect(&stat.BC)
		_, yc4, err4 := colLine.Intersect(&stat.AD)
		if (c.YCenter+float64(c.Height/2) == stat.AB.Y1) && (xc1 > stat.AB.X1 && xc1 < stat.AB.X2) && ABSign(c.SpdVec) && c.AllCross(statObj) && err1 == nil {
			disKeyWS = disKeyS
		}

		if (c.YCenter-float64(c.Height/2) == stat.DC.Y1) && (xc2 > stat.DC.X1 && xc2 < stat.DC.X2) && DCSign(c.SpdVec) && c.AllCross(statObj) && err2 == nil {
			disKeyWS = disKeyW
		}

		if (c.XCenter-float64(c.Width/2) == stat.BC.X1) && (yc3 > stat.BC.Y1 && yc3 < stat.BC.Y2) && BCSign(c.SpdVec) && c.AllCross(statObj) && err3 == nil {
			disKeyAD = disKeyA
		}

		if (c.XCenter+float64(c.Width/2) == stat.AD.X1) && (yc4 > stat.AD.Y1 && yc4 < stat.AD.Y2) && ADSign(c.SpdVec) && c.AllCross(statObj) && err4 == nil {
			disKeyAD = disKeyD

		}
		if !c.AllCross(statObj) {
			disKeyWS = disKeyNone
			disKeyAD = disKeyNone
		}
	}
}
