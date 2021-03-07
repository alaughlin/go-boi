package gameboy

import (
	"io/ioutil"
)

// Console holds all the moving parts
type Console struct {
	cpu    *cpu
	memory *memory
}

// InitializeConsole initializes all the values a Gameboy needs to start
func InitializeConsole(romPath string) *Console {
	console := &Console{
		cpu:    initializeCPU(),
		memory: initializeMemory(),
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

	for i := 0; i < 32768; i++ {
		console.memory.write(uint16(i), romData[i])
	}
}

func (console *Console) Tick() {
	console.cpu.ExecuteOpcode(console.memory)
}

func (console *Console) GetVRAM() *[]byte {
	return console.memory.getVRAM()
}
