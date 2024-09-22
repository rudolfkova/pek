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
	redCircle  *entity.Object
	blueCircle *entity.Object
	pump1      *entity.Object
	pump2      *entity.Object
	character  *entity.Character
	wall1      *entity.Object
	wall2      *entity.Object
	allstat    []*entity.Object
}

func NewGame() *Game {
	g := &Game{}
	redCircle := entity.NewObject(20, 250, 10, 10, color.RGBA{255, 0, 0, 255})
	blueCircle := entity.NewObject(20, 300, 10, 10, color.RGBA{0, 0, 255, 255})
	entity.NewObjectSpd(redCircle, 10, 5)
	entity.NewObjectSpd(blueCircle, 5, 5)
	pump1 := entity.NewObject(100, 100, 100, 100, color.RGBA{0, 255, 0, 255})
	pump2 := entity.NewObject(300, 300, 100, 100, color.RGBA{0, 255, 0, 255})
	wall1 := entity.NewObject(500, 0, 40, 400, color.RGBA{26, 35, 73, 100})
	wall2 := entity.NewObject(600, screenHeight-400, 40, 400, color.RGBA{26, 35, 73, 100})
	character := entity.NewCharacter("Артем", 200, 300, 20, 20, color.RGBA{R: 0, G: 0, B: 255, A: 255})
	g.character = character
	g.blueCircle = blueCircle
	g.pump2 = pump2
	g.redCircle = redCircle
	g.pump1 = pump1
	g.wall1 = wall1
	g.wall2 = wall2
	w1, err1 := wall1.Split()
	if err1!= nil {
		panic(err1)
	}
	g.allstat = w1
	physics.InitStat(g.allstat...)
	w2, err2 := wall2.Split()
	if err2!= nil {
		panic(err2)
	}
	g.allstat = w2
	physics.InitStat(g.allstat...)
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
	physics.DynCollision()
	physics.DynMove()
	physics.ScreenCollision(screenWidth, screenHeight)

	physics.CharacterMove(g.character)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 0, 0, 255})
	cords.DebugCoordsObject(screen, g.redCircle)
	cords.DebugCharacter(screen, g.character, g.allstat)
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
	// Стены
	opWall1 := &ebiten.DrawImageOptions{}
	opWall1.GeoM.Translate(g.wall1.X, g.wall1.Y)
	screen.DrawImage(g.wall1.Img, opWall1)
	opWall2 := &ebiten.DrawImageOptions{}
	opWall2.GeoM.Translate(g.wall2.X, g.wall2.Y)
	screen.DrawImage(g.wall2.Img, opWall2)
	//Вектора
	cords.DebugCollision(screen, g.character, g.pump1, g.pump2)

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
