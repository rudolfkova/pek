package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"

	"github.com/rudolfkova/pek/cords"
	"github.com/rudolfkova/pek/entity"
	"github.com/rudolfkova/pek/physics"
)

const (
	screenWidth  = 1000
	screenHeight = 500
)

type Game struct {
	redCircle  *entity.Object
	blueCircle *entity.Object
	pump1      *entity.Object
	pump2      *entity.Object
	character  *entity.Character
}

func NewGame() *Game {
	g := &Game{}
	redCircle := entity.NewObject(20, 250, 10, 10, color.RGBA{255, 0, 0, 255})
	blueCircle := entity.NewObject(20, 300, 10, 10, color.RGBA{0, 0, 255, 255})
	entity.NewObjectSpd(redCircle, 10, 5)
	entity.NewObjectSpd(blueCircle, 15, 2)
	pump1 := entity.NewObject(100, 100, 100, 100, color.RGBA{0, 255, 0, 255})
	pump2 := entity.NewObject(300, 300, 100, 100, color.RGBA{0, 255, 0, 255})
	character := entity.NewCharacter("Артем", 200, 300, 10, 20, color.RGBA{R: 0, G: 0, B: 255, A: 255})
	g.character = character
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
	physics.NewStatVec()

	return g
}

func (g *Game) Update() error {
	physics.Collision()
	physics.Move()
	physics.ScreenCollision(screenWidth, screenHeight)
	physics.CharacterMove(g.character)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 0, 0, 255})
	cords.DebugCoordsObject(screen, g.redCircle)
	cords.DebugCharacter(screen, g.character, g.pump1)
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
	// Персонаж
	opCharacter := &ebiten.DrawImageOptions{}
	opCharacter.GeoM.Translate(g.character.X, g.character.Y)
	screen.DrawImage(g.character.Img, opCharacter)
	//Вектор скорости
	vector.StrokeLine(screen, float32(g.character.SpdVec.X1), float32(g.character.SpdVec.Y1), float32(g.character.SpdVec.X2), float32(g.character.SpdVec.Y2), 1, color.RGBA{255, 255, 255, 255}, true)
	xc, yc, err := g.character.SpdVec.Intersect(&g.pump1.AB)
	if err == nil && xc > g.pump1.X && xc < g.pump1.X+float64(g.pump1.Width) {
		vector.StrokeLine(screen, float32(g.character.SpdVec.X1), float32(g.character.SpdVec.Y1), float32(xc), float32(yc), 1, color.RGBA{255, 255, 255, 255}, true)
	}

	xc2, yc2, err2 := g.character.SpdVec.Intersect(&g.pump1.BC)
	if err2 == nil && yc2 > g.pump1.BC.Y1 && yc2 < g.pump1.BC.Y2 {
		vector.StrokeLine(screen, float32(g.character.SpdVec.X1), float32(g.character.SpdVec.Y1), float32(xc2), float32(yc2), 1, color.RGBA{255, 255, 255, 255}, true)
	}

	// xc1, yc1, err1 := g.character.SpdVec.Intersect(&g.pump1.BC)
	// if err1 == nil && yc1 > g.pump1.Y && yc1 < g.pump1.Y + float64(g.pump1.Height) {
	// 	vector.StrokeLine(screen, float32(g.character.SpdVec.X1), float32(g.character.SpdVec.Y1), float32(xc1), float32(yc1), 1, color.RGBA{255, 255, 255, 255}, true)
	// }

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	g := NewGame()
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Андрей Паньков")
	ebiten.SetVsyncEnabled(true)
	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}
