package gameboy

type cpu struct {
	a  *byte
	b  *byte
	c  *byte
	d  *byte
	e  *byte
	f  *byte
	h  *byte
	l  *byte
	sp uint16
	pc uint16
}

var (
	flagZ int = 7
	flagN int = 6
	flagH int = 5
	flagC int = 4
)

func initializeCPU() *cpu {
	c := byte(0x13)
	e := byte(0xD8)
	f := byte(0xB0)
	h := byte(0x01)
	l := byte(0x4D)

	return &cpu{
		c:  &c,
		e:  &e,
		f:  &f,
		h:  &h,
		l:  &l,
		sp: 0xFFFE,
		pc: 0x100,
	}
}

func (cpu *cpu) af() uint16 {
	return double(*cpu.a, *cpu.f)
}

func (cpu *cpu) bc() uint16 {
	return double(*cpu.b, *cpu.c)
}

func (cpu *cpu) de() uint16 {
	return double(*cpu.d, *cpu.e)
}

func (cpu *cpu) hl() uint16 {
	return double(*cpu.h, *cpu.l)
}

func double(n1 byte, n2 byte) uint16 {
	return uint16(n1)<<8 | uint16(n2)
}

func (cpu *cpu) setFlag(flagPos int) {
	*cpu.f |= (1 << flagPos)
}

func (cpu *cpu) clearFlag(flagPos int) {
	*cpu.f &^= (1 << flagPos)
}

func (cpu *cpu) ExecuteOpcode(memory *memory) {
	opcode := memory.read(cpu.pc)

	switch opcode {
	// 8-bit loads
	case 0x06:
		cpu.ld_r(cpu.b, memory.read(cpu.pc+1), 2)
	case 0x0E:
		cpu.ld_r(cpu.c, memory.read(cpu.pc+1), 2)
	case 0x16:
		cpu.ld_r(cpu.d, memory.read(cpu.pc+1), 2)
	case 0x1E:
		cpu.ld_r(cpu.e, memory.read(cpu.pc+1), 2)
	case 0x26:
		cpu.ld_r(cpu.h, memory.read(cpu.pc+1), 2)
	case 0x2E:
		cpu.ld_r(cpu.l, memory.read(cpu.pc+1), 2)
	case 0x7F:
		cpu.ld_r(cpu.a, *cpu.a, 1)
	case 0x78:
		cpu.ld_r(cpu.a, *cpu.b, 1)
	case 0x79:
		cpu.ld_r(cpu.a, *cpu.c, 1)
	case 0x7A:
		cpu.ld_r(cpu.a, *cpu.d, 1)
	case 0x7B:
		cpu.ld_r(cpu.a, *cpu.e, 1)
	case 0x7C:
		cpu.ld_r(cpu.a, *cpu.h, 1)
	case 0x7D:
		cpu.ld_r(cpu.a, *cpu.l, 1)
	case 0x7E:
		cpu.ld_r(cpu.a, memory.read(cpu.hl()), 1)
	case 0x40:
		cpu.ld_r(cpu.b, *cpu.b, 1)
	case 0x41:
		cpu.ld_r(cpu.b, *cpu.c, 1)
	case 0x42:
		cpu.ld_r(cpu.b, *cpu.d, 1)
	case 0x43:
		cpu.ld_r(cpu.b, *cpu.e, 1)
	case 0x44:
		cpu.ld_r(cpu.b, *cpu.h, 1)
	case 0x45:
		cpu.ld_r(cpu.b, *cpu.l, 1)
	case 0x46:
		cpu.ld_r(cpu.b, memory.read(cpu.hl()), 1)
	case 0x48:
		cpu.ld_r(cpu.c, *cpu.b, 1)
	case 0x49:
		cpu.ld_r(cpu.c, *cpu.c, 1)
	case 0x4A:
		cpu.ld_r(cpu.c, *cpu.d, 1)
	case 0x4B:
		cpu.ld_r(cpu.c, *cpu.e, 1)
	case 0x4C:
		cpu.ld_r(cpu.c, *cpu.h, 1)
	case 0x4D:
		cpu.ld_r(cpu.c, *cpu.l, 1)
	case 0x4E:
		cpu.ld_r(cpu.c, memory.read(cpu.hl()), 1)
	case 0x50:
		cpu.ld_r(cpu.d, *cpu.b, 1)
	case 0x51:
		cpu.ld_r(cpu.d, *cpu.c, 1)
	case 0x52:
		cpu.ld_r(cpu.d, *cpu.d, 1)
	case 0x53:
		cpu.ld_r(cpu.d, *cpu.e, 1)
	case 0x54:
		cpu.ld_r(cpu.d, *cpu.h, 1)
	case 0x55:
		cpu.ld_r(cpu.d, *cpu.l, 1)
	case 0x56:
		cpu.ld_r(cpu.d, memory.read(cpu.hl()), 1)
	case 0x58:
		cpu.ld_r(cpu.e, *cpu.b, 1)
	case 0x59:
		cpu.ld_r(cpu.e, *cpu.c, 1)
	case 0x5A:
		cpu.ld_r(cpu.e, *cpu.d, 1)
	case 0x5B:
		cpu.ld_r(cpu.e, *cpu.e, 1)
	case 0x5C:
		cpu.ld_r(cpu.e, *cpu.h, 1)
	case 0x5D:
		cpu.ld_r(cpu.e, *cpu.l, 1)
	case 0x5E:
		cpu.ld_r(cpu.e, memory.read(cpu.hl()), 1)
	case 0x60:
		cpu.ld_r(cpu.h, *cpu.b, 1)
	case 0x61:
		cpu.ld_r(cpu.h, *cpu.c, 1)
	case 0x62:
		cpu.ld_r(cpu.h, *cpu.d, 1)
	case 0x63:
		cpu.ld_r(cpu.h, *cpu.e, 1)
	case 0x64:
		cpu.ld_r(cpu.h, *cpu.h, 1)
	case 0x65:
		cpu.ld_r(cpu.h, *cpu.l, 1)
	case 0x66:
		cpu.ld_r(cpu.h, memory.read(cpu.hl()), 1)
	case 0x68:
		cpu.ld_r(cpu.l, *cpu.b, 1)
	case 0x69:
		cpu.ld_r(cpu.l, *cpu.c, 1)
	case 0x6A:
		cpu.ld_r(cpu.l, *cpu.d, 1)
	case 0x6B:
		cpu.ld_r(cpu.l, *cpu.e, 1)
	case 0x6C:
		cpu.ld_r(cpu.l, *cpu.h, 1)
	case 0x6D:
		cpu.ld_r(cpu.l, *cpu.l, 1)
	case 0x6E:
		cpu.ld_r(cpu.l, memory.read(cpu.hl()), 1)
	case 0x70:
		cpu.ld_addr(cpu.hl(), *cpu.b, memory, 1)
	case 0x71:
		cpu.ld_addr(cpu.hl(), *cpu.c, memory, 1)
	case 0x72:
		cpu.ld_addr(cpu.hl(), *cpu.d, memory, 1)
	case 0x73:
		cpu.ld_addr(cpu.hl(), *cpu.e, memory, 1)
	case 0x74:
		cpu.ld_addr(cpu.hl(), *cpu.h, memory, 1)
	case 0x75:
		cpu.ld_addr(cpu.hl(), *cpu.l, memory, 1)
	case 0x36:
		cpu.ld_addr(cpu.hl(), memory.read(cpu.pc+1), memory, 2)
	case 0x0A:
		cpu.ld_r(cpu.a, memory.read(cpu.bc()), 1)
	case 0x1A:
		cpu.ld_r(cpu.a, memory.read(cpu.de()), 1)
	case 0xFA:
		cpu.ld_r(cpu.a, memory.read(uint16(cpu.pc+2)<<8|uint16(cpu.pc+1)), 3)
	case 0x3E:
		cpu.ld_r(cpu.a, memory.read(cpu.pc+1), 2)
	case 0x47:
		cpu.ld_r(cpu.b, *cpu.a, 1)
	case 0x4F:
		cpu.ld_r(cpu.c, *cpu.a, 1)
	case 0x57:
		cpu.ld_r(cpu.d, *cpu.a, 1)
	case 0x5F:
		cpu.ld_r(cpu.e, *cpu.a, 1)
	case 0x67:
		cpu.ld_r(cpu.h, *cpu.a, 1)
	case 0x6F:
		cpu.ld_r(cpu.l, *cpu.a, 1)
	case 0x02:
		cpu.ld_addr(cpu.bc(), *cpu.a, memory, 1)
	case 0x12:
		cpu.ld_addr(cpu.de(), *cpu.a, memory, 1)
	case 0x77:
		cpu.ld_addr(cpu.hl(), *cpu.a, memory, 1)
	case 0xEA:
		cpu.ld_addr(uint16(cpu.pc+2)<<8|uint16(cpu.pc+1), *cpu.a, memory, 3)
	case 0xF2:
		cpu.ld_r(cpu.a, memory.read(0xFF00+uint16(*cpu.c)), 1)
	case 0xE2:
		cpu.ld_addr(0xFF00+uint16(*cpu.c), *cpu.a, memory, 1)
	case 0x3A:
		cpu.ld_r(cpu.a, memory.read(cpu.hl()), 1)
		memory.decrement(cpu.hl())
	case 0x32:
		cpu.ld_addr(cpu.hl(), *cpu.a, memory, 1)
		memory.decrement(cpu.hl())
	case 0x2A:
		cpu.ld_r(cpu.a, memory.read(cpu.hl()), 1)
		memory.increment(cpu.hl())
	case 0x22:
		cpu.ld_addr(cpu.hl(), *cpu.a, memory, 1)
		memory.increment(cpu.hl())
	case 0xE0:
		cpu.ld_addr(0xFF00+uint16(memory.read(cpu.pc+1)), *cpu.a, memory, 1)
	case 0xF0:
		cpu.ld_r(cpu.a, memory.read(0xFF00+uint16(memory.read(cpu.pc+1))), 1)
	// 16-bit loads
	case 0x01:
		cpu.ld_r_double(cpu.b, cpu.c, memory.readDouble(cpu.pc+1), 3)
	case 0x11:
		cpu.ld_r_double(cpu.d, cpu.e, memory.readDouble(cpu.pc+1), 3)
	case 0x21:
		cpu.ld_r_double(cpu.h, cpu.l, memory.readDouble(cpu.pc+1), 3)
	case 0x31:
		cpu.ld_sp(memory.readDouble(cpu.pc+1), 3)
	case 0xF9:
		cpu.ld_sp(cpu.hl(), 1)
	case 0xF8:
		// TODO: this is awful. check this. also need to set flag bits.
		cpu.ld_r_double(cpu.h, cpu.l, uint16(int16(cpu.sp)+int16(memory.read(cpu.pc+1))), 2)
	case 0x08:
		cpu.ld_addr_double(double(memory.read(cpu.pc+1), memory.read(cpu.pc+2)), cpu.sp, memory, 3)
	case 0xF5:
		cpu.push(memory.readDouble(cpu.af()), memory)
	case 0xC5:
		cpu.push(memory.readDouble(cpu.bc()), memory)
	case 0xD5:
		cpu.push(memory.readDouble(cpu.de()), memory)
	case 0xE5:
		cpu.push(memory.readDouble(cpu.hl()), memory)
	case 0xF1:
		cpu.pop(cpu.a, cpu.f, memory)
	case 0xC1:
		cpu.pop(cpu.b, cpu.c, memory)
	case 0xD1:
		cpu.pop(cpu.d, cpu.e, memory)
	case 0xE1:
		cpu.pop(cpu.h, cpu.l, memory)
	}
}

// 8-bit loads
func (cpu *cpu) ld_r(r *byte, n byte, incrementBy uint16) {
	*r = n
	cpu.pc += incrementBy
}

func (cpu *cpu) ld_addr(addr uint16, n byte, memory *memory, incrementBy uint16) {
	memory.write(addr, n)
	cpu.pc += incrementBy
}

// 16-bit loads
func (cpu *cpu) ld_r_double(r1 *byte, r2 *byte, nn uint16, incrementBy uint16) {
	*r1 = uint8(nn >> 8)
	*r2 = uint8(nn & 0x00FF)
	cpu.pc += incrementBy
}

func (cpu *cpu) ld_addr_double(addr uint16, nn uint16, memory *memory, incrementBy uint16) {
	memory.writeDouble(addr, nn)
	cpu.pc += incrementBy
}

func (cpu *cpu) ld_sp(nn uint16, incrementBy uint16) {
	cpu.sp = nn
	cpu.pc += incrementBy
}

func (cpu *cpu) push(nn uint16, memory *memory) {
	cpu.sp--
	memory.write(cpu.sp, uint8(nn>>8))
	cpu.sp--
	memory.write(cpu.sp, uint8(nn&0x00FF))
}

func (cpu *cpu) pop(r1 *byte, r2 *byte, memory *memory) {
	*r1 = memory.read(cpu.sp)
	cpu.sp++
	*r2 = memory.read(cpu.sp)
	cpu.sp++
}
