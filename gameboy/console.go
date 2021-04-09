package gameboy

import (
	"io/ioutil"
)

// Console holds all the moving parts
type Console struct {
	cpu     *cpu
	memory  *memory
	display *display
}

// InitializeConsole initializes all the moving parts
func InitializeConsole(romPath string, width int, height int) *Console {
	console := &Console{
		cpu:     initializeCPU(),
		memory:  initializeMemory(),
		display: initalizeDisplay(width, height),
	}

	console.loadGame(romPath)
	return console
}

// LoadGame takes a path to a ROM and loads it into memory
func (console *Console) loadGame(path string) {
	romData, err := ioutil.ReadFile(path)
	if err != nil {
		panic("ROM not found")
	}

	console.memory.loadGame(romData)
}

// Tick executes a single instruction
func (console *Console) Tick() int {
	return console.cpu.ExecuteOpcode(console.memory)
}

// GetScreenData returns an array of rgba values to draw
func (console *Console) GetScreenData() []byte {
	return console.display.ScreenData
}
