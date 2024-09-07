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
func Collision(dyn entity.Object, stat entity.Object) (float64, float64) {
	if XCross(dyn, stat) && YCross(dyn, stat) {
		if (dyn.X - stat.X) > (dyn.Y - stat.Y) {
			return dyn.XSpeed, -dyn.YSpeed
		}
		if (dyn.X - stat.X) < (dyn.Y - stat.Y) {
			return -dyn.XSpeed, dyn.YSpeed
		}
		if (dyn.X - (stat.X + float64(stat.Width))) < ((stat.Y + float64(stat.Height)) - dyn.Y) {
			return -dyn.XSpeed, dyn.YSpeed
		}
		if (dyn.X - (stat.X + float64(stat.Width))) > ((stat.Y + float64(stat.Height)) - dyn.Y) {
			return dyn.XSpeed, -dyn.YSpeed
		}
	}
	return dyn.XSpeed, dyn.YSpeed
	// return dyn.X+float64(dyn.Width)/2 > stat.X-float64(stat.Width)/2 && dyn.X-float64(dyn.Width)/2 < stat.X+float64(stat.Width)/2 && dyn.Y+float64(dyn.Height)/2 > stat.Y-float64(stat.Height)/2 && dyn.Y-float64(dyn.Height)/2 < stat.Y+float64(stat.Height)/2
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
