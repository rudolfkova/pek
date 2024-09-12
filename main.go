package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/rudolfkova/pek/cords"
	"github.com/rudolfkova/pek/entity"
	"github.com/rudolfkova/pek/physics"
)

const (
	screenWidth  = 1000
	screenHeight = 500
)

type Game struct {
	redCircle *entity.Object
	blueCircle *entity.Object
	pump1     *entity.Object
	pump2     *entity.Object
}

func NewGame() *Game {
	g := &Game{}
	redCircle := entity.NewObject(20, 250, 10, 10, color.RGBA{255, 0, 0, 255})
	blueCircle := entity.NewObject(20, 300, 10, 10, color.RGBA{0, 0, 255, 255})
	entity.NewObjectSpd(redCircle, 5, 5)
	entity.NewObjectSpd(blueCircle, 5, 5)
	pump1 := entity.NewObject(500, 300, 500, 20, color.RGBA{0, 255, 0, 255})
	pump2 := entity.NewObject(0, 200, 500, 20, color.RGBA{0, 255, 0, 255})
	g.blueCircle = blueCircle
	g.pump2 = pump2
	g.redCircle = redCircle
	g.pump1 = pump1
	physics.InitDyn(
		g.redCircle,
		g.blueCircle,
	)
	physics.InitStat(
		g.pump1,
		g.pump2,
	)

	return g
}

func (g *Game) Update() error {
	physics.Collision()
	g.redCircle.X += g.redCircle.XSpeed
	g.redCircle.Y += g.redCircle.YSpeed
	g.blueCircle.X += g.blueCircle.XSpeed
	g.blueCircle.Y += g.blueCircle.YSpeed

	g.redCircle.XSpeed, g.redCircle.YSpeed = physics.ScreenCollision(*g.redCircle, screenWidth, screenHeight)
	g.blueCircle.XSpeed, g.blueCircle.YSpeed = physics.ScreenCollision(*g.blueCircle, screenWidth, screenHeight)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0xff, 0xff, 0xff, 0xff})
	cords.DrawPlane(screen, screenWidth, screenHeight)
	cords.DebugCoordsObject(screen, g.redCircle)
	// Красный круг
	opRedCircle := &ebiten.DrawImageOptions{}
	opRedCircle.GeoM.Translate(g.redCircle.X, g.redCircle.Y)
	screen.DrawImage(g.redCircle.Img, opRedCircle)
	// Труба1
	opPump1 := &ebiten.DrawImageOptions{}
	opPump1.GeoM.Translate(g.pump1.X, g.pump1.Y)
	screen.DrawImage(g.pump1.Img, opPump1)
	cords.MousePos(screen)
	// Труба2
	opPump2 := &ebiten.DrawImageOptions{}
	opPump2.GeoM.Translate(g.pump2.X, g.pump2.Y)
	screen.DrawImage(g.pump2.Img, opPump2)
	// Синий круг
	opBlueCircle := &ebiten.DrawImageOptions{}
	opBlueCircle.GeoM.Translate(g.blueCircle.X, g.blueCircle.Y)
	screen.DrawImage(g.blueCircle.Img, opBlueCircle)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	g := NewGame()
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Синий круг")
	ebiten.SetVsyncEnabled(false)
	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}
