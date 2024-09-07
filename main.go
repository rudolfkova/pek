package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/rudolfkova/pek/cords"
	"github.com/rudolfkova/pek/entity"
	"github.com/rudolfkova/pek/physics"
)

const (
	screenWidth  = 640
	screenHeight = 320
)

type Game struct {
	//Вода
	water       *ebiten.Image
	xWater      float64
	yWater      float64
	xSpeedWater float64
	ySpeedWater float64
	//Труба
	pump      *ebiten.Image
	xPump     float64
	yPump     float64
	redCircle *entity.Object
	pump1     *entity.Object
}

func NewGame() *Game {
	g := &Game{
		water:       ebiten.NewImage(10, 10),
		xWater:      1,
		yWater:      1,
		xSpeedWater: 0.3,
		ySpeedWater: 0.3,
		pump:        ebiten.NewImage(10, 100),
		xPump:       100,
		yPump:       50,
	}
	g.water.Fill(color.RGBA{0, 0, 255, 255})
	g.pump.Fill(color.RGBA{255, 255, 255, 255})
	return g
}

func (g *Game) Update() error {

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0xff, 0xff, 0xff, 0xff})
	cords.DrawPlane(screen, screenWidth, screenHeight)
	cords.DebugCoords(screen, g.xWater, g.yWater, g.xSpeedWater, g.ySpeedWater, 0, 0)
	cords.DebugCoordsObject(screen, g.redCircle)
	cords.MousePos(screen)
	// Перемещаем круг
	g.xWater += g.xSpeedWater
	g.yWater += g.ySpeedWater
	g.redCircle.X += g.redCircle.XSpeed
	g.redCircle.Y += g.redCircle.YSpeed
	g.redCircle.XSpeed, g.redCircle.YSpeed = physics.Collision(*g.redCircle, *g.pump1)
	g.redCircle.XSpeed, g.redCircle.YSpeed = physics.ScreenCollision(*g.redCircle, screenWidth, screenHeight)
	if g.xWater+5 > g.xPump-5 && g.xWater-5 < g.xPump+5 && g.yWater+5 > g.yPump-50 && g.yWater-5 < g.yPump+50 {
		g.xSpeedWater = -g.xSpeedWater
	}
	if g.xWater+10 > screenWidth || g.xWater < 0 {
		g.xSpeedWater = -g.xSpeedWater
	}
	if g.yWater+10 > screenHeight || g.yWater < 0 {
		g.ySpeedWater = -g.ySpeedWater
	}
	g.xSpeedWater, g.ySpeedWater = cords.MouseStop(g.xSpeedWater, g.ySpeedWater)
	// Вода
	opWater := &ebiten.DrawImageOptions{}
	opWater.GeoM.Translate(g.xWater, g.yWater)
	screen.DrawImage(g.water, opWater)
	// Труба
	opPump := &ebiten.DrawImageOptions{}
	opPump.GeoM.Translate(g.xPump, g.yPump)
	screen.DrawImage(g.pump, opPump)
	// Красный круг
	opRedCircle := &ebiten.DrawImageOptions{}
	opRedCircle.GeoM.Translate(g.redCircle.X, g.redCircle.Y)
	screen.DrawImage(g.redCircle.Img, opRedCircle)
	// Труба1
	opPump1 := &ebiten.DrawImageOptions{}
	opPump1.GeoM.Translate(g.pump1.X, g.pump1.Y)
	screen.DrawImage(g.pump1.Img, opPump1)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	redCircle := entity.NewObject(100, 100, 10, 10, color.RGBA{255, 0, 0, 255})
	entity.NewObjectSpd(redCircle, 0.2, 0.2)
	pump1 := entity.NewObject(200, 100, 100, 100, color.RGBA{255, 255, 255, 255})
	g := NewGame()
	g.redCircle = redCircle
	g.pump1 = pump1
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Синий круг")
	ebiten.SetVsyncEnabled(false)
	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}
