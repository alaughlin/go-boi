package gameboy

import (
	"fmt"
	"math/bits"
)

type flags struct {
	Z, N, H, C bool
}

type cpu struct {
	a, b, c, d, e, h, l *byte
	sp, pc              uint16
	flags               *flags
	cycles              uint
}

var (
	flagZ int = 7
	flagN int = 6
	flagH int = 5
	flagC int = 4
)

func initializeCPU() *cpu {
	a := byte(0x00)
	b := byte(0x00)
	c := byte(0x13)
	d := byte(0x00)
	e := byte(0xD8)
	h := byte(0x01)
	l := byte(0x4D)
	flags := flags{Z: false, N: false, H: false, C: false}

	return &cpu{
		a:     &a,
		b:     &b,
		c:     &c,
		d:     &d,
		e:     &e,
		h:     &h,
		l:     &l,
		sp:    0xFFFE,
		pc:    0x100,
		flags: &flags,
	}
}

func flagsToByte(flags flags) uint8 {
	var f uint8

	if flags.Z {
		setBit(&f, flagZ)
	}
	if flags.N {
		setBit(&f, flagN)
	}
	if flags.H {
		setBit(&f, flagH)
	}
	if flags.C {
		setBit(&f, flagC)
	}

	return f
}

func byteToFlags(n byte) flags {
	flags := flags{}

	if getBit(n, flagZ) == 1 {
		flags.Z = true
	}
	if getBit(n, flagN) == 1 {
		flags.N = true
	}
	if getBit(n, flagH) == 1 {
		flags.H = true
	}
	if getBit(n, flagC) == 1 {
		flags.C = true
	}

	return flags
}

func getBit(n byte, bitPos int) byte {
	return (n >> bitPos) & 1
}

func setBit(n *byte, bitPos int) {
	*n |= (1 << bitPos)
}

func clearBit(n *byte, bitPos int) {
	*n &^= (1 << bitPos)
}

func (cpu *cpu) af() uint16 {
	f := flagsToByte(*cpu.flags)
	return double(*cpu.a, f)
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

func (cpu *cpu) ExecuteOpcode(memory *memory) {
	opcode := memory.read(cpu.pc)
	fmt.Printf("%X: %X\n", cpu.pc, opcode)

	switch opcode {
	case 0x06:
		cpu.ld_r(cpu.b, memory.read(cpu.pc+1), 2, 8)
	case 0x0E:
		cpu.ld_r(cpu.c, memory.read(cpu.pc+1), 2, 8)
	case 0x16:
		cpu.ld_r(cpu.d, memory.read(cpu.pc+1), 2, 8)
	case 0x1E:
		cpu.ld_r(cpu.e, memory.read(cpu.pc+1), 2, 8)
	case 0x26:
		cpu.ld_r(cpu.h, memory.read(cpu.pc+1), 2, 8)
	case 0x2E:
		cpu.ld_r(cpu.l, memory.read(cpu.pc+1), 2, 8)
	case 0x7F:
		cpu.ld_r(cpu.a, *cpu.a, 1, 4)
	case 0x78:
		cpu.ld_r(cpu.a, *cpu.b, 1, 4)
	case 0x79:
		cpu.ld_r(cpu.a, *cpu.c, 1, 4)
	case 0x7A:
		cpu.ld_r(cpu.a, *cpu.d, 1, 4)
	case 0x7B:
		cpu.ld_r(cpu.a, *cpu.e, 1, 4)
	case 0x7C:
		cpu.ld_r(cpu.a, *cpu.h, 1, 4)
	case 0x7D:
		cpu.ld_r(cpu.a, *cpu.l, 1, 4)
	case 0x7E:
		cpu.ld_r(cpu.a, memory.read(cpu.hl()), 1, 8)
	case 0x40:
		cpu.ld_r(cpu.b, *cpu.b, 1, 4)
	case 0x41:
		cpu.ld_r(cpu.b, *cpu.c, 1, 4)
	case 0x42:
		cpu.ld_r(cpu.b, *cpu.d, 1, 4)
	case 0x43:
		cpu.ld_r(cpu.b, *cpu.e, 1, 4)
	case 0x44:
		cpu.ld_r(cpu.b, *cpu.h, 1, 4)
	case 0x45:
		cpu.ld_r(cpu.b, *cpu.l, 1, 4)
	case 0x46:
		cpu.ld_r(cpu.b, memory.read(cpu.hl()), 1, 8)
	case 0x48:
		cpu.ld_r(cpu.c, *cpu.b, 1, 4)
	case 0x49:
		cpu.ld_r(cpu.c, *cpu.c, 1, 4)
	case 0x4A:
		cpu.ld_r(cpu.c, *cpu.d, 1, 4)
	case 0x4B:
		cpu.ld_r(cpu.c, *cpu.e, 1, 4)
	case 0x4C:
		cpu.ld_r(cpu.c, *cpu.h, 1, 4)
	case 0x4D:
		cpu.ld_r(cpu.c, *cpu.l, 1, 4)
	case 0x4E:
		cpu.ld_r(cpu.c, memory.read(cpu.hl()), 1, 8)
	case 0x50:
		cpu.ld_r(cpu.d, *cpu.b, 1, 4)
	case 0x51:
		cpu.ld_r(cpu.d, *cpu.c, 1, 4)
	case 0x52:
		cpu.ld_r(cpu.d, *cpu.d, 1, 4)
	case 0x53:
		cpu.ld_r(cpu.d, *cpu.e, 1, 4)
	case 0x54:
		cpu.ld_r(cpu.d, *cpu.h, 1, 4)
	case 0x55:
		cpu.ld_r(cpu.d, *cpu.l, 1, 4)
	case 0x56:
		cpu.ld_r(cpu.d, memory.read(cpu.hl()), 1, 8)
	case 0x58:
		cpu.ld_r(cpu.e, *cpu.b, 1, 4)
	case 0x59:
		cpu.ld_r(cpu.e, *cpu.c, 1, 4)
	case 0x5A:
		cpu.ld_r(cpu.e, *cpu.d, 1, 4)
	case 0x5B:
		cpu.ld_r(cpu.e, *cpu.e, 1, 4)
	case 0x5C:
		cpu.ld_r(cpu.e, *cpu.h, 1, 4)
	case 0x5D:
		cpu.ld_r(cpu.e, *cpu.l, 1, 4)
	case 0x5E:
		cpu.ld_r(cpu.e, memory.read(cpu.hl()), 1, 8)
	case 0x60:
		cpu.ld_r(cpu.h, *cpu.b, 1, 4)
	case 0x61:
		cpu.ld_r(cpu.h, *cpu.c, 1, 4)
	case 0x62:
		cpu.ld_r(cpu.h, *cpu.d, 1, 4)
	case 0x63:
		cpu.ld_r(cpu.h, *cpu.e, 1, 4)
	case 0x64:
		cpu.ld_r(cpu.h, *cpu.h, 1, 4)
	case 0x65:
		cpu.ld_r(cpu.h, *cpu.l, 1, 4)
	case 0x66:
		cpu.ld_r(cpu.h, memory.read(cpu.hl()), 1, 8)
	case 0x68:
		cpu.ld_r(cpu.l, *cpu.b, 1, 4)
	case 0x69:
		cpu.ld_r(cpu.l, *cpu.c, 1, 4)
	case 0x6A:
		cpu.ld_r(cpu.l, *cpu.d, 1, 4)
	case 0x6B:
		cpu.ld_r(cpu.l, *cpu.e, 1, 4)
	case 0x6C:
		cpu.ld_r(cpu.l, *cpu.h, 1, 4)
	case 0x6D:
		cpu.ld_r(cpu.l, *cpu.l, 1, 4)
	case 0x6E:
		cpu.ld_r(cpu.l, memory.read(cpu.hl()), 1, 8)
	case 0x70:
		cpu.ld_addr(cpu.hl(), *cpu.b, memory, 1, 8)
	case 0x71:
		cpu.ld_addr(cpu.hl(), *cpu.c, memory, 1, 8)
	case 0x72:
		cpu.ld_addr(cpu.hl(), *cpu.d, memory, 1, 8)
	case 0x73:
		cpu.ld_addr(cpu.hl(), *cpu.e, memory, 1, 8)
	case 0x74:
		cpu.ld_addr(cpu.hl(), *cpu.h, memory, 1, 8)
	case 0x75:
		cpu.ld_addr(cpu.hl(), *cpu.l, memory, 1, 8)
	case 0x36:
		cpu.ld_addr(cpu.hl(), memory.read(cpu.pc+1), memory, 2, 12)
	case 0x0A:
		cpu.ld_r(cpu.a, memory.read(cpu.bc()), 1, 8)
	case 0x1A:
		cpu.ld_r(cpu.a, memory.read(cpu.de()), 1, 8)
	case 0xFA:
		cpu.ld_r(cpu.a, memory.read(uint16(cpu.pc+2)<<8|uint16(cpu.pc+1)), 3, 16)
	case 0x3E:
		cpu.ld_r(cpu.a, memory.read(cpu.pc+1), 2, 8)
	case 0x47:
		cpu.ld_r(cpu.b, *cpu.a, 1, 4)
	case 0x4F:
		cpu.ld_r(cpu.c, *cpu.a, 1, 4)
	case 0x57:
		cpu.ld_r(cpu.d, *cpu.a, 1, 4)
	case 0x5F:
		cpu.ld_r(cpu.e, *cpu.a, 1, 4)
	case 0x67:
		cpu.ld_r(cpu.h, *cpu.a, 1, 4)
	case 0x6F:
		cpu.ld_r(cpu.l, *cpu.a, 1, 4)
	case 0x02:
		cpu.ld_addr(cpu.bc(), *cpu.a, memory, 1, 8)
	case 0x12:
		cpu.ld_addr(cpu.de(), *cpu.a, memory, 1, 8)
	case 0x77:
		cpu.ld_addr(cpu.hl(), *cpu.a, memory, 1, 8)
	case 0xEA:
		cpu.ld_addr(uint16(cpu.pc+2)<<8|uint16(cpu.pc+1), *cpu.a, memory, 3, 16)
	case 0xF2:
		cpu.ld_r(cpu.a, memory.read(0xFF00+uint16(*cpu.c)), 1, 8)
	case 0xE2:
		cpu.ld_addr(0xFF00+uint16(*cpu.c), *cpu.a, memory, 1, 8)
	case 0x3A:
		cpu.ld_r(cpu.a, memory.read(cpu.hl()), 1, 8)
		memory.decrement(cpu.hl())
	case 0x32:
		cpu.ld_addr(cpu.hl(), *cpu.a, memory, 1, 8)
		memory.decrement(cpu.hl())
	case 0x2A:
		cpu.ld_r(cpu.a, memory.read(cpu.hl()), 1, 8)
		memory.increment(cpu.hl())
	case 0x22:
		cpu.ld_addr(cpu.hl(), *cpu.a, memory, 1, 8)
		memory.increment(cpu.hl())
	case 0xE0:
		cpu.ld_addr(0xFF00+uint16(memory.read(cpu.pc+1)), *cpu.a, memory, 2, 12)
	case 0xF0:
		cpu.ld_r(cpu.a, memory.read(0xFF00+uint16(memory.read(cpu.pc+1))), 2, 12)
	case 0x01:
		cpu.ld_r_double(cpu.b, cpu.c, memory.readDouble(cpu.pc+1), 3, 12)
	case 0x11:
		cpu.ld_r_double(cpu.d, cpu.e, memory.readDouble(cpu.pc+1), 3, 12)
	case 0x21:
		cpu.ld_r_double(cpu.h, cpu.l, memory.readDouble(cpu.pc+1), 3, 12)
	case 0x31:
		cpu.ld_sp(memory.readDouble(cpu.pc+1), 3, 12)
	case 0xF9:
		cpu.ld_sp(cpu.hl(), 1, 8)
	case 0xF8:
		cpu.ld_r_double(cpu.h, cpu.l, cpu.sp+uint16(memory.read(cpu.pc+1)), 2, 12)
	case 0x08:
		cpu.ld_addr_double(double(memory.read(cpu.pc+1), memory.read(cpu.pc+2)), cpu.sp&0x00FF<<8|cpu.sp>>8, memory, 3, 20)
	case 0xF5:
		cpu.push(memory.readDouble(cpu.af()), memory, 16)
	case 0xC5:
		cpu.push(memory.readDouble(cpu.bc()), memory, 16)
	case 0xD5:
		cpu.push(memory.readDouble(cpu.de()), memory, 16)
	case 0xE5:
		cpu.push(memory.readDouble(cpu.hl()), memory, 16)
	case 0xF1:
		cpu.pop_flag(cpu.a, cpu.flags, memory, 12)
	case 0xC1:
		cpu.pop(cpu.b, cpu.c, memory, 12)
	case 0xD1:
		cpu.pop(cpu.d, cpu.e, memory, 12)
	case 0xE1:
		cpu.pop(cpu.h, cpu.l, memory, 12)
	case 0x87:
		cpu.add_r(cpu.a, *cpu.a, 1, 4)
	case 0x80:
		cpu.add_r(cpu.a, *cpu.b, 1, 4)
	case 0x81:
		cpu.add_r(cpu.a, *cpu.c, 1, 4)
	case 0x82:
		cpu.add_r(cpu.a, *cpu.d, 1, 4)
	case 0x83:
		cpu.add_r(cpu.a, *cpu.e, 1, 4)
	case 0x84:
		cpu.add_r(cpu.a, *cpu.h, 1, 4)
	case 0x85:
		cpu.add_r(cpu.a, *cpu.l, 1, 4)
	case 0x86:
		cpu.add_r(cpu.a, memory.read(cpu.hl()), 1, 8)
	case 0xC6:
		cpu.add_r(cpu.a, memory.read(cpu.pc+1), 2, 8)
	case 0x8F:
		cpu.adc_r(cpu.a, *cpu.a, 1, 8)
	case 0x88:
		cpu.adc_r(cpu.a, *cpu.b, 1, 8)
	case 0x89:
		cpu.adc_r(cpu.a, *cpu.c, 1, 8)
	case 0x8A:
		cpu.adc_r(cpu.a, *cpu.d, 1, 8)
	case 0x8B:
		cpu.adc_r(cpu.a, *cpu.e, 1, 8)
	case 0x8C:
		cpu.adc_r(cpu.a, *cpu.h, 1, 8)
	case 0x8D:
		cpu.adc_r(cpu.a, *cpu.l, 1, 8)
	case 0x8E:
		cpu.adc_r(cpu.a, memory.read(cpu.hl()), 1, 8)
	case 0xCE:
		cpu.adc_r(cpu.a, memory.read(cpu.pc+1), 2, 8)
	case 0x97:
		cpu.sub_r(cpu.a, *cpu.a, 1, 4)
	case 0x90:
		cpu.sub_r(cpu.a, *cpu.b, 1, 4)
	case 0x91:
		cpu.sub_r(cpu.a, *cpu.c, 1, 4)
	case 0x92:
		cpu.sub_r(cpu.a, *cpu.d, 1, 4)
	case 0x93:
		cpu.sub_r(cpu.a, *cpu.e, 1, 4)
	case 0x94:
		cpu.sub_r(cpu.a, *cpu.h, 1, 4)
	case 0x95:
		cpu.sub_r(cpu.a, *cpu.l, 1, 4)
	case 0x96:
		cpu.sub_r(cpu.a, memory.read(cpu.hl()), 1, 8)
	case 0xD6:
		cpu.sub_r(cpu.a, memory.read(cpu.pc+1), 2, 8)
	case 0x9F:
		cpu.sbc_r(cpu.a, *cpu.a, 1, 4)
	case 0x98:
		cpu.sbc_r(cpu.a, *cpu.b, 1, 4)
	case 0x99:
		cpu.sbc_r(cpu.a, *cpu.c, 1, 4)
	case 0x9A:
		cpu.sbc_r(cpu.a, *cpu.d, 1, 4)
	case 0x9B:
		cpu.sbc_r(cpu.a, *cpu.e, 1, 4)
	case 0x9C:
		cpu.sbc_r(cpu.a, *cpu.h, 1, 4)
	case 0x9D:
		cpu.sbc_r(cpu.a, *cpu.l, 1, 4)
	case 0x9E:
		cpu.sbc_r(cpu.a, memory.read(cpu.hl()), 1, 8)
	case 0xA7:
		cpu.and_r(cpu.a, *cpu.a, 1, 4)
	case 0xA0:
		cpu.and_r(cpu.a, *cpu.b, 1, 4)
	case 0xA1:
		cpu.and_r(cpu.a, *cpu.c, 1, 4)
	case 0xA2:
		cpu.and_r(cpu.a, *cpu.d, 1, 4)
	case 0xA3:
		cpu.and_r(cpu.a, *cpu.e, 1, 4)
	case 0xA4:
		cpu.and_r(cpu.a, *cpu.h, 1, 4)
	case 0xA5:
		cpu.and_r(cpu.a, *cpu.l, 1, 4)
	case 0xA6:
		cpu.and_r(cpu.a, memory.read(cpu.hl()), 1, 8)
	case 0xE6:
		cpu.and_r(cpu.a, memory.read(cpu.pc+1), 2, 8)
	case 0xB7:
		cpu.or_r(cpu.a, *cpu.a, 1, 4)
	case 0xB0:
		cpu.or_r(cpu.a, *cpu.b, 1, 4)
	case 0xB1:
		cpu.or_r(cpu.a, *cpu.c, 1, 4)
	case 0xB2:
		cpu.or_r(cpu.a, *cpu.d, 1, 4)
	case 0xB3:
		cpu.or_r(cpu.a, *cpu.e, 1, 4)
	case 0xB4:
		cpu.or_r(cpu.a, *cpu.h, 1, 4)
	case 0xB5:
		cpu.or_r(cpu.a, *cpu.l, 1, 4)
	case 0xB6:
		cpu.or_r(cpu.a, memory.read(cpu.hl()), 1, 8)
	case 0xF6:
		cpu.or_r(cpu.a, memory.read(cpu.pc+1), 2, 8)
	case 0xAF:
		cpu.xor_r(cpu.a, *cpu.a, 1, 4)
	case 0xA8:
		cpu.xor_r(cpu.a, *cpu.b, 1, 4)
	case 0xA9:
		cpu.xor_r(cpu.a, *cpu.c, 1, 4)
	case 0xAA:
		cpu.xor_r(cpu.a, *cpu.d, 1, 4)
	case 0xAB:
		cpu.xor_r(cpu.a, *cpu.e, 1, 4)
	case 0xAC:
		cpu.xor_r(cpu.a, *cpu.h, 1, 4)
	case 0xAD:
		cpu.xor_r(cpu.a, *cpu.l, 1, 4)
	case 0xAE:
		cpu.xor_r(cpu.a, memory.read(cpu.hl()), 1, 8)
	case 0xEE:
		cpu.xor_r(cpu.a, memory.read(cpu.pc+1), 2, 8)
	case 0xBF:
		cpu.cp_r(cpu.a, *cpu.a, 1, 4)
	case 0xB8:
		cpu.cp_r(cpu.a, *cpu.b, 1, 4)
	case 0xB9:
		cpu.cp_r(cpu.a, *cpu.c, 1, 4)
	case 0xBA:
		cpu.cp_r(cpu.a, *cpu.d, 1, 4)
	case 0xBB:
		cpu.cp_r(cpu.a, *cpu.e, 1, 4)
	case 0xBC:
		cpu.cp_r(cpu.a, *cpu.h, 1, 4)
	case 0xBD:
		cpu.cp_r(cpu.a, *cpu.l, 1, 4)
	case 0xBE:
		cpu.cp_r(cpu.a, memory.read(cpu.hl()), 1, 8)
	case 0xFE:
		cpu.cp_r(cpu.a, memory.read(cpu.pc+1), 2, 8)
	case 0x3C:
		cpu.inc_r(cpu.a, 4)
	case 0x04:
		cpu.inc_r(cpu.b, 4)
	case 0x0C:
		cpu.inc_r(cpu.c, 4)
	case 0x14:
		cpu.inc_r(cpu.d, 4)
	case 0x1C:
		cpu.inc_r(cpu.e, 4)
	case 0x24:
		cpu.inc_r(cpu.h, 4)
	case 0x2C:
		cpu.inc_r(cpu.l, 4)
	case 0x34:
		cpu.inc_addr(cpu.hl(), memory, 12)
	case 0x3D:
		cpu.dec_r(cpu.a, 4)
	case 0x05:
		cpu.dec_r(cpu.b, 4)
	case 0x0D:
		cpu.dec_r(cpu.c, 4)
	case 0x15:
		cpu.dec_r(cpu.d, 4)
	case 0x1D:
		cpu.dec_r(cpu.e, 4)
	case 0x25:
		cpu.dec_r(cpu.h, 4)
	case 0x2D:
		cpu.dec_r(cpu.l, 4)
	case 0x35:
		cpu.dec_addr(cpu.hl(), memory, 12)
	case 0x09:
		cpu.add_r_double(cpu.h, cpu.l, cpu.bc(), 8)
	case 0x19:
		cpu.add_r_double(cpu.h, cpu.l, cpu.de(), 8)
	case 0x29:
		cpu.add_r_double(cpu.h, cpu.l, cpu.hl(), 8)
	case 0x39:
		cpu.add_r_double(cpu.h, cpu.l, cpu.sp, 8)
	case 0xE8:
		cpu.add_sp(memory.read(cpu.pc+1), 16)
	case 0x03:
		cpu.inc_r_double(cpu.b, cpu.c, 8)
	case 0x13:
		cpu.inc_r_double(cpu.d, cpu.e, 8)
	case 0x23:
		cpu.inc_r_double(cpu.h, cpu.l, 8)
	case 0x33:
		cpu.inc_sp(8)
	case 0x0B:
		cpu.dec_r_double(cpu.b, cpu.c, 8)
	case 0x1B:
		cpu.dec_r_double(cpu.d, cpu.e, 8)
	case 0x2B:
		cpu.dec_r_double(cpu.h, cpu.l, 8)
	case 0x3B:
		cpu.dec_sp(8)
	case 0xCB:
		switch memory.read(cpu.pc + 1) {
		case 0x37:
			cpu.swap_r(cpu.a, 8)
		case 0x30:
			cpu.swap_r(cpu.b, 8)
		case 0x31:
			cpu.swap_r(cpu.c, 8)
		case 0x32:
			cpu.swap_r(cpu.d, 8)
		case 0x33:
			cpu.swap_r(cpu.e, 8)
		case 0x34:
			cpu.swap_r(cpu.h, 8)
		case 0x35:
			cpu.swap_r(cpu.l, 8)
		case 0x36:
			cpu.swap_addr(cpu.hl(), memory, 16)
		case 0x07:
			cpu.rlc_r(cpu.a)
		case 0x00:
			cpu.rlc_r(cpu.b)
		case 0x01:
			cpu.rlc_r(cpu.c)
		case 0x02:
			cpu.rlc_r(cpu.d)
		case 0x03:
			cpu.rlc_r(cpu.e)
		case 0x04:
			cpu.rlc_r(cpu.h)
		case 0x05:
			cpu.rlc_r(cpu.l)
		case 0x06:
			cpu.rlc_addr(cpu.hl(), memory)
		case 0x17:
			cpu.rl_r(cpu.a)
		case 0x10:
			cpu.rl_r(cpu.b)
		case 0x11:
			cpu.rl_r(cpu.c)
		case 0x12:
			cpu.rl_r(cpu.d)
		case 0x13:
			cpu.rl_r(cpu.e)
		case 0x14:
			cpu.rl_r(cpu.h)
		case 0x15:
			cpu.rl_r(cpu.l)
		case 0x16:
			cpu.rl_addr(cpu.hl(), memory)
		case 0x0F:
			cpu.rrc_r(cpu.a)
		case 0x08:
			cpu.rrc_r(cpu.b)
		case 0x09:
			cpu.rrc_r(cpu.c)
		case 0x0A:
			cpu.rrc_r(cpu.d)
		case 0x0B:
			cpu.rrc_r(cpu.e)
		case 0x0C:
			cpu.rrc_r(cpu.h)
		case 0x0D:
			cpu.rrc_r(cpu.l)
		case 0x0E:
			cpu.rrc_addr(cpu.hl(), memory)
		case 0x1F:
			cpu.rr_r(cpu.a)
		case 0x18:
			cpu.rr_r(cpu.b)
		case 0x19:
			cpu.rr_r(cpu.c)
		case 0x1A:
			cpu.rr_r(cpu.d)
		case 0x1B:
			cpu.rr_r(cpu.e)
		case 0x1C:
			cpu.rr_r(cpu.h)
		case 0x1D:
			cpu.rr_r(cpu.l)
		case 0x1E:
			cpu.rr_addr(cpu.hl(), memory)
		case 0x27:
			cpu.sla_r(cpu.a)
		case 0x20:
			cpu.sla_r(cpu.b)
		case 0x21:
			cpu.sla_r(cpu.c)
		case 0x22:
			cpu.sla_r(cpu.d)
		case 0x23:
			cpu.sla_r(cpu.e)
		case 0x24:
			cpu.sla_r(cpu.h)
		case 0x25:
			cpu.sla_r(cpu.l)
		case 0x26:
			cpu.sla_addr(cpu.hl(), memory)
		case 0x2F:
			cpu.sra_r(cpu.a)
		case 0x28:
			cpu.sra_r(cpu.b)
		case 0x29:
			cpu.sra_r(cpu.c)
		case 0x2A:
			cpu.sra_r(cpu.d)
		case 0x2B:
			cpu.sra_r(cpu.e)
		case 0x2C:
			cpu.sra_r(cpu.h)
		case 0x2D:
			cpu.sra_r(cpu.l)
		case 0x2E:
			cpu.sra_addr(cpu.hl(), memory)
		case 0x3F:
			cpu.srl_r(cpu.a)
		case 0x38:
			cpu.srl_r(cpu.b)
		case 0x39:
			cpu.srl_r(cpu.c)
		case 0x3A:
			cpu.srl_r(cpu.d)
		case 0x3B:
			cpu.srl_r(cpu.e)
		case 0x3C:
			cpu.srl_r(cpu.h)
		case 0x3D:
			cpu.srl_r(cpu.l)
		case 0x3E:
			cpu.srl_addr(cpu.hl(), memory)
		case 0x47:
			cpu.bit_r(cpu.a, memory.read(cpu.pc+1))
		case 0x40:
			cpu.bit_r(cpu.b, memory.read(cpu.pc+1))
		case 0x41:
			cpu.bit_r(cpu.c, memory.read(cpu.pc+1))
		case 0x42:
			cpu.bit_r(cpu.d, memory.read(cpu.pc+1))
		case 0x43:
			cpu.bit_r(cpu.e, memory.read(cpu.pc+1))
		case 0x44:
			cpu.bit_r(cpu.h, memory.read(cpu.pc+1))
		case 0x45:
			cpu.bit_r(cpu.l, memory.read(cpu.pc+1))
		case 0x46:
			cpu.bit_addr(cpu.hl(), memory.read(cpu.pc+1), memory)
		case 0xC7:
			cpu.set_r(cpu.a, memory.read(cpu.pc+1))
		case 0xC0:
			cpu.set_r(cpu.b, memory.read(cpu.pc+1))
		case 0xC1:
			cpu.set_r(cpu.c, memory.read(cpu.pc+1))
		case 0xC2:
			cpu.set_r(cpu.d, memory.read(cpu.pc+1))
		case 0xC3:
			cpu.set_r(cpu.e, memory.read(cpu.pc+1))
		case 0xC4:
			cpu.set_r(cpu.h, memory.read(cpu.pc+1))
		case 0xC5:
			cpu.set_r(cpu.l, memory.read(cpu.pc+1))
		case 0xC6:
			cpu.set_addr(cpu.hl(), memory.read(cpu.pc+1), memory)
		case 0x87:
			cpu.res_r(cpu.a, memory.read(cpu.pc+1))
		case 0x80:
			cpu.res_r(cpu.b, memory.read(cpu.pc+1))
		case 0x81:
			cpu.res_r(cpu.c, memory.read(cpu.pc+1))
		case 0x82:
			cpu.res_r(cpu.d, memory.read(cpu.pc+1))
		case 0x83:
			cpu.res_r(cpu.e, memory.read(cpu.pc+1))
		case 0x84:
			cpu.res_r(cpu.h, memory.read(cpu.pc+1))
		case 0x85:
			cpu.res_r(cpu.l, memory.read(cpu.pc+1))
		case 0x86:
			cpu.res_addr(cpu.hl(), memory.read(cpu.pc+1), memory)
		default:
			panic(fmt.Sprintf("unknown instruction: CB %X", opcode))
		}
	case 0x27:
		cpu.da_r(cpu.a)
	case 0x2F:
		cpu.cpl_r(cpu.a, 4)
	case 0x3F:
		cpu.ccf(4)
	case 0x37:
		cpu.scf(4)
	case 0x00:
		cpu.nop(4)
	case 0x76:
		cpu.halt(4)
	case 0x10:
		cpu.stop(4)
	case 0xF3:
		cpu.di(4)
	case 0xFB:
		cpu.ei(4)
	case 0x07:
		cpu.rlca()
	case 0x17:
		cpu.rla()
	case 0x0F:
		cpu.rrca()
	case 0x1F:
		cpu.rra()
	case 0xC3:
		cpu.jp_nn(double(memory.read(cpu.pc+2), memory.read(cpu.pc+1)))
	case 0xC2:
		cpu.jp_nn_cc(double(memory.read(cpu.pc+2), memory.read(cpu.pc+1)), flagZ, 0)
	case 0xCA:
		cpu.jp_nn_cc(double(memory.read(cpu.pc+2), memory.read(cpu.pc+1)), flagZ, 1)
	case 0xD2:
		cpu.jp_nn_cc(double(memory.read(cpu.pc+2), memory.read(cpu.pc+1)), flagC, 0)
	case 0xDA:
		cpu.jp_nn_cc(double(memory.read(cpu.pc+2), memory.read(cpu.pc+1)), flagC, 1)
	case 0xE9:
		cpu.jp_nn(memory.readDouble(cpu.hl()))
	case 0x18:
		cpu.jr_n(memory.read(cpu.pc + 1))
	case 0x20:
		cpu.jr_n_cc(memory.read(cpu.pc+1), flagZ, 0)
	case 0x28:
		cpu.jr_n_cc(memory.read(cpu.pc+1), flagZ, 1)
	case 0x30:
		cpu.jr_n_cc(memory.read(cpu.pc+1), flagC, 0)
	case 0x38:
		cpu.jr_n_cc(memory.read(cpu.pc+1), flagC, 1)
	case 0xCD:
		cpu.call_nn(double(memory.read(cpu.pc+2), memory.read(cpu.pc+1)), cpu.pc+3, memory)
	case 0xC4:
		cpu.call_cc_nn(double(memory.read(cpu.pc+2), memory.read(cpu.pc+1)), cpu.pc+3, flagZ, 0, memory)
	case 0xCC:
		cpu.call_cc_nn(double(memory.read(cpu.pc+2), memory.read(cpu.pc+1)), cpu.pc+3, flagZ, 1, memory)
	case 0xD4:
		cpu.call_cc_nn(double(memory.read(cpu.pc+2), memory.read(cpu.pc+1)), cpu.pc+3, flagC, 0, memory)
	case 0xDC:
		cpu.call_cc_nn(double(memory.read(cpu.pc+2), memory.read(cpu.pc+1)), cpu.pc+3, flagC, 1, memory)
	case 0xC7:
		cpu.rst_n(0x00, memory)
	case 0xCF:
		cpu.rst_n(0x08, memory)
	case 0xD7:
		cpu.rst_n(0x10, memory)
	case 0xDF:
		cpu.rst_n(0x18, memory)
	case 0xE7:
		cpu.rst_n(0x20, memory)
	case 0xEF:
		cpu.rst_n(0x28, memory)
	case 0xF7:
		cpu.rst_n(0x30, memory)
	case 0xFF:
		cpu.rst_n(0x38, memory)
	case 0xC9:
		cpu.ret(memory)
	case 0xC0:
		cpu.ret_cc(flagZ, 0, memory)
	case 0xC8:
		cpu.ret_cc(flagZ, 1, memory)
	case 0xD0:
		cpu.ret_cc(flagC, 0, memory)
	case 0xD8:
		cpu.ret_cc(flagC, 1, memory)
	case 0xD9:
		cpu.reti()
	default:
		panic(fmt.Sprintf("unknown instruction: %X", opcode))
	}
}

// 8-Bit Loads
func (cpu *cpu) ld_r(r *byte, n byte, incrementBy uint16, cycles uint) {
	*r = n
	cpu.cycles = cycles
	cpu.pc += incrementBy
}

func (cpu *cpu) ld_addr(addr uint16, n byte, memory *memory, incrementBy uint16, cycles uint) {
	if addr == 0xFFFF {
		*cpu.flags = byteToFlags(n)
	} else {
		memory.write(addr, n)
	}
	cpu.cycles = cycles
	cpu.pc += incrementBy
}

// 16-Bit Loads
func (cpu *cpu) ld_r_double(r1 *byte, r2 *byte, nn uint16, incrementBy uint16, cycles uint) {
	*r1 = uint8(nn >> 8)
	*r2 = uint8(nn & 0x00FF)
	cpu.flags.Z = false
	cpu.flags.N = false
	cpu.cycles = cycles
	cpu.pc += incrementBy
}

func (cpu *cpu) ld_addr_double(addr uint16, nn uint16, memory *memory, incrementBy uint16, cycles uint) {
	memory.writeDouble(addr, nn)
	cpu.cycles = cycles
	cpu.pc += incrementBy
}

func (cpu *cpu) ld_sp(nn uint16, incrementBy uint16, cycles uint) {
	cpu.sp = nn
	cpu.cycles = cycles
	cpu.pc += incrementBy
}

func (cpu *cpu) push(nn uint16, memory *memory, cycles uint) {
	memory.write(cpu.sp-1, uint8(nn>>8))
	memory.write(cpu.sp-2, uint8(nn&0x00FF))
	cpu.sp -= 2
	cpu.cycles = cycles
	cpu.pc++
}

func (cpu *cpu) pop(r1 *byte, r2 *byte, memory *memory, cycles uint) {
	*r1 = memory.read(cpu.sp)
	*r2 = memory.read(cpu.sp + 1)
	cpu.sp += 2
	cpu.pc++
}

func (cpu *cpu) pop_flag(r1 *byte, flags *flags, memory *memory, cycles uint) {
	*r1 = memory.read(cpu.sp)
	cpu.sp++
	*flags = byteToFlags(memory.read(cpu.sp))
	cpu.sp++
	cpu.pc++
}

// 8-Bit ALU
func (cpu *cpu) add_r(r *byte, n byte, incrementBy uint16, cycles uint) {
	res := uint16(*r) + uint16(n)
	cpu.flags.Z = (res & 0xFF) == 0
	cpu.flags.N = false
	cpu.flags.H = res > 0xF
	cpu.flags.C = res > 0xFF
	*r = uint8(res)
	cpu.cycles = cycles
	cpu.pc += incrementBy
}

func (cpu *cpu) adc_r(r *byte, n byte, incrementBy uint16, cycles uint) {
	res := uint16(*r) + uint16(n) + (uint16(flagsToByte(*cpu.flags) & byte(flagC)))
	cpu.flags.Z = (res & 0xFF) == 0
	cpu.flags.N = false
	cpu.flags.H = res > 0xF
	cpu.flags.C = res > 0xFF
	*r = uint8(res)
	cpu.cycles = cycles
	cpu.pc += incrementBy
}

func (cpu *cpu) sub_r(r *byte, n byte, incrementBy uint16, cycles uint) {
	res := int16(*r) - int16(n)
	cpu.flags.Z = (res & 0xFF) == 0
	cpu.flags.N = true
	cpu.flags.H = (int8(*r&0xF) - int8(n&0xF)) < 0
	cpu.flags.C = res < 0
	*r = uint8(res)
	cpu.cycles = cycles
	cpu.pc += incrementBy
}

func (cpu *cpu) sbc_r(r *byte, n byte, incrementBy uint16, cycles uint) {
	res := uint16(*r) - uint16(n) - (uint16(flagsToByte(*cpu.flags) & byte(flagC)))
	cpu.flags.Z = (res & 0xFF) == 0
	cpu.flags.N = true
	cpu.flags.H = (int8(*r&0xF) - int8(n&0xF)) < 0
	cpu.flags.C = res < 0
	*r = uint8(res)
	cpu.cycles = cycles
	cpu.pc += incrementBy
}

func (cpu *cpu) and_r(r *byte, n byte, incrementBy uint16, cycles uint) {
	res := *r & n
	cpu.flags.Z = res == 0
	cpu.flags.N = false
	cpu.flags.H = true
	cpu.flags.C = false
	*r = res
	cpu.cycles = cycles
	cpu.pc += incrementBy
}

func (cpu *cpu) or_r(r *byte, n byte, incrementBy uint16, cycles uint) {
	res := *r | n
	cpu.flags.Z = res == 0
	cpu.flags.N = false
	cpu.flags.H = false
	cpu.flags.C = false
	*r = res
	cpu.cycles = cycles
	cpu.pc += incrementBy
}

func (cpu *cpu) xor_r(r *byte, n byte, incrementBy uint16, cycles uint) {
	res := *r ^ n
	cpu.flags.Z = res == 0
	cpu.flags.N = false
	cpu.flags.H = false
	cpu.flags.C = false
	*r = res
	cpu.cycles = cycles
	cpu.pc += incrementBy
}

func (cpu *cpu) cp_r(r *byte, n byte, incrementBy uint16, cycles uint) {
	cpu.flags.Z = *r == n
	cpu.flags.N = true
	cpu.flags.H = *r&0xF < n&0xF
	cpu.flags.C = *r < n
	cpu.cycles = cycles
	cpu.pc += incrementBy
}

func (cpu *cpu) inc_r(r *byte, cycles uint) {
	res := uint16(*r) + 1
	cpu.flags.Z = res == 0
	cpu.flags.N = false
	cpu.flags.H = res > 0xFF
	*r = uint8(res)
	cpu.cycles = cycles
	cpu.pc++
}

func (cpu *cpu) inc_addr(addr uint16, memory *memory, cycles uint) {
	res := uint16(memory.read(addr)) + 1
	cpu.flags.Z = res == 0
	cpu.flags.N = false
	cpu.flags.H = res > 0xFF
	memory.write(addr, uint8(res))
	cpu.cycles = cycles
	cpu.pc++
}

func (cpu *cpu) dec_r(r *byte, cycles uint) {
	res := int16(*r) - 1
	cpu.flags.Z = res == 0
	cpu.flags.N = true
	cpu.flags.H = res < 0
	*r = uint8(res)
	cpu.cycles = cycles
	cpu.pc++
}

func (cpu *cpu) dec_addr(addr uint16, memory *memory, cycles uint) {
	res := int16(memory.read(addr)) - 1
	cpu.flags.Z = res == 0
	cpu.flags.N = true
	cpu.flags.H = res < 0
	memory.write(addr, uint8(res))
	cpu.cycles = 12
	cpu.pc++
}

// 16-Bit ALU
func (cpu *cpu) add_r_double(r1 *byte, r2 *byte, nn uint16, cycles uint) {
	res := uint32(double(*r1, *r2)) + uint32(nn)
	cpu.flags.N = false
	cpu.flags.H = res > 0xFFF
	cpu.flags.C = res > 0xFFFF
	*r1 = uint8(uint16(res) >> 8)
	*r2 = uint8(uint16(res) & 0x00FF)
	cpu.cycles = cycles
	cpu.pc++
}

func (cpu *cpu) add_sp(n byte, cycles uint) {
	res := uint32(cpu.sp) + uint32(n)
	cpu.flags.Z = false
	cpu.flags.N = false
	cpu.flags.H = res > 0xFFF
	cpu.flags.C = res > 0xFFFF
	cpu.sp = uint16(res)
	cpu.cycles = cycles
	cpu.pc += 2
}

func (cpu *cpu) inc_r_double(r1 *byte, r2 *byte, cycles uint) {
	res := double(*r1, *r2) + 1
	*r1 = uint8(res >> 8)
	*r2 = uint8(res & 0x00FF)
	cpu.cycles = cycles
	cpu.pc++
}

func (cpu *cpu) inc_sp(cycles uint) {
	cpu.sp++
	cpu.cycles = cycles
	cpu.pc++
}

func (cpu *cpu) dec_r_double(r1 *byte, r2 *byte, cycles uint) {
	res := double(*r1, *r2) - 1
	*r1 = uint8(res >> 8)
	*r2 = uint8(res & 0x00FF)
	cpu.cycles = cycles
	cpu.pc++
}

func (cpu *cpu) dec_sp(cycles uint) {
	cpu.sp--
	cpu.cycles = cycles
	cpu.pc++
}

// Misc
func (cpu *cpu) swap_r(r *byte, cycles uint) {
	*r = *r&0x0F<<4 | *r>>4
	cpu.flags.Z = *r == 0
	cpu.flags.N = false
	cpu.flags.H = false
	cpu.flags.C = false
	cpu.pc += 2
}

func (cpu *cpu) swap_addr(addr uint16, memory *memory, cycles uint) {
	n := memory.read(addr)
	memory.write(addr, n&0x0F<<4|n>>4)
	cpu.flags.Z = n == 0
	cpu.flags.N = false
	cpu.flags.H = false
	cpu.flags.C = false
	cpu.cycles = cycles
	cpu.pc += 2
}

func (cpu *cpu) da_r(r *byte) {
	// TODO
	// TODO: set flags
	cpu.pc++
}

func (cpu *cpu) cpl_r(r *byte, cycles uint) {
	*r ^= *r
	cpu.flags.N = true
	cpu.flags.H = true
	cpu.cycles = cycles
	cpu.pc++
}

func (cpu *cpu) ccf(cycles uint) {
	cpu.flags.N = false
	cpu.flags.H = false
	cpu.flags.C = !cpu.flags.C
	cpu.cycles = cycles
	cpu.pc++
}

func (cpu *cpu) scf(cycles uint) {
	cpu.flags.N = false
	cpu.flags.H = false
	cpu.flags.C = true
	cpu.cycles = cycles
	cpu.pc++
}

func (cpu *cpu) nop(cycles uint) {
	cpu.cycles = cycles
	cpu.pc++
}

func (cpu *cpu) halt(cycles uint) {
	// TODO
	cpu.cycles = cycles
}

func (cpu *cpu) stop(cycles uint) {
	/// TODO: disable opcode execution until button pressed
	cpu.cycles = cycles
	cpu.pc++
}

func (cpu *cpu) di(cycles uint) {
	// TODO: disable interrupts AFTER following instruction
	cpu.cycles = cycles
	cpu.pc++
}

func (cpu *cpu) ei(cycles uint) {
	// TODO: enable interrupts AFTER following instruction
	cpu.cycles = cycles
	cpu.pc++
}

// Rotates and Shifts
func (cpu *cpu) rlca() {
	cpu.flags.Z = false
	cpu.flags.N = false
	cpu.flags.H = false
	cpu.flags.C = 128&*cpu.a == 1
	*cpu.a = byte(bits.RotateLeft8(uint8(*cpu.a), 1))
	cpu.pc++
}

func (cpu *cpu) rla() {
	oldC := cpu.flags.C
	cpu.flags.Z = false
	cpu.flags.N = false
	cpu.flags.H = false
	cpu.flags.C = 128&*cpu.a == 1
	*cpu.a = byte(bits.RotateLeft8(uint8(*cpu.a), 1))
	if oldC {
		setBit(cpu.a, 0)
	} else {
		clearBit(cpu.a, 0)
	}
	cpu.pc++
}

func (cpu *cpu) rrca() {
	cpu.flags.Z = false
	cpu.flags.N = false
	cpu.flags.H = false
	cpu.flags.C = 1&*cpu.a == 1
	*cpu.a = byte(bits.RotateLeft8(*cpu.a, -1))
	cpu.pc++
}

func (cpu *cpu) rra() {
	oldC := cpu.flags.C
	cpu.flags.Z = false
	cpu.flags.N = false
	cpu.flags.H = false
	cpu.flags.C = 1&*cpu.a == 1
	*cpu.a = byte(bits.RotateLeft8(*cpu.a, -1))
	if oldC {
		setBit(cpu.a, 7)
	} else {
		clearBit(cpu.a, 7)
	}
	cpu.pc++
}

func (cpu *cpu) rlc_r(r *byte) {
	cpu.flags.N = false
	cpu.flags.H = false
	cpu.flags.C = 128&*r == 1
	*r = byte(bits.RotateLeft8(uint8(*r), 1))
	cpu.flags.Z = *r == 0
	cpu.pc++
}

func (cpu *cpu) rlc_addr(addr uint16, memory *memory) {
	n := memory.read(addr)
	cpu.flags.C = 128&n == 0
	res := byte(bits.RotateLeft8(uint8(n), 1))
	cpu.flags.Z = res == 0
	cpu.flags.N = false
	cpu.flags.H = false
	memory.write(addr, res)
	cpu.pc++
}

func (cpu *cpu) rl_r(r *byte) {
	oldC := cpu.flags.C
	cpu.flags.N = false
	cpu.flags.H = false
	cpu.flags.C = 128&*r == 1
	*r = byte(bits.RotateLeft8(*r, 1))
	cpu.flags.Z = *r == 0
	if oldC {
		setBit(r, 0)
	} else {
		clearBit(r, 0)
	}
	cpu.pc++
}

func (cpu *cpu) rl_addr(addr uint16, memory *memory) {
	oldC := cpu.flags.C
	cpu.flags.N = false
	cpu.flags.H = false
	n := memory.read(addr)
	cpu.flags.C = 128&n != 0
	res := byte(bits.RotateLeft8(uint8(n), 1))
	cpu.flags.Z = res == 0
	if oldC {
		setBit(&res, 0)
	} else {
		setBit(&res, 0)
	}
	memory.write(addr, res)
	cpu.flags.Z = false
	cpu.flags.N = false
	cpu.flags.H = false
	cpu.pc++
}

func (cpu *cpu) rrc_r(r *byte) {
	cpu.flags.N = false
	cpu.flags.H = false
	cpu.flags.C = 1&*r == 1
	*r = byte(bits.RotateLeft8(*r, -1))
	cpu.flags.Z = *r == 0
	cpu.pc++
}

func (cpu *cpu) rr_r(r *byte) {
	oldC := flagsToByte(*cpu.flags) & 16
	cpu.flags.N = false
	cpu.flags.H = false
	cpu.flags.C = 1&*r == 1
	*r = byte(bits.RotateLeft8(*r, -1))
	cpu.flags.Z = *r == 0
	if oldC == 0 {
		clearBit(r, 7)
	} else {
		setBit(r, 7)
	}
	cpu.pc++
}

func (cpu *cpu) rrc_addr(addr uint16, memory *memory) {
	cpu.flags.N = false
	cpu.flags.H = false
	n := memory.read(addr)
	cpu.flags.C = 1&n != 0
	res := (bits.RotateLeft8(n, -1))
	cpu.flags.Z = res == 0
	memory.write(addr, res)
	cpu.pc++
}

func (cpu *cpu) rr_addr(addr uint16, memory *memory) {
	oldC := cpu.flags.C
	cpu.flags.N = false
	cpu.flags.H = false
	n := memory.read(addr)
	cpu.flags.C = 1&n != 0
	res := n >> 1
	cpu.flags.Z = res == 0
	if oldC {
		setBit(&n, 7)
	} else {
		clearBit(&n, 7)
	}
	memory.write(addr, res)
	cpu.pc++
}

func (cpu *cpu) sla_r(r *byte) {
	cpu.flags.C = 128&*r != 0
	*r <<= 1
	cpu.flags.Z = *r == 0
	cpu.flags.N = false
	cpu.flags.H = false
	cpu.pc++
}

func (cpu *cpu) sla_addr(addr uint16, memory *memory) {
	n := memory.read(addr)
	cpu.flags.C = 128&n != 0
	n <<= 1
	memory.write(addr, n)
	cpu.flags.Z = n == 0
	cpu.flags.N = false
	cpu.flags.H = false
	cpu.pc++
}

func (cpu *cpu) sra_r(r *byte) {
	old7 := 128 & *r
	cpu.flags.C = 1&*r != 0
	*r >>= 1
	if old7 == 1 {
		setBit(r, 7)
	}
	cpu.flags.Z = *r == 0
	cpu.flags.N = false
	cpu.flags.H = false
	cpu.pc++
}

func (cpu *cpu) sra_addr(addr uint16, memory *memory) {
	n := memory.read(addr)
	old7 := 128 & n
	cpu.flags.C = 1&n != 0
	n >>= 1
	if old7 == 1 {
		setBit(&n, 7)
	}
	memory.write(addr, n)
	cpu.flags.Z = n == 0
	cpu.flags.N = false
	cpu.flags.H = false
	cpu.pc++
}

func (cpu *cpu) srl_r(r *byte) {
	cpu.flags.C = 1&*r == 1
	*r >>= 1
	clearBit(r, 7)
	cpu.flags.Z = *r == 0
	cpu.flags.N = false
	cpu.flags.H = false
	cpu.pc++
}

func (cpu *cpu) srl_addr(addr uint16, memory *memory) {
	n := memory.read(addr)
	cpu.flags.C = 1&n == 1
	n >>= 1
	clearBit(&n, 7)
	memory.write(addr, n)
	cpu.flags.Z = n == 0
	cpu.flags.N = false
	cpu.flags.H = false
	cpu.pc++
}

// Bit Opcodes
func (cpu *cpu) bit_r(r *byte, bitPos byte) {
	cpu.flags.Z = (*r>>bitPos)&1 == 0
	cpu.flags.N = false
	cpu.flags.H = true
	cpu.pc += 2
}

func (cpu *cpu) bit_addr(addr uint16, bitPos byte, memory *memory) {
	n := memory.read(addr)
	cpu.flags.Z = (n>>bitPos)&1 == 0
	cpu.flags.N = false
	cpu.flags.H = true
	cpu.pc += 2
}

func (cpu *cpu) set_r(r *byte, bitPos byte) {
	setBit(r, int(bitPos))
	cpu.pc += 2
}

func (cpu *cpu) set_addr(addr uint16, bitPos byte, memory *memory) {
	n := memory.read(addr)
	setBit(&n, int(bitPos))
	memory.write(addr, n)
	cpu.pc += 2
}

func (cpu *cpu) res_r(r *byte, bitPos byte) {
	clearBit(r, int(bitPos))
	cpu.pc += 2
}

func (cpu *cpu) res_addr(addr uint16, bitPos byte, memory *memory) {
	n := memory.read(addr)
	clearBit(&n, int(bitPos))
	memory.write(addr, n)
	cpu.pc += 2
}

// Jumps
func (cpu *cpu) jp_nn(addr uint16) {
	cpu.pc = addr
}

func (cpu *cpu) jr_n(distance byte) {
	singedDist := int8(distance)
	if singedDist >= 0 {
		cpu.pc += uint16(singedDist)
	} else {
		cpu.pc -= uint16(-singedDist)
	}
}

func (cpu *cpu) jp_nn_cc(addr uint16, flagPos int, expected byte) {
	if getBit(flagsToByte(*cpu.flags), flagPos) == expected {
		cpu.pc = addr
	} else {
		cpu.pc += 3
	}
}

func (cpu *cpu) jr_n_cc(distance uint8, flagPos int, expected byte) {
	if getBit(flagsToByte(*cpu.flags), flagPos) == expected {
		signedDist := int8(distance)
		if signedDist >= 0 {
			cpu.pc += uint16(signedDist)
		} else {
			cpu.pc -= uint16(-signedDist)
		}
	} else {
		cpu.pc += 2
	}
}

// Calls
func (cpu *cpu) call_nn(addr uint16, nextInstructionAddr uint16, memory *memory) {
	memory.write(cpu.sp-1, uint8(nextInstructionAddr&0x00FF))
	memory.write(cpu.sp-2, uint8(nextInstructionAddr>>8&0x00FF))
	cpu.sp -= 2
	cpu.pc = addr
}

func (cpu *cpu) call_cc_nn(addr uint16, nextInstructionAddr uint16, flagPos int, expected byte, memory *memory) {
	if getBit(flagsToByte(*cpu.flags), flagPos) == expected {
		cpu.call_nn(addr, nextInstructionAddr, memory)
	} else {
		cpu.pc += 3
	}
}

// Restarts
func (cpu *cpu) rst_n(n byte, memory *memory) {
	memory.write(cpu.sp-1, uint8(cpu.pc>>8&0x00FF))
	memory.write(cpu.sp-2, uint8(cpu.pc&0x00FF))
	cpu.sp -= 2
	cpu.pc = uint16(n)
}

// Returns
func (cpu *cpu) ret(memory *memory) {
	pc := uint16(memory.read(cpu.sp+1)) | uint16(memory.read(cpu.sp))
	cpu.sp += 2
	cpu.pc = pc
}

func (cpu *cpu) reti() {
	// TODO
}

func (cpu *cpu) ret_cc(flagPos int, expected byte, memory *memory) {
	if getBit(flagsToByte(*cpu.flags), flagPos) == expected {
		cpu.ret(memory)
	} else {
		cpu.pc++
	}
}
