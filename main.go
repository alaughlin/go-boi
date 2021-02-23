package main

import "github.com/alaughlin/go-boi/gameboy"

func main() {
	gameboy := gameboy.InitializeConsole()
	gameboy.LoadGame("./roms/tetris.gb")
	gameboy.Start()
}
