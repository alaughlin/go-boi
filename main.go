package main

import (
	"image/color"
	"log"

	"github.com/alaughlin/go-boi/gameboy"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	width           = 160
	height          = 144
	scaleFactor     = 4
	cyclesPerUpdate = 69905
)

var (
	colors map[int]color.RGBA = map[int]color.RGBA{0x3: {255, 255, 255, 255}, 0x2: {170, 170, 170, 255}, 0x1: {85, 85, 85, 255}, 0x0: {0, 0, 0, 255}}
)

// App holds the gameboy
type App struct {
	Gameboy *gameboy.Console
}

// Update executes 60 times/second
func (g *App) Update() error {
	cycles := 0
	for cycles < cyclesPerUpdate {
		cycles += g.Gameboy.Tick()
	}
	return nil
}

// Draw takes the display data and draws it to the screen
func (g *App) Draw(screen *ebiten.Image) {
	screen.Fill(color.White)
	square := ebiten.NewImage(100, 100)
	square.Fill(colors[0x2])
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(0, 0)
	screen.DrawImage(square, options)
}

// Layout defines the internal resolution which is later scaled
func (g *App) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return width, height
}

func main() {
	app := &App{
		Gameboy: gameboy.InitializeConsole("./roms/blargg/03-op sp,hl.gb", width, height),
	}
	ebiten.SetWindowSize(width*scaleFactor, height*scaleFactor)
	ebiten.SetWindowTitle("GoBoi")
	if err := ebiten.RunGame(app); err != nil {
		log.Fatal(err)
	}
}
