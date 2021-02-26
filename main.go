package main

import (
	"image/color"
	"log"

	"github.com/alaughlin/go-boi/gameboy"
	"github.com/hajimehoshi/ebiten/v2"
)

type App struct {
	Gameboy *gameboy.Console
}

// Update executes 60 times/second
func (g *App) Update() error {
	g.Gameboy.Tick()
	return nil
}

// Draw ...draws to the screen
func (g *App) Draw(screen *ebiten.Image) {
	square := ebiten.NewImage(10, 10)
	square.Fill(color.White)
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(150, 134)
	screen.DrawImage(square, options)
}

// Layout defines the internal resolution which is later scaled
func (g *App) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 160, 144
}

func main() {
	app := &App{
		Gameboy: gameboy.InitializeConsole("./roms/tetris.gb"),
	}
	ebiten.SetWindowSize(640, 576)
	ebiten.SetWindowTitle("GoBoi")
	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(app); err != nil {
		log.Fatal(err)
	}
}
