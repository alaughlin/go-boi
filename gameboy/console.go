package gameboy

import (
	"io/ioutil"
	"time"
)

// Console holds all the moving parts
type Console struct {
	cpu    *cpu
	memory *memory
}

// InitializeConsole initializes all the values a Gameboy needs to start
func InitializeConsole() *Console {
	return &Console{
		cpu:    initializeCPU(),
		memory: initializeMemory(),
	}
}

// LoadGame takes a path to a ROM and loads it into memory
func (console *Console) LoadGame(path string) {
	romData, err := ioutil.ReadFile(path)
	if err != nil {
		panic("ROM not found")
	}

	for i := 0; i < 32768; i++ {
		console.memory.write(uint16(i), romData[i])
	}
}

// Start begins the main loop
func (console *Console) Start() {
	for {
		console.cpu.ExecuteOpcode(console.memory)
		time.Sleep(time.Duration(1000))
	}
}
