package physics

import (
	"github.com/rudolfkova/pek/entity"
)

func XCross(dyn entity.Object, stat entity.Object) bool {
	return dyn.X+float64(dyn.Width) >= stat.X && dyn.X <= stat.X+float64(stat.Width)
}

func YCross(dyn entity.Object, stat entity.Object) bool {
	return dyn.Y+float64(dyn.Height) >= stat.Y && dyn.Y <= stat.Y+float64(stat.Height)
}

var dynObj []*entity.Object
var statObj []*entity.Object

func InitDyn(dyn ...*entity.Object) {
	dynObj = append(dynObj, dyn...)
}
func InitStat(stat ...*entity.Object) {
	statObj = append(statObj, stat...)
}

func Collision() {
	for _, d := range dynObj {
		for _, s := range statObj {
			if *d.Collision == 1 && XCross(*d, *s) && YCross(*d, *s) {
				*d.Collision = 0
				d.XSpeed = -d.XSpeed
			}
			if *d.Collision == 2 && XCross(*d, *s) && YCross(*d, *s) {
				*d.Collision = 0
				d.YSpeed = -d.YSpeed
			}
			switch true {
			case YCross(*d, *s):
				*d.Collision = 1
			case XCross(*d, *s):
				*d.Collision = 2
			}
		}
	}
}

func ScreenCollision(dyn entity.Object, screenWidth int, screenHeight int) (float64, float64) {
	if dyn.X+float64(dyn.Width) > float64(screenWidth) || dyn.X < 0 {
		dyn.XSpeed = -dyn.XSpeed
		return dyn.XSpeed, dyn.YSpeed
	}
	if dyn.Y+float64(dyn.Height) > float64(screenHeight) || dyn.Y < 0 {
		dyn.YSpeed = -dyn.YSpeed
		return dyn.XSpeed, dyn.YSpeed
	}
	return dyn.XSpeed, dyn.YSpeed
}
