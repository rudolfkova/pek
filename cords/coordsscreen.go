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
