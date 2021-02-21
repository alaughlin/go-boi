package gameboy

type memory struct {
	bank0      *[]byte
	bank1      *[]byte
	vram       *[]byte
	eram       *[]byte
	wram0      *[]byte
	wram1      *[]byte
	oam        *[]byte
	io         *[]byte
	hram       *[]byte
	interrupts *byte
}

func initializeMemory() *memory {
	bank0 := make([]byte, 16384)
	bank1 := make([]byte, 16384)
	vram := make([]byte, 8192)
	eram := make([]byte, 8192)
	wram0 := make([]byte, 4096)
	wram1 := make([]byte, 4096)
	oam := make([]byte, 160)
	io := make([]byte, 128)
	hram := make([]byte, 126)
	interrupts := byte(0)

	memory := &memory{
		bank0:      &bank0,
		bank1:      &bank1,
		vram:       &vram,
		eram:       &eram,
		wram0:      &wram0,
		wram1:      &wram1,
		oam:        &oam,
		io:         &io,
		hram:       &hram,
		interrupts: &interrupts,
	}

	memory.initializeValues()

	return memory
}

func (memory *memory) initializeValues() {
	memory.write(0xFF10, 0x80)
	memory.write(0xFF11, 0xBF)
	memory.write(0xFF12, 0xF3)
	memory.write(0xFF14, 0xBF)
	memory.write(0xFF16, 0x3F)
	memory.write(0xFF19, 0xBF)
	memory.write(0xFF1A, 0x7F)
	memory.write(0xFF1B, 0xFF)
	memory.write(0xFF1C, 0x9F)
	memory.write(0xFF1E, 0xBF)
	memory.write(0xFF20, 0xFF)
	memory.write(0xFF23, 0xBF)
	memory.write(0xFF24, 0x77)
	memory.write(0xFF25, 0xF3)
	memory.write(0xFF26, 0xF1)
	memory.write(0xFF40, 0x91)
	memory.write(0xFF47, 0xFC)
	memory.write(0xFF48, 0xFF)
	memory.write(0xFF49, 0xFF)
}

func (memory *memory) read(address uint16) byte {
	slice, offset := memory.mapAddress(address)
	return (*slice)[address-offset]
}

func (memory *memory) write(address uint16, value byte) {
	slice, offset := memory.mapAddress(address)
	(*slice)[address-offset] = value
}

func (memory *memory) decrement(address uint16) {
	slice, offset := memory.mapAddress(address)
	(*slice)[address-offset]--
}

func (memory *memory) increment(address uint16) {
	slice, offset := memory.mapAddress(address)
	(*slice)[address-offset]++
}

func (memory *memory) mapAddress(address uint16) (*[]byte, uint16) {
	if address < 0x4000 {
		return memory.bank0, 0
	} else if address < 0x8000 {
		return memory.bank1, 0x4000
	} else if address < 0xA000 {
		return memory.vram, 0x8000
	} else if address < 0xC000 {
		return memory.eram, 0xA000
	} else if address < 0xD000 {
		return memory.wram0, 0xC000
	} else if address < 0xE000 {
		return memory.wram1, 0xD000
	} else if address < 0xFE00 {
		panic("accessing echo ram")
	} else if address < 0xFEA0 {
		return memory.oam, 0xFE00
	} else if address < 0xFF00 {
		panic("accessing unusable ram")
	} else if address < 0xFF80 {
		return memory.io, 0xFF00
	} else if address < 0xFFFF {
		return memory.hram, 0xFF80
	} else {
		panic("ram address out of bounds")
	}
}
