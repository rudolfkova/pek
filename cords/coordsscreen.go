package cords

import (
	"fmt"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/rudolfkova/pek/entity"
	"github.com/rudolfkova/pek/physics"
)

// Создаёт координатную плоскость на экране
func DrawPlane(screen *ebiten.Image, screenWidth int, screenHeight int) {
	for i := 0; i < screenWidth; i += 20 {
		for j := 0; j < screenHeight; j += 20 {
			vector.DrawFilledRect(screen, float32(i), float32(j), 19, 19, color.RGBA{0x00, 0x00, 0x00, 0xff}, false)
		}
	}
}

// Указывает координаты и скорость цели
func DebugCoords(screen *ebiten.Image, x float64, y float64, xSpd float64, ySpd float64, xPos int, yPos int) {
	spd := math.Sqrt(math.Pow(xSpd, 2) + math.Pow(ySpd, 2))
	ebitenutil.DebugPrintAt(
		screen,
		fmt.Sprintf(
			"x: %.2f, y: %.2f, xSpd: %.2f, ySpd: %.2f, spd: %.2f",
			x,
			y,
			xSpd,
			ySpd,
			spd,
		),
		xPos, yPos,
	)
}

func DebugCoordsObject(screen *ebiten.Image, obj *entity.Object) {
	spd := math.Sqrt(math.Pow(obj.XSpeed, 2) + math.Pow(obj.YSpeed, 2))
	xPos, yPos := 0, 20
	ebitenutil.DebugPrintAt(
		screen,
		fmt.Sprintf(
			"x: %.2f, y: %.2f, xSpd: %.2f, ySpd: %.2f, spd: %.2f",
			obj.X,
			obj.Y,
			obj.XSpeed,
			obj.YSpeed,
			spd,
		),
		xPos, yPos,
	)
}

func DebugCharacter(screen *ebiten.Image, obj *entity.Character, stat *entity.Object) {
	xPos, yPos := 0, 0
	ebitenutil.DebugPrintAt(
		screen,
		fmt.Sprintf(
			"XCrossCharacter: %t, YCrossCharacter:%t, XSpeed:%.2f, YSpeed:%.2f",
			physics.XCrossCharacter(*obj, *stat),
			physics.YCrossCharacter(*obj, *stat),
			obj.XSpeed,
			obj.YSpeed,
		),
		xPos, yPos,
	)
}

func DebugCollision(screen *ebiten.Image, c *entity.Character, stat ...*entity.Object) {
	vector.StrokeLine(screen, float32(c.SpdVec.X1), float32(c.SpdVec.Y1), float32(c.SpdVec.X2), float32(c.SpdVec.Y2), 1, color.RGBA{255, 255, 255, 255}, true)
	for _, stat := range stat {
		xc1, yc1, err1 := c.SpdVec.Intersect(&stat.AB)
		xc2, yc2, err2 := c.SpdVec.Intersect(&stat.DC)
		xc3, yc3, err3 := c.SpdVec.Intersect(&stat.BC)
		xc4, yc4, err4 := c.SpdVec.Intersect(&stat.AD)
		var len1 float64
		var len2 float64
		var len3 float64
		var len4 float64
		var minLen []float64

		if err1 == nil && xc1 > stat.AB.X1 && xc1 < stat.AB.X2 {
			len1 = math.Sqrt(math.Pow(xc1-c.SpdVec.X1, 2) + math.Pow(yc1-c.SpdVec.Y1, 2))
			minLen = append(minLen, len1)
		}

		if err2 == nil && xc2 > stat.DC.X1 && xc2 < stat.DC.X2 {
			len2 = math.Sqrt(math.Pow(xc2-c.SpdVec.X1, 2) + math.Pow(yc2-c.SpdVec.Y1, 2))
			minLen = append(minLen, len2)
		}

		if err3 == nil && yc3 > stat.BC.Y1 && yc3 < stat.BC.Y2 {
			len3 = math.Sqrt(math.Pow(xc3-c.SpdVec.X1, 2) + math.Pow(yc3-c.SpdVec.Y1, 2))
			minLen = append(minLen, len3)
		}

		if err4 == nil && yc4 > stat.AD.Y1 && yc4 < stat.AD.Y2 {
			len4 = math.Sqrt(math.Pow(xc4-c.SpdVec.X1, 2) + math.Pow(yc4-c.SpdVec.Y1, 2))
			minLen = append(minLen, len4)
		}

		if len(minLen) > 0 {
			minLenValue := minLen[0]
			for _, value := range minLen {
				minLenValue = math.Min(minLenValue, value)
			}
			if minLenValue == len1 {
				// if c.SpdVec.Signs(xc1, yc1) {
				vector.StrokeLine(screen, float32(c.SpdVec.X1), float32(c.SpdVec.Y1), float32(xc1), float32(yc1), 1, color.RGBA{255, 255, 255, 255}, true)
				// }
			}
			if minLenValue == len2 {
				// if c.SpdVec.Signs(xc2, yc2) {
				vector.StrokeLine(screen, float32(c.SpdVec.X1), float32(c.SpdVec.Y1), float32(xc2), float32(yc2), 1, color.RGBA{255, 255, 255, 255}, true)
				// }
			}
			if minLenValue == len3 {
				// if c.SpdVec.Signs(xc3, yc3) {
				vector.StrokeLine(screen, float32(c.SpdVec.X1), float32(c.SpdVec.Y1), float32(xc3), float32(yc3), 1, color.RGBA{255, 255, 255, 255}, true)
				// }
			}
			if minLenValue == len4 {
				// if c.SpdVec.Signs(xc4, yc4) {
				vector.StrokeLine(screen, float32(c.SpdVec.X1), float32(c.SpdVec.Y1), float32(xc4), float32(yc4), 1, color.RGBA{255, 255, 255, 255}, true)
				// }
			}
		}

	}
}
