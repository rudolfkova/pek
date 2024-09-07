package cords

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Указывает координаты курсора при нажатии левой клавиши мыши
func MousePos(screen *ebiten.Image) {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		ebitenutil.DebugPrintAt(
			screen,
			fmt.Sprintf(
				"x: %d, y: %d",
				x,
				y,
			),
			x, y,
		)
	}
}

func MouseStop(xSpd float64, ySpd float64) (float64, float64) {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
		return 0, 0
	}
	return xSpd, ySpd
}
