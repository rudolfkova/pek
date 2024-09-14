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
		s.CD = *vec.NewVec(X1_DC, Y1_DC, X2_DC, Y2_DC)
		s.DA = *vec.NewVec(X1_AD, Y1_AD, X2_AD, Y2_AD)
		statVec = append(statVec, &s.AB)
		statVec = append(statVec, &s.BC)
		statVec = append(statVec, &s.CD)
		statVec = append(statVec, &s.DA)
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

func CharacterMove(c *entity.Character) {
	//for _, s := range statObj {
	c.XCenter = c.X + float64(c.Width)/2
	c.YCenter = c.Y + float64(c.Height)/2
	c.SpdVec = *vec.NewVec(c.XCenter, c.YCenter, c.XCenter+c.XSpeed, c.YCenter+c.YSpeed)
	//if XCrossCharacter(*c, *s) && YCrossCharacter(*c, *s) {
	//	switch true {
	//	case c.XCenter-s.X < s.X+float64(s.Width)-c.XCenter:
	//		c.X = c.X - 1
	//	case c.XCenter-s.X > s.X+float64(s.Width)-c.XCenter:
	//		c.X = c.X + 1
	//	case c.YCenter-s.Y < s.Y+float64(s.Height)-c.YCenter:
	//		c.Y = c.Y - 1
	//	case c.YCenter-s.Y > s.Y+float64(s.Height)-c.YCenter:
	//		c.Y = c.Y + 1
	//	}
	//}
	//}
	//if !AnyCross(c) {
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		if c.YSpeed >= 0 {
			c.YSpeed = -c.YSpeedConst
		}
		c.Y += c.YSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		if c.YSpeed <= 0 {
			c.YSpeed = c.YSpeedConst
		}
		c.Y += c.YSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		if c.XSpeed >= 0 {
			c.XSpeed = -c.XSpeedConst
		}
		c.X += c.XSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		if c.XSpeed <= 0 {
			c.XSpeed = c.XSpeedConst
		}
		c.X += c.XSpeed
	}
	if !ebiten.IsKeyPressed(ebiten.KeyW) && !ebiten.IsKeyPressed(ebiten.KeyS) {
		c.YSpeed = 0
	}
	if !ebiten.IsKeyPressed(ebiten.KeyA) && !ebiten.IsKeyPressed(ebiten.KeyD) {
		c.XSpeed = 0
	}
	//}
}

func VecCollision(c *entity.Character) {
	// for _, s := range statObj {
	// 	x1, y1, err1 := c.SpdVec.Intersect(&s.AB)
	// 	x2, y2, err2 := c.SpdVec.Intersect(&s.BC)
	// 	x3, y3, err3 := c.SpdVec.Intersect(&s.CD)
	// 	x4, y4, err4 := c.SpdVec.Intersect(&s.DA)
	// 	if err1 == nil && err2 == nil && err3 == nil && err4 == nil {

	// 	}
	// }
}
